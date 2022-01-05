// Package ionic provides a direct representation of the endpoints and objects
// within the Ion Channel API
package ionic

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "time"

    "github.com/ion-channel/ionic/pagination"
    "github.com/ion-channel/ionic/requests"
    "github.com/ion-channel/ionic/responses"
)

const (
    maxIdleConns        = 25
    maxIdleConnsPerHost = 25
    maxPagingLimit      = 100

    DefaultBaseURL = "https://api.ionchannel.io"
)

// IonClient represents a communication layer with the Ion Channel API
type IonClient struct {
    baseURL              url.URL
    client               http.Client
    session              Session
    sessionAutoRenewStop chan struct{}
    // requestModifier is a function that is run on every request
    requestModifier requests.RequestModifier
}

// IonClientOptions represents the options available when creating a new IonClient.
// All the options are optional and will be replaced with working defaults if left empty/nil.
type IonClientOptions struct {
    BaseURL         string
    Client          *http.Client
    RequestModifier requests.RequestModifier
}

// New takes the base URL of the API and returns a client for talking to the API
// and an error if any issues instantiating the client are encountered.
// DEPRECATED: this function is deprecated as of 2022/01/01. Use NewWithOptions instead.
func New(baseURL string) (*IonClient, error) {
    return NewWithOptions(IonClientOptions{BaseURL: baseURL})
}

// NewWithOptions takes an IonClientOptions to construct a client for talking to the API.
// Returns the client and any error that occurs.
func NewWithOptions(options IonClientOptions) (*IonClient, error) {
    if options.BaseURL == "" {
        options.BaseURL = DefaultBaseURL
    }

    if options.Client == nil {
        options.Client = &http.Client{
            Transport: &http.Transport{
                MaxIdleConnsPerHost: maxIdleConnsPerHost,
                MaxIdleConns:        maxIdleConns,
            },
        }
    }

    u, err := url.Parse(options.BaseURL)
    if err != nil {
        return nil, fmt.Errorf("ionic: invalid URL: %v", err.Error())
    }

    ic := &IonClient{
        baseURL:         *u,
        client:          *options.Client,
        requestModifier: options.RequestModifier,
    }

    return ic, nil
}

// NewWithClient takes the base URL of the API and an existing HTTP client.  It
// returns a client for talking to the API and an error if any issues
// instantiating the client are encountered.
// DEPRECATED: this function is deprecated as of 2022/01/01. Use NewWithOptions instead.
func NewWithClient(baseURL string, client http.Client) (*IonClient, error) {
    return NewWithOptions(IonClientOptions{
        BaseURL: baseURL,
        Client:  &client,
    })
}

// WithRequestModifier can be used to create a new temporary IonClient with all the same options as the one this
// method is called on.
// NOTICE: the receiver (ic IonClient) is NOT a pointer, so the receiver is not mutated. This returns a copy.
func (ic IonClient) WithRequestModifier(requestModifier requests.RequestModifier) IonClient {
    ic.requestModifier = requestModifier
    return ic
}

// Delete takes an endpoint, token, params, and headers to pass as a delete call to the
// API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
func (ic *IonClient) Delete(endpoint, token string, params url.Values, headers http.Header) (json.RawMessage, error) {
    return requests.Delete(ic.client, ic.baseURL, endpoint, token, params, headers, ic.requestModifier)
}

// Head takes an endpoint, token, params, headers, and pagination params to pass as a
// head call to the API.  It will return any errors it encounters with the API.
func (ic *IonClient) Head(endpoint, token string, params url.Values, headers http.Header, page pagination.Pagination) error {
    return requests.Head(ic.client, ic.baseURL, endpoint, token, params, headers, page, ic.requestModifier)
}

// Get takes an endpoint, token, params, headers, and pagination params to pass as a
// get call to the API.  It will return a json RawMessage for the response and
// any errors it encounters with the API.
func (ic *IonClient) Get(endpoint, token string, params url.Values, headers http.Header, page pagination.Pagination) (json.RawMessage, *responses.Meta, error) {
    return requests.Get(ic.client, ic.baseURL, endpoint, token, params, headers, page, ic.requestModifier)
}

// Post takes an endpoint, token, params, payload, and headers to pass as a post call
// to the API.  It will return a json RawMessage for the response and any errors
// it encounters with the API.
func (ic *IonClient) Post(endpoint, token string, params url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
    return requests.Post(ic.client, ic.baseURL, endpoint, token, params, payload, headers, ic.requestModifier)
}

// Put takes an endpoint, token, params, payload, and headers to pass as a put call to
// the API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
func (ic *IonClient) Put(endpoint, token string, params url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
    return requests.Put(ic.client, ic.baseURL, endpoint, token, params, payload, headers, ic.requestModifier)
}

// Patch takes an endpoint, token, params, payload, and headers to pass as a patch call to
// the API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
func (ic *IonClient) Patch(endpoint, token string, params url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
    return requests.Patch(ic.client, ic.baseURL, endpoint, token, params, payload, headers, ic.requestModifier)
}

// SetSession sets the client's internal Session that can be used to authenticate when making API requests.
// The session can safely be set to null.
// Example: myClient.GetSelf(myClient.Session().BearerToken)
func (ic *IonClient) SetSession(session Session) {
    ic.session = session
}

// Session returns the client's internal Session.
// This Session is set and renewed automatically if the EnableSessionAutoRenew method is used.
func (ic *IonClient) Session() Session {
    return ic.session
}

// EnableSessionAutoRenew enables the periodic automatic renewal of the IonClient session using the given
// login information, ensuring that the client will always have a valid session token.
// To make the client stop automatically renewing its session, use the DisableSessionAutoRenew method.
func (ic *IonClient) EnableSessionAutoRenew(username, password string) error {
    // try to log in with these credentials immediately, abort if it fails
    session, err := ic.Login(username, password)
    if err != nil {
        return err
    }

    ic.SetSession(session)

    go ic.autoRenewSessionWorker(username, password)

    return nil
}

// DisableSessionAutoRenew makes the IonClient stop automatically renewing its session.
func (ic *IonClient) DisableSessionAutoRenew() {
    close(ic.sessionAutoRenewStop)
}

func (ic *IonClient) autoRenewSessionWorker(username, password string) {
    ticker := time.NewTicker(7 * time.Minute) // sessions expire after 15 minutes
    ic.sessionAutoRenewStop = make(chan struct{})

    for {
        select {
        case <-ticker.C:
            session, err := ic.Login(username, password)
            if err != nil {
                fmt.Printf("ERROR - failed to automatically renew IonClient session: %s", err.Error())
                continue // don't blow everything up in case it was just a temporary issue
            }

            ic.SetSession(session)
        case <-ic.sessionAutoRenewStop:
            ticker.Stop()
            return
        }
    }
}

// Package ionic provides a direct representation of the endpoints and objects
// within the Ion Channel API.
// Use NewDefault or NewWithOptions to create a client.
package ionic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/ion-channel/ionic/pagination"
	"github.com/ion-channel/ionic/requests"
	"github.com/ion-channel/ionic/responses"
)

type contextKey int

const (
	maxIdleConns        = 25
	maxIdleConnsPerHost = 25
	maxPagingLimit      = 100

	ionClientContextKey contextKey = iota
)

// IonClient represents a communication layer with the Ion Channel API
type IonClient struct {
	baseURL              url.URL
	client               http.Client
	session              Session
	sessionAutoRenewStop chan struct{}
	// the global context used for every request this IonClient makes
	ctx context.Context
}

// IonClientOptions represents the options available when creating a new IonClient.
// All the options are optional and will be replaced with working defaults if left empty/nil.
// Some options can be set via environment variables; prefix the envconfig value with "IONIC_" to get the variable name.
type IonClientOptions struct {
	BaseURL string          `envconfig:"BASE_URL" default:"https://api.ionchannel.io"`
	Client  *http.Client    `ignored:"true"`
	Context context.Context `ignored:"true"`
}

// New takes the base URL of the API and returns a client for talking to the API
// and an error if any issues instantiating the client are encountered.
// DEPRECATED: this function is deprecated as of 2022/01/01. Use NewWithOptions instead.
func New(baseURL string) (*IonClient, error) {
	return NewWithOptions(IonClientOptions{BaseURL: baseURL})
}

// NewDefault returns a new default IonClient.
// Some defaults can be overridden using environment variables, see the IonClientOptions struct.
func NewDefault() *IonClient {
	ic, _ := NewWithOptions(IonClientOptions{})

	return ic
}

// NewWithOptions takes an IonClientOptions to construct a client for talking to the API.
// Returns the client and any error that occurs.
// The defaults provided by an empty IonClientOptions object are sane and functional, so all the options are optional.
// Some defaults can be overridden using environment variables, see the IonClientOptions struct.
func NewWithOptions(options IonClientOptions) (*IonClient, error) {
	var defaultOptions IonClientOptions
	err := envconfig.Process("ionic", &defaultOptions)
	if err != nil {
		log.Fatalf("failed to initialize Ionic: %v", err.Error())
	}

	if options.BaseURL == "" {
		options.BaseURL = defaultOptions.BaseURL
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
		baseURL: *u,
		client:  *options.Client,
		ctx:     options.Context,
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

// InjectIntoContext returns the given context with the given IonClient added to it.
func InjectIntoContext(ctx context.Context, ionClient *IonClient) context.Context {
	return context.WithValue(ctx, ionClientContextKey, ionClient)
}

// FromContext retrieves an IonClient from the given context, or returns a new default client if one is not found.
func FromContext(ctx context.Context) *IonClient {
	ionClient, ok := ctx.Value(ionClientContextKey).(*IonClient)
	if !ok {
		ionClient = NewDefault()
	}

	return ionClient
}

// WithContext can be used to create a new temporary IonClient with all the same options as the one this
// method is called on.
// NOTICE: the receiver (ic IonClient) is NOT a pointer, so the receiver is not mutated. This returns a copy.
func (ic IonClient) WithContext(ctx context.Context) *IonClient {
	ic.ctx = ctx
	return &ic
}

// Delete takes an endpoint, token, params, and headers to pass as a delete call to the
// API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
func (ic *IonClient) Delete(endpoint, token string, params url.Values, headers http.Header) (json.RawMessage, error) {
	return requests.Delete(ic.ctx, ic.client, ic.baseURL, endpoint, token, params, headers)
}

// Head takes an endpoint, token, params, headers, and pagination params to pass as a
// head call to the API.  It will return any errors it encounters with the API.
func (ic *IonClient) Head(endpoint, token string, params url.Values, headers http.Header, page pagination.Pagination) error {
	return requests.Head(ic.ctx, ic.client, ic.baseURL, endpoint, token, params, headers, page)
}

// Get takes an endpoint, token, params, headers, and pagination params to pass as a
// get call to the API.  It will return a json RawMessage for the response and
// any errors it encounters with the API.
func (ic *IonClient) Get(endpoint, token string, params url.Values, headers http.Header, page pagination.Pagination) (json.RawMessage, *responses.Meta, error) {
	return requests.Get(ic.ctx, ic.client, ic.baseURL, endpoint, token, params, headers, page)
}

// Post takes an endpoint, token, params, payload, and headers to pass as a post call
// to the API.  It will return a json RawMessage for the response and any errors
// it encounters with the API.
func (ic *IonClient) Post(endpoint, token string, params url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
	return requests.Post(ic.ctx, ic.client, ic.baseURL, endpoint, token, params, payload, headers)
}

// Put takes an endpoint, token, params, payload, and headers to pass as a put call to
// the API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
func (ic *IonClient) Put(endpoint, token string, params url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
	return requests.Put(ic.ctx, ic.client, ic.baseURL, endpoint, token, params, payload, headers)
}

// Patch takes an endpoint, token, params, payload, and headers to pass as a patch call to
// the API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
func (ic *IonClient) Patch(endpoint, token string, params url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
	return requests.Patch(ic.ctx, ic.client, ic.baseURL, endpoint, token, params, payload, headers)
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

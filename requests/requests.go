package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/ion-channel/ionic/errors"
	"github.com/ion-channel/ionic/pagination"
	"github.com/ion-channel/ionic/responses"
)

const (
	maxPagingLimit = 100
)

// ByIDs represents a list of ids in a request
type ByIDs struct {
	IDs []string `json:"ids"`
}

// ByIDsAndTeamID represents data to give team id and a slice of project IDs
type ByIDsAndTeamID struct {
	TeamID string   `json:"team_id"`
	IDs    []string `json:"ids"`
}

// request is an internal container for all the relevant data that makes up an HTTP request
type request struct {
	Client     http.Client
	Headers    http.Header
	Method     string
	BaseURL    url.URL
	Endpoint   string
	Params     url.Values
	Payload    bytes.Buffer
	Pagination pagination.Pagination
	Token      string
}

func do(req request) (json.RawMessage, *responses.Meta, error) {
	if req.Pagination == (pagination.Pagination{}) || req.Pagination.Limit > 0 {
		ir, err := _do(req)
		if err != nil {
			return nil, nil, err
		}

		return ir.Data, &ir.Meta, nil
	}

	req.Pagination = pagination.New(0, maxPagingLimit)
	data := []byte("[")

	total := 1
	for req.Pagination.Offset < total {
		ir, err := _do(req)
		if err != nil {
			err.Prepend("api: paging")
			return nil, nil, err
		}

		data = append(data, ir.Data[1:len(ir.Data)-1]...)
		data = append(data, []byte(",")...)
		req.Pagination.Up()
		total = ir.Meta.TotalCount
	}

	data = append(data[:len(data)-1], []byte("]")...)
	return data, &responses.Meta{TotalCount: total}, nil
}

func _do(req request) (*responses.IonResponse, *errors.IonError) {
	u := createURL(req.BaseURL, req.Endpoint, req.Params, req.Pagination)

	httpReq, err := http.NewRequest(strings.ToUpper(req.Method), u.String(), &req.Payload)
	if err != nil {
		return nil, errors.Errors("no body", 0, "http request: failed to create: %v", err.Error())
	}

	if req.Headers != nil {
		httpReq.Header = req.Headers
	}

	if req.Token != "" {
		httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %v", req.Token))
	}

	resp, err := req.Client.Do(httpReq)
	if err != nil {
		return nil, errors.Errors("no body", 0, "http request: failed: %v", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Errors("no body", resp.StatusCode, "response body: failed to read: %v", err.Error())
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.Errors(string(body), resp.StatusCode, "api error response: %s", string(body))
	}

	if strings.ToUpper(req.Method) == "HEAD" || strings.ToUpper(req.Method) == "DELETE" {
		return &responses.IonResponse{}, nil
	}

	var ir responses.IonResponse
	err = json.Unmarshal(body, &ir)
	if err != nil {
		return nil, errors.Errors(string(body), resp.StatusCode, "api: malformed response: %v", err.Error())
	}

	return &ir, nil
}

func createURL(baseURL url.URL, endpoint string, params url.Values, page pagination.Pagination) url.URL {
	baseURL.Path = endpoint

	if params == nil {
		params = url.Values{}
	}

	// add pagination params to the URL if given
	if page != (pagination.Pagination{}) {
		page.AddParams(&params)
	}

	baseURL.RawQuery = params.Encode()
	return baseURL
}

// Delete takes a client, baseURL, endpoint, token, params, and headers to pass as a delete call to the
// API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
// It is used internally by the SDK
func Delete(client http.Client, baseURL url.URL, endpoint, token string, params url.Values, headers http.Header) (json.RawMessage, error) {
	req := request{
		Client:   client,
		Headers:  headers,
		Method:   "DELETE",
		BaseURL:  baseURL,
		Endpoint: endpoint,
		Params:   params,
		Token:    token,
	}
	r, _, err := do(req)
	return r, err
}

// Head takes a client, baseURL, endpoint, token, params, headers, and pagination params to pass as a
// head call to the API.  It will return any errors it encounters with the API.
// It is used internally by the SDK
func Head(client http.Client, baseURL url.URL, endpoint, token string, params url.Values, headers http.Header, page pagination.Pagination) error {
	req := request{
		Client:     client,
		Headers:    headers,
		Method:     "HEAD",
		BaseURL:    baseURL,
		Endpoint:   endpoint,
		Params:     params,
		Token:      token,
		Pagination: page,
	}
	_, _, err := do(req)
	return err
}

// Get takes a client, baseURL, endpoint, token, params, headers, and pagination params to pass as a
// get call to the API.  It will return a json RawMessage for the response and
// any errors it encounters with the API.
// It is used internally by the SDK
func Get(client http.Client, baseURL url.URL, endpoint, token string, params url.Values, headers http.Header, page pagination.Pagination) (json.RawMessage, *responses.Meta, error) {
	req := request{
		Client:     client,
		Headers:    headers,
		Method:     "GET",
		BaseURL:    baseURL,
		Endpoint:   endpoint,
		Params:     params,
		Token:      token,
		Pagination: page,
	}
	r, m, err := do(req)
	return r, m, err
}

// Post takes a client, baseURL, endpoint, token, params, payload, and headers to pass as a post call
// to the API.  It will return a json RawMessage for the response and any errors
// it encounters with the API.
// It is used internally by the SDK
func Post(client http.Client, baseURL url.URL, endpoint, token string, params url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
	req := request{
		Client:   client,
		Headers:  headers,
		Method:   "POST",
		BaseURL:  baseURL,
		Endpoint: endpoint,
		Params:   params,
		Token:    token,
		Payload:  payload,
	}
	r, _, err := do(req)
	return r, err
}

// Put takes a client, baseURL, endpoint, token, params, payload, and headers to pass as a put call to
// the API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
// It is used internally by the SDK
func Put(client http.Client, baseURL url.URL, endpoint, token string, params url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
	req := request{
		Client:   client,
		Headers:  headers,
		Method:   "PUT",
		BaseURL:  baseURL,
		Endpoint: endpoint,
		Params:   params,
		Token:    token,
		Payload:  payload,
	}
	r, _, err := do(req)
	return r, err
}

// Patch takes a client, baseURL, endpoint, token, params, payload, and headers to pass as a patch call to
// the API.  It will return a json RawMessage for the response and any errors it
// encounters with the API.
// It is used internally by the SDK
func Patch(client http.Client, baseURL url.URL, endpoint, token string, params url.Values, payload bytes.Buffer, headers http.Header) (json.RawMessage, error) {
	req := request{
		Client:   client,
		Headers:  headers,
		Method:   "PATCH",
		BaseURL:  baseURL,
		Endpoint: endpoint,
		Params:   params,
		Token:    token,
		Payload:  payload,
	}
	r, _, err := do(req)
	return r, err
}

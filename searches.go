package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ion-channel/ionic/pagination"
	"net/url"
	"time"

	"github.com/ion-channel/ionic/community"
	"github.com/ion-channel/ionic/dependencies"
	"github.com/ion-channel/ionic/products"
	"github.com/ion-channel/ionic/responses"
	"github.com/ion-channel/ionic/searches"
)

const (
	searchEndpoint = "v1/search"
)

// SearchMatch structure for holding multiple search response types
type SearchMatch struct {
	// Requires common fields to be explicitly
	// defined here
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at"`
	Confidence float32   `json:"confidence"`
	Version    string    `json:"version,omitempty"`
	Org        string    `json:"org,omitempty"`
	Type       string    `json:"type,omitempty"`
	ExternalID string    `json:"external_id,omitempty"`

	// Clean up the output
	Vulnerabilities *interface{} `json:"vulnerabilities,omitempty"`
	Source          *interface{} `json:"source,omitempty"`
	References      *interface{} `json:"references,omitempty"`
	Aliases         *interface{} `json:"aliases,omitempty"`
	Dependencies    *interface{} `json:"dependencies,omitempty"`

	*community.Repo          `json:",omitempty"`
	*products.Product        `json:",omitempty"`
	*dependencies.Dependency `json:",omitempty"`
	*searches.Report         `json:",omitempty"`
}

// GetSearch takes a query to perform and a to be searched param
// a productidentifier search against the Ion API, assembling a slice of Ionic
// products.ProductSearchResponse objects
func (ic *IonClient) GetSearch(query, tbs, token string) ([]SearchMatch, *responses.Meta, error) {
	params := url.Values{}
	params.Set("q", query)
	params.Set("tbs", tbs)

	b, m, err := ic.Get(searchEndpoint, token, params, nil, pagination.Pagination{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get productidentifiers search: %v", err.Error())
	}

	var results []SearchMatch
	err = json.Unmarshal(b, &results)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal product search results: %v", err.Error())
	}
	return results, m, nil

}

// BulkSearch takes one or more query strings and a "to be searched" param, then performs a productidentifier search
// against the Ion API, returning a map of the original query string(s) to SearchMatch objects
func (ic *IonClient) BulkSearch(queries []string, tbs, token string) (map[string][]SearchMatch, error) {
	params := url.Values{}
	params.Set("tbs", tbs)

	body, err := json.Marshal(queries)
	if err != nil {
		return nil, fmt.Errorf("session: failed to marshal login body: %v", err.Error())
	}

	buff := bytes.NewBuffer(body)

	b, err := ic.Post(searchEndpoint, token, params, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get productidentifiers search: %v", err.Error())
	}

	var results map[string][]SearchMatch
	err = json.Unmarshal(b, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal product search results: %v", err.Error())
	}

	return results, nil

}

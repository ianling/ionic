package ionic

import (
	"encoding/json"
	"fmt"
	"github.com/ion-channel/ionic/pagination"
	"github.com/ion-channel/ionic/projects"
	"net/url"
)

// CreateSoftwareListRequest defines the body of a CreateSoftwareList request.
type CreateSoftwareListRequest struct {
	Name             string `json:"name"`
	OrganizationId   string `json:"org_id"`
	Version          string `json:"version"`
	Supplier         string `json:"supplier_name"`
	ContactName      string `json:"contact_name"`
	ContactEmail     string `json:"contact_email"`
	RulesetID        string `json:"ruleset_id"`
	MonitorFrequency string `json:"monitor_frequency"`
}

// GetSoftwareListRequest defines the parameters available for a GetSoftwareList request.
type GetSoftwareListRequest struct {
	// ID (required)	-- Software List ID
	ID string
}

// GetSoftwareListsRequest defines the parameters available for a GetSoftwareLists request.
type GetSoftwareListsRequest struct {
	// OrganizationID (required)	-- Organization ID
	OrganizationID string
	// Status (optional)			-- filters on the Status field of the Software Lists, if provided. Ignored if blank.
	Status string
}

// GetSoftwareList returns the requested Software List or any error that occurred.
func (ic *IonClient) GetSoftwareList(req GetSoftwareListRequest, token string) (projects.SoftwareList, error) {
	var sbom projects.SoftwareList

	params := url.Values{}
	params.Set("id", req.ID)

	b, _, err := ic.Get(projects.GetSoftwareListEndpoint, token, params, nil, pagination.Pagination{})
	if err != nil {
		return sbom, fmt.Errorf("failed to get SoftwareList: %s", err.Error())
	}

	err = json.Unmarshal(b, &sbom)
	if err != nil {
		return sbom, fmt.Errorf("failed to unmarshal SoftwareList: %s", err.Error())
	}

	return sbom, nil
}

// GetSoftwareLists retrieves an organization's Software Lists, filtered on the status, if given, or any error that occurred.
func (ic *IonClient) GetSoftwareLists(req GetSoftwareListsRequest, token string) ([]projects.SoftwareList, error) {
	params := url.Values{}
	params.Set("org_id", req.OrganizationID)
	params.Set("status", req.Status)

	b, _, err := ic.Get(projects.GetSoftwareListsEndpoint, token, params, nil, pagination.Pagination{})
	if err != nil {
		return nil, fmt.Errorf("failed to get SBOMs: %s", err.Error())
	}

	var sboms []projects.SoftwareList
	err = json.Unmarshal(b, &sboms)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal SBOMs: %s", err.Error())
	}

	return sboms, nil
}

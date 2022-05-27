package ionic

import (
	"encoding/json"
	"fmt"
	"github.com/ion-channel/ionic/pagination"
	"net/url"
)

const (
	// GetSoftwareListEndpoint is the path to the endpoint for retrieving a specific Software List
	GetSoftwareListEndpoint = "v1/project/getSBOM"
	// GetSoftwareListsEndpoint is the path to the endpoint for retrieving an organization's Software Lists
	GetSoftwareListsEndpoint = "v1/project/getSBOMs"
)

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
func (ic *IonClient) GetSoftwareList(req GetSoftwareListRequest, token string) (SoftwareList, error) {
	var sbom SoftwareList

	params := url.Values{}
	params.Set("id", req.ID)

	b, _, err := ic.Get(GetSoftwareListEndpoint, token, params, nil, pagination.Pagination{})
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
func (ic *IonClient) GetSoftwareLists(req GetSoftwareListsRequest, token string) ([]SoftwareList, error) {
	params := url.Values{}
	params.Set("org_id", req.OrganizationID)
	params.Set("status", req.Status)

	b, _, err := ic.Get(GetSoftwareListsEndpoint, token, params, nil, pagination.Pagination{})
	if err != nil {
		return nil, fmt.Errorf("failed to get SBOMs: %s", err.Error())
	}

	var sboms []SoftwareList
	err = json.Unmarshal(b, &sboms)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal SBOMs: %s", err.Error())
	}

	return sboms, nil
}

package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/pagination"
	"github.com/ion-channel/ionic/software_lists"
)

const (
	// GetSoftwareListEndpoint is the path to the endpoint for retrieving a specific Software List
	GetSoftwareListEndpoint = "v1/project/getSBOM"
	// GetSoftwareListsEndpoint is the path to the endpoint for retrieving an organization's Software Lists
	GetSoftwareListsEndpoint = "v1/project/getSBOMs"
	// GetSoftwareListEndpoint is the path to the endpoint for deleting an organization's Software List
	DeleteSoftwareListEndpoint = "/v1/project/deleteSBOM"
	// UpdateSoftwareListEndpoint is the path to the endpoint for updating an organization's Software List
	UpdateSoftwareListEndpoint = "/v1/project/updateSBOM"
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

// DeleteSoftwareList deletes the requested Software List or any error that occurred.
func (ic *IonClient) DeleteSoftwareList(id string, token string) error {
	params := url.Values{}
	params.Set("id", id)

	_, err := ic.Delete(DeleteSoftwareListEndpoint, token, params, nil)
	if err != nil {
		return fmt.Errorf("failed to delete SoftwareList: %s", err.Error())
	}

	return nil
}

// UpdateSoftwareList updates the requested Software List or any error that occurred.
func (ic *IonClient) UpdateSoftwareList(sbom software_lists.SoftwareList, token string) (*software_lists.SoftwareList, error) {
	b, err := json.Marshal(sbom)
	if err != nil {
		return nil, fmt.Errorf("session: failed to marshal login body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	b, err = ic.Put(UpdateSoftwareListEndpoint, token, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete SoftwareList: %s", err.Error())
	}

	err = json.Unmarshal(b, &sbom)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal delete response: %s", err.Error())
	}

	return &sbom, nil
}

// GetSoftwareList returns the requested Software List or any error that occurred.
func (ic *IonClient) GetSoftwareList(req GetSoftwareListRequest, token string) (software_lists.SoftwareList, error) {
	var sbom software_lists.SoftwareList

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
func (ic *IonClient) GetSoftwareLists(req GetSoftwareListsRequest, token string) ([]software_lists.SoftwareList, error) {
	var sboms []software_lists.SoftwareList

	params := url.Values{}
	params.Set("org_id", req.OrganizationID)
	params.Set("status", req.Status)

	b, _, err := ic.Get(GetSoftwareListsEndpoint, token, params, nil, pagination.Pagination{})
	if err != nil {
		return nil, fmt.Errorf("failed to get SBOMs: %s", err.Error())
	}

	/* TODO GetSoftwareListsEndpoint should return SoftwareList, not SoftwareInventory.

	   Then we can simply unmarshal a SoftwareList and
	   no need to extract sboms from Inventory.
	*/
	inventory := software_lists.SoftwareInventory{}
	err = json.Unmarshal(b, &inventory)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal into sbom: %s %s", err.Error(), b)
	}

	// Extract software list from inventory
	sboms = inventory.SoftwareLists

	return sboms, nil
}

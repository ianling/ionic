package ionic

import (
	"encoding/json"
	"fmt"
	"github.com/ion-channel/ionic/pagination"
	"github.com/ion-channel/ionic/projects"
	"net/url"
)

// GetSbom takes an SBOM ID.
// Returns the requested SBOM, or any error that occurred.
func (ic *IonClient) GetSbom(id, token string) (projects.SBOM, error) {
	var sbom projects.SBOM

	params := url.Values{}
	params.Set("id", id)

	b, _, err := ic.Get(projects.GetSbomEndpoint, token, params, nil, pagination.Pagination{})
	if err != nil {
		return sbom, fmt.Errorf("failed to get SBOM: %s", err.Error())
	}

	err = json.Unmarshal(b, &sbom)
	if err != nil {
		return sbom, fmt.Errorf("failed to unmarshal SBOM: %s", err.Error())
	}

	return sbom, nil
}

// GetSboms takes an organization ID, and a status to filter on, if given.
// Returns the organization's SBOMs, filtered on the status if given, or any error that occurred.
func (ic *IonClient) GetSboms(orgID, status, token string) ([]projects.SBOM, error) {
	params := url.Values{}
	params.Set("org_id", orgID)
	params.Set("status", status)

	b, _, err := ic.Get(projects.GetSbomsEndpoint, token, params, nil, pagination.Pagination{})
	if err != nil {
		return nil, fmt.Errorf("failed to get SBOMs: %s", err.Error())
	}

	var sboms []projects.SBOM
	err = json.Unmarshal(b, &sboms)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal SBOMs: %s", err.Error())
	}

	return sboms, nil
}

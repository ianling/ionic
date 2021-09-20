package ionic

import (
    "encoding/json"
    "fmt"
    "github.com/ion-channel/ionic/projects"
    "net/url"
)

// GetSboms takes an organization ID, and a status to filter on, if given.
// Returns the organization's SBOMs, filtered on the status if given, or any error that occurred.
func (ic *IonClient) GetSboms(orgId, status, token string) ([]projects.SBOM, error) {
    params := &url.Values{}
    params.Set("org_id", orgId)
    params.Set("status", status)

    b, _, err := ic.Get(projects.GetSbomsEndpoint, token, params, nil, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to get project: %s", err.Error())
    }

    var sbom []projects.SBOM
    err = json.Unmarshal(b, &sbom)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal SBOMs: %s", err.Error())
    }

    return sbom, nil
}

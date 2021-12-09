package projects

import "time"

const (
	// GetSbomEndpoint is the current endpoint for retrieving an organization's SBOMs
	GetSbomEndpoint = "v1/project/getSBOM"
	// GetSbomsEndpoint is the current endpoint for retrieving an organization's SBOMs
	GetSbomsEndpoint = "v1/project/getSBOMs"
)

// SBOM represents a software list containing zero or more SBOMEntry objects.
type SBOM struct {
	ID         string      `json:"id"`
	Name       string      `json:"sbom_name"`
	Version    string      `json:"sbom_version"`
	Supplier   string      `json:"supplier_name"`
	SbomStatus string      `json:"sbom_status"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	EntryCount int         `json:"entry_count"`
	Entries    []SBOMEntry `json:"entries"`
	TeamID     string      `json:"team_id"`
	OrgID      string      `json:"org_id"`
	RulesetID  string      `json:"ruleset_id"`
}

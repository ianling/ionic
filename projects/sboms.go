package projects

import "time"

const (
	// GetSbomEndpoint is the current endpoint for retrieving an organization's SBOMs
	GetSbomEndpoint = "v1/project/getSBOM"
	// GetSbomsEndpoint is the current endpoint for retrieving an organization's SBOMs
	GetSbomsEndpoint = "v1/project/getSBOMs"
)

// SBOMMetadata contains various piece of metadata about a particular SBOM
type SBOMMetadata struct {
	EntryCount                  int `json:"entry_count"`
	ResolvedEntryCount          int `json:"resolved_entry_count"`
	PartiallyResolvedEntryCount int `json:"partially_resolved_entry_count"`
	UnresolvedEntryCount        int `json:"unresolved_entry_count"`
}

// SBOM represents a software list containing zero or more SBOMEntry objects.
type SBOM struct {
	ID               string       `json:"id"`
	Name             string       `json:"name"`
	Version          string       `json:"version"`
	Supplier         string       `json:"supplier"`
	ContactName      string       `json:"contact_name"`
	ContactEmail     string       `json:"contact_email"`
	MonitorFrequency string       `json:"monitor_frequency"`
	SbomStatus       string       `json:"status"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
	DeletedAt        *time.Time   `json:"deleted_at"`
	EntryCount       int          `json:"entry_count"`
	Metadata         SBOMMetadata `json:"metadata"`
	Stats            Stats        `json:"stats"`
	Entries          []SBOMEntry  `json:"entries"`
	TeamID           string       `json:"team_id"`
	OrgID            string       `json:"org_id"`
	RulesetID        string       `json:"ruleset_id"`
}

type Risk struct {
	Score  int `json:"score"`
	Scopes struct {
		Ecosystem   int `json:"ecosystem"`
		SupplyChain int `json:"supply_chain"`
		Software    int `json:"software"`
	} `json:"scopes"`
}

type Compliance struct {
	Passing int `json:"passing"`
	Failing int `json:"failing"`
}

type Resolution struct {
	Resolved          int `json:"resolved"`
	PartiallyResolved int `json:"partially_resolved"`
	Unresolved        int `json:"unresolved"`
}

type Stats struct {
	Risk       Risk       `json:"risk"`
	Compliance Compliance `json:"compliance"`
	Resolution Resolution `json:"resolution"`
}

type SoftwareInventorySummary struct {
	ID            string `json:"id"`
	Organization  Stats  `json:"organization"`
	SoftwareLists []SBOM `json:"software_lists"`
}

package projects

import "time"

const (
	// GetSoftwareListEndpoint is the path to the endpoint for retrieving a specific Software List
	GetSoftwareListEndpoint = "v1/project/getSBOM"
	// GetSoftwareListsEndpoint is the path to the endpoint for retrieving an organization's Software Lists
	GetSoftwareListsEndpoint = "v1/project/getSBOMs"
)

// SoftwareList represents a software list containing zero or more Component objects.
type SoftwareList struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	Version          string      `json:"version"`
	Supplier         string      `json:"supplier"`
	ContactName      string      `json:"contact_name"`
	ContactEmail     string      `json:"contact_email"`
	MonitorFrequency string      `json:"monitor_frequency"`
	Status           string      `json:"status"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	DeletedAt        *time.Time  `json:"deleted_at"`
	EntryCount       int         `json:"entry_count"`
	Stats            Stats       `json:"stats"`
	Entries          []Component `json:"entries"`
	TeamID           string      `json:"team_id"`
	OrgID            string      `json:"org_id"`
	RulesetID        string      `json:"ruleset_id"`
}

type Risk struct {
	Score  int `json:"score"`
	Scopes struct {
		Ecosystem   int `json:"ecosystem"`
		SupplyChain int `json:"supplyChain"`
		Software    int `json:"software"`
	} `json:"scopes"`
}

type Compliance struct {
	Passing int `json:"passing"`
	Failing int `json:"failing"`
}

type Resolution struct {
	Resolved          int `json:"resolved"`
	PartiallyResolved int `json:"partiallyResolved"`
	Unresolved        int `json:"unresolved"`
}

type Stats struct {
	Risk       Risk       `json:"risk"`
	Compliance Compliance `json:"compliance"`
	Resolution Resolution `json:"resolution"`
}

type SoftwareInventorySummary struct {
	ID            string         `json:"id"`
	Organization  Stats          `json:"organization"`
	SoftwareLists []SoftwareList `json:"softwareLists"`
}

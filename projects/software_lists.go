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
	Metrics          Metrics     `json:"metrics"`
	Entries          []Component `json:"entries"`
	TeamID           string      `json:"team_id"`
	OrgID            string      `json:"org_id"`
	RulesetID        string      `json:"ruleset_id"`
}

type Risk struct {
	Score  *int            `json:"score"`
	Scopes map[string]*int `json:"scopes"`
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

type Metrics struct {
	Risk       Risk       `json:"risk"`
	Compliance Compliance `json:"compliance"`
	Resolution Resolution `json:"resolution"`
}

type SoftwareInventorySummary struct {
	ID            string         `json:"id"`
	Organization  Metrics        `json:"organization"`
	SoftwareLists []SoftwareList `json:"softwareLists"`
}

// NewMetrics returns a Metrics struct.
// Used to ensure Metrics.Risk.Scopes is not nil.
func NewMetrics() Metrics {
	return Metrics{
		Risk: Risk{
			Scopes: make(map[string]*int),
		},
	}
}

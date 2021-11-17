package projects

import "time"

const (
	// GetSbomEndpoint is the current endpoint for retrieving an organization's SBOMs
	GetSbomEndpoint = "v1/project/getSBOM"
	// GetSbomsEndpoint is the current endpoint for retrieving an organization's SBOMs
	GetSbomsEndpoint = "v1/project/getSBOMs"
)

// SourceDetails contains the original package information retrieved directly from an uploaded SBOM.
type SourceDetails struct {
	Name    string `json:"sbom_name"`
	Org     string `json:"sbom_org"`
	Version string `json:"sbom_version"`
}

// SBOMEntryType represents a search result's type
type SBOMEntryType string

const (
	// SBOMEntryTypePackage represents the Package search result type
	SBOMEntryTypePackage SBOMEntryType = "package"
	// SBOMEntryTypeRepo represents the Repo search result type
	SBOMEntryTypeRepo SBOMEntryType = "repo"
	// SBOMEntryTypeProduct represents the Product search result type
	SBOMEntryTypeProduct SBOMEntryType = "product"
	// SBOMEntryTypeError denotes that the search result represents an error.
	// The error message can be found in the ErrMsg field
	SBOMEntryTypeError SBOMEntryType = "error"
)

// SBOMEntryStatus represents a search result's status
type SBOMEntryStatus string

const (
	// SBOMEntryStatusNoResolution means no results were found for the entry
	SBOMEntryStatusNoResolution SBOMEntryStatus = "no-resolution"
	// SBOMEntryStatusPartialResolution means some results were found for the entry, but none have been selected
	SBOMEntryStatusPartialResolution SBOMEntryStatus = "partial-resolution"
	// SBOMEntryStatusResolved means a result has been selected
	SBOMEntryStatusResolved SBOMEntryStatus = "resolved"
	// SBOMEntryStatusErrored means the API experienced an internal error while generating results for the search
	SBOMEntryStatusErrored SBOMEntryStatus = "errored"
	// SBOMEntryStatusDeleted means the entry was deleted
	SBOMEntryStatusDeleted SBOMEntryStatus = "deleted"
)

// SBOMEntry represents a single entry within an SBOM.
type SBOMEntry struct {
	ID             string          `json:"id"`
	SBOMID         string          `json:"sbom_id"`
	Type           SBOMEntryType   `json:"type"`
	Confidence     float32         `json:"confidence"`
	Name           string          `json:"name"`
	Org            string          `json:"org"`
	Version        string          `json:"version"`
	IonID          string          `json:"ion_id"`
	Selected       bool            `json:"selected"`
	LocationInSBOM int             `json:"location_in_sbom"`
	Source         SourceDetails   `json:"source"`
	ErrMsg         string          `json:"error_message"`
	ProductID      string          `json:"product_id"` // CPE
	PackageID      string          `json:"package_id"` // PURL
	Repo           string          `json:"repo"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	Status         SBOMEntryStatus `json:"status"`
}

// SBOM represents a software list containing zero or more SBOMEntry objects.
type SBOM struct {
	ID         string      `json:"id"`
	Name       string      `json:"sbom_name"`
	Version    string      `json:"sbom_version"`
	Supplier   string      `json:"supplier_name"`
	SbomType   string      `json:"sbom_type"`
	SbomStatus string      `json:"sbom_status"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	EntryCount int         `json:"entry_count"`
	Entries    []SBOMEntry `json:"entries"`
	TeamID     string      `json:"team_id"`
	OrgID      string      `json:"org_id"`
	RulesetID  string      `json:"ruleset_id"`
}

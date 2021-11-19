package projects

import "time"

const (
	// GetSbomEndpoint is the current endpoint for retrieving an organization's SBOMs
	GetSbomEndpoint = "v1/project/getSBOM"
	// GetSbomsEndpoint is the current endpoint for retrieving an organization's SBOMs
	GetSbomsEndpoint = "v1/project/getSBOMs"
)

// SBOMSearchResultType represents a search result's type
type SBOMSearchResultType string

const (
	// SBOMSearchResultTypePackage represents the Package search result type
	SBOMSearchResultTypePackage SBOMSearchResultType = "package"
	// SBOMSearchResultTypeRepo represents the Repo search result type
	SBOMSearchResultTypeRepo SBOMSearchResultType = "repo"
	// SBOMSearchResultTypeProduct represents the Product search result type
	SBOMSearchResultTypeProduct SBOMSearchResultType = "product"
	// SBOMSearchResultTypeError denotes that the search result represents an error.
	// The error message can be found in the ErrMsg field
	SBOMSearchResultTypeError SBOMSearchResultType = "error"
)

// SBOMSearchResultGeneric contains the fields common to all SBOM search result types
type SBOMSearchResultGeneric struct {
	ID         string  `json:"id"`
	Confidence float32 `json:"confidence"`
	Selected   bool    `json:"selected"`
	Name       string  `json:"name"`
	Org        string  `json:"org"`
	Version    string  `json:"version"`
}

// SBOMPackageSearchResult contains the fields specific to Package search results
type SBOMPackageSearchResult struct {
	SBOMSearchResultGeneric
	PURL string `json:"purl"`
}

// SBOMProductSearchResult contains the fields specific to Product search results
type SBOMProductSearchResult struct {
	SBOMSearchResultGeneric
	CPE string `json:"cpe"`
}

// SBOMRepoSearchResult contains the fields specific to Repo search results
type SBOMRepoSearchResult struct {
	SBOMSearchResultGeneric
	RepoURL string `json:"repo_url"`
}

// SBOMSearchResults is a container for all the different search result types we can find for an SBOM entry
type SBOMSearchResults struct {
	Package []SBOMPackageSearchResult `json:"package"`
	Product []SBOMProductSearchResult `json:"product"`
	Repo    []SBOMRepoSearchResult    `json:"repo"`
}

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
	ID             string            `json:"id"`
	SBOMID         string            `json:"sbom_id"`
	LocationInSBOM int               `json:"location_in_sbom"`
	Name           string            `json:"name"`
	Org            string            `json:"org"`
	Version        string            `json:"version"`
	Status         SBOMEntryStatus   `json:"status"`
	SearchResults  SBOMSearchResults `json:"search_results"`
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
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

package projects

// SearchResultType represents a search result's type
type SearchResultType string

const (
	// SearchResultTypePackage represents the Package search result type
	SearchResultTypePackage SearchResultType = "package"
	// SearchResultTypeRepo represents the Repo search result type
	SearchResultTypeRepo SearchResultType = "repo"
	// SearchResultTypeProduct represents the Product search result type
	SearchResultTypeProduct SearchResultType = "product"
)

// SearchResultGeneric contains the fields common to all SoftwareList search result types
type SearchResultGeneric struct {
	ID                    string  `json:"id"`
	Confidence            float32 `json:"confidence"`
	IsUserInput           bool    `json:"is_user_input"`
	Selected              bool    `json:"selected"`
	AutomaticallySelected bool    `json:"automatically_selected"`
	Name                  string  `json:"name"`
	Org                   string  `json:"org"`
	Version               string  `json:"version"`
}

// PackageSearchResult contains the fields specific to Package search results
type PackageSearchResult struct {
	SearchResultGeneric
	PURL string `json:"purl"`
}

// ProductSearchResult contains the fields specific to Product search results
type ProductSearchResult struct {
	SearchResultGeneric
	CPE string `json:"cpe"`
}

// RepoSearchResult contains the fields specific to Repo search results
type RepoSearchResult struct {
	SearchResultGeneric
	RepoURL string `json:"repo_url"`
}

// SearchResults is a container for all the different search result types we can find for a SoftwareList Component
type SearchResults struct {
	Package []PackageSearchResult `json:"package"`
	Product []ProductSearchResult `json:"product"`
	Repo    []RepoSearchResult    `json:"repo"`
}

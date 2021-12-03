package projects

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

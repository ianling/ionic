package software_lists

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Component struct {
	ID            string                `json:"id"`
	SbomID        string                `json:"sbom_id"`
	ProjectID     *string               `json:"project_id"`
	Name          string                `json:"name"`
	Version       string                `json:"version"`
	Org           string                `json:"org"`
	Status        ComponentStatus       `json:"status"`
	SearchResults SearchResults         `json:"search_results"`
	Suggestions   []ComponentSuggestion `json:"suggestions"`
	CreatedAt     time.Time             `json:"created_at"`
	UpdatedAt     time.Time             `json:"updated_at"`
	DeletedAt     *time.Time            `json:"deleted_at"`
	ErrorMessage  *string               `json:"error_message"`
}

type ComponentSuggestion struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ComponentStatus string

const (
	ComponentStatusNoResolution      ComponentStatus = "no_resolution"
	ComponentStatusPartialResolution ComponentStatus = "partial_resolution"
	ComponentStatusResolved          ComponentStatus = "resolved"
	ComponentStatusErrored           ComponentStatus = "errored"
	ComponentStatusDeleted           ComponentStatus = "deleted"
)

var AllComponentStatus = []ComponentStatus{
	ComponentStatusNoResolution,
	ComponentStatusPartialResolution,
	ComponentStatusResolved,
	ComponentStatusErrored,
	ComponentStatusDeleted,
}

func (e ComponentStatus) IsValid() bool {
	switch e {
	case ComponentStatusNoResolution, ComponentStatusPartialResolution, ComponentStatusResolved, ComponentStatusErrored, ComponentStatusDeleted:
		return true
	}
	return false
}

func (e ComponentStatus) String() string {
	return string(e)
}

func (e *ComponentStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ComponentStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ComponentStatus", str)
	}
	return nil
}

func (e ComponentStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SearchResult struct {
	ID                    string  `json:"id"`
	Confidence            float64 `json:"confidence"`
	IsUserInput           bool    `json:"is_user_input"`
	Selected              bool    `json:"selected"`
	AutomaticallySelected bool    `json:"automatically_selected"`
	Name                  string  `json:"name"`
	Org                   string  `json:"org"`
	Version               string  `json:"version"`
}

type SearchResultType string

const (
	SearchResultTypePackage SearchResultType = "package"
	SearchResultTypeProduct SearchResultType = "product"
	SearchResultTypeRepo    SearchResultType = "repo"
)

var AllSearchResultType = []SearchResultType{
	SearchResultTypePackage,
	SearchResultTypeProduct,
	SearchResultTypeRepo,
}

func (e SearchResultType) IsValid() bool {
	switch e {
	case SearchResultTypePackage, SearchResultTypeProduct, SearchResultTypeRepo:
		return true
	}
	return false
}

func (e SearchResultType) String() string {
	return string(e)
}

func (e *SearchResultType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SearchResultType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SearchResultType", str)
	}
	return nil
}

func (e SearchResultType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SearchResults struct {
	Package []PackageSearchResult `json:"package"`
	Repo    []RepoSearchResult    `json:"repo"`
	Product []ProductSearchResult `json:"product"`
}

type PackageSearchResult struct {
	SearchResult
	Purl string `json:"purl"`
}

type ProductSearchResult struct {
	SearchResult
	Cpe string `json:"cpe"`
}

type RepoSearchResult struct {
	SearchResult
	RepoURL string `json:"repo_url"`
}

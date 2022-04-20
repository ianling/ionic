package projects

import "time"

// ComponentStatus represents a search result's status
type ComponentStatus string

const (
	// ComponentStatusNoResolution means no results were found for the component
	ComponentStatusNoResolution ComponentStatus = "no-resolution"
	// ComponentStatusPartialResolution means some results were found for the component, but none have been selected
	ComponentStatusPartialResolution ComponentStatus = "partial-resolution"
	// ComponentStatusResolved means a result has been selected
	ComponentStatusResolved ComponentStatus = "resolved"
	// ComponentStatusErrored means the API experienced an internal error while generating results for the search
	ComponentStatusErrored ComponentStatus = "errored"
	// ComponentStatusDeleted means the component was deleted
	ComponentStatusDeleted ComponentStatus = "deleted"
)

// Component represents a single component within a SoftwareList.
type Component struct {
	ID             string            `json:"id"`
	SoftwareListID string            `json:"sbom_id"`
	Name           string            `json:"name"`
	Org            string            `json:"org"`
	Version        string            `json:"version"`
	Status         ComponentStatus   `json:"status"`
	SearchResults  SearchResults     `json:"search_results"`
	Suggestions    map[string]string `json:"suggestions"`
	ErrorMessage   string            `json:"error_message"`
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
	DeletedAt      *time.Time        `json:"deleted_at"`
}

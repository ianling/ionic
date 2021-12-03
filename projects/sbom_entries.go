package projects

import "time"

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
    Name           string            `json:"name"`
    Org            string            `json:"org"`
    Version        string            `json:"version"`
    Status         SBOMEntryStatus   `json:"status"`
    SearchResults  SBOMSearchResults `json:"search_results"`
    CreatedAt      time.Time         `json:"created_at"`
    UpdatedAt      time.Time         `json:"updated_at"`
}

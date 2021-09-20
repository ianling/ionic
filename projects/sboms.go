package projects

import "time"

const (
    // GetSbomsEndpoint is the current endpoint for retrieving an organization's SBOMs
    GetSbomsEndpoint = "v1/project/getSBOMs"
)

type SourceDetails struct {
    Name    string `json:"sbom_name"`
    Org     string `json:"sbom_org"`
    Version string `json:"sbom_version"`
}

type SBOMEntry struct {
    ID             string        `json:"id"`
    SBOMID        string        `json:"sbom_id"`
    Confidence     float32       `json:"confidence"`
    Name           string        `json:"name"`
    Org            string        `json:"org"`
    Version        string        `json:"version"`
    IonID         string        `json:"ion_id"`
    Selected       bool          `json:"selected"`
    LocationInSBOM int           `json:"location_in_sbom"`
    Source         SourceDetails `json:"source"`
}

type SBOM struct {
    ID         string      `json:"id"`
    SbomType   string      `json:"sbom_type"`
    SbomStatus string      `json:"sbom_status"`
    CreatedAt  time.Time   `json:"created_at"`
    Entries    []SBOMEntry `json:"entries"`
    TeamID     string      `json:"team_id"`
    OrgID      string      `json:"org_id"`
    RulesetID  string      `json:"ruleset_id"`
}

package reports

import (
	"net/url"
	"strconv"
)

const (
	// ReportGetSBOMEndpoint is the endpoint for generating SBOMs
	ReportGetSBOMEndpoint = "v1/report/getSBOM"
)

// SBOMFormat is a string enum for the accepted SBOM formats that we can export
type SBOMFormat string

const (
	// SBOMFormatSPDX is the enum value for the SPDX SBOM format
	SBOMFormatSPDX SBOMFormat = "SPDX"
	// SBOMFormatCycloneDX is the enum value for the CycloneDX SBOM format
	SBOMFormatCycloneDX SBOMFormat = "CycloneDX"
	// SBOMFormatGitlab is the enum value for the Gitlab SBOM format
	SBOMFormatGitlab SBOMFormat = "Gitlab"
)

// SBOMExportOptions represents all the different settings a user can specify for how the SBOM is exported.
// Specify only one of the following data sources to generate the SBOM from:
//  * ProjectIDs (a slice of one or more project IDs)
//  * TeamID (the ID of a team containing one or more projects; this will use all the team's projects)
//  * SBOMID (the ID of a software list containing one or more components; this will use all the list's components)
//  * SBOMEntryIDs (a slice of one or more software list component IDs)
// Format (required) specifies which format/standard the SBOM will be exported in.
// IncludeDependencies will include all the direct and transitive dependencies of each item in the SBOM if true,
// or exclude all the dependencies if false, leaving only the items themselves.
// TeamIsTopLevel applies only to SBOMs generated using the TeamID field. If true, the top-level item in the SBOM's
// hierarchy will be the team. If false, all the team's projects will be on the top level of the hierarchy.
type SBOMExportOptions struct {
	ProjectIDs   []string `json:"ids"`
	TeamID       string   `json:"team_id"`
	SBOMID       string   `json:"sbom_id"`
	SBOMEntryIDs []string `json:"sbom_entry_ids"`

	Format              SBOMFormat `json:"sbom_type"`
	IncludeDependencies bool       `json:"include_dependencies"`
	TeamIsTopLevel      bool       `json:"team_top_level"`
}

// Params converts an SBOMExportOptions object into a URL param object for use in making an API request
func (options SBOMExportOptions) Params() url.Values {
	params := url.Values{}
	params.Set("sbom_type", string(options.Format))
	params.Set("include_dependencies", strconv.FormatBool(options.IncludeDependencies))
	params.Set("team_top_level", strconv.FormatBool(options.TeamIsTopLevel))

	return params
}

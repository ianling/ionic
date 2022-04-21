package reports

import (
	"net/url"
	"strconv"
)

const (
	// ReportExportSBOMEndpoint is the endpoint for generating SBOMs
	ReportExportSBOMEndpoint = "v1/report/getSBOM"
)

// Standard is a string enum of supported SBOM standards
type Standard string

// Encoding is a string enum of supported file/data encoding standards
type Encoding string

const (
	// StandardUnknown denotes that the Standard is unknown
	StandardUnknown Standard = "unknown"
	// EncodingUnknown denotes that the Encoding is unknown
	EncodingUnknown Encoding = "unknown"

	// StandardCycloneDX indicates that the SBOM adheres to the CycloneDX standard
	StandardCycloneDX Standard = "CycloneDX"
	// StandardIonChannel indicates that the SBOM adheres to the Ion Channel standard
	StandardIonChannel Standard = "IonChannel"
	// StandardSPDX indicates that the SBOM adheres to the SPDX standard
	StandardSPDX Standard = "SPDX"

	// EncodingCSV indicates that the SBOM is stored in CSV format
	EncodingCSV Encoding = "CSV"
	// EncodingJSON indicates that the SBOM is stored in JSON format
	EncodingJSON Encoding = "JSON"
	// EncodingTagValue indicates that the SBOM is stored in tag-value format
	EncodingTagValue Encoding = "tag-value"
	// EncodingXLSX indicates that the SBOM is stored in XLSX format
	EncodingXLSX Encoding = "XLSX"
	// EncodingXML indicates that the SBOM is stored in XML format
	EncodingXML Encoding = "XML"
	// EncodingYAML indicates that the SBOM is stored in YAML format
	EncodingYAML Encoding = "YAML"
)

// SBOMExportOptions represents all the different settings a user can specify for how the SBOM is exported.
// Specify only one of the following data sources to generate the SBOM from:
//  * ProjectIDs (a slice of one or more project IDs)
//  * TeamID (the ID of a team containing one or more projects; this will use all the team's projects)
//  * SoftwareListID (the ID of a software list containing one or more components; this will use all the list's components)
//  * ComponentIDs (a slice of one or more software list component IDs)
// Format (required) specifies which format/standard the SBOM will be exported in.
// IncludeDependencies will include all the direct and transitive dependencies of each item in the SBOM if true,
// or exclude all the dependencies if false, leaving only the items themselves.
// TeamIsTopLevel applies only to SBOMs generated using the TeamID field. If true, the top-level item in the SBOM's
// hierarchy will be the team. If false, all the team's projects will be on the top level of the hierarchy.
type SBOMExportOptions struct {
	ProjectIDs     []string `json:"ids"`
	TeamID         string   `json:"team_id"`
	SoftwareListID string   `json:"sbom_id"`
	ComponentIDs   []string `json:"sbom_entry_ids"`

	Standard            Standard `json:"sbom_type"`
	Encoding            Encoding `json:"encoding"`
	IncludeDependencies bool     `json:"include_dependencies"`
	TeamIsTopLevel      bool     `json:"team_top_level"`
}

// Params converts an SBOMExportOptions object into a URL param object for use in making an API request
func (options SBOMExportOptions) Params() url.Values {
	params := url.Values{}
	params.Set("sbom_type", string(options.Standard))
	params.Set("encoding", string(options.Encoding))
	params.Set("include_dependencies", strconv.FormatBool(options.IncludeDependencies))
	params.Set("team_top_level", strconv.FormatBool(options.TeamIsTopLevel))

	return params
}

package scans

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ion-channel/ionic/dependencies"
	"github.com/ion-channel/ionic/risk"
	"github.com/ion-channel/ionic/secrets"
	"github.com/ion-channel/ionic/vulnerabilities"
)

// UntranslatedResults represents a result of a specific type that has not been
// translated for use in reports
type UntranslatedResults struct {
	AboutYML                *AboutYMLResults                `json:"about_yml,omitempty"`
	Buildsystem             *BuildsystemResults             `json:"buildsystems,omitempty"`
	Community               *CommunityResults               `json:"community,omitempty"`
	Coverage                *CoverageResults                `json:"coverage,omitempty"`
	Dependency              *DependencyResults              `json:"dependency,omitempty"`
	Difference              *DifferenceResults              `json:"difference,omitempty"`
	Ecosystem               *EcosystemResults               `json:"ecosystems,omitempty"`
	ExternalVulnerabilities *ExternalVulnerabilitiesResults `json:"external_vulnerability,omitempty"`
	License                 *LicenseResults                 `json:"license,omitempty"`
	Virus                   *VirusResults                   `json:"clamav,omitempty"`
	VirusDetails            *ClamavDetails                  `json:"clam_av_details,omitempty"`
	Vulnerability           *VulnerabilityResults           `json:"vulnerabilities,omitempty"`
	Secret                  *SecretResults                  `json:"secrets,omitempty"`
	Risk                    *RiskResults                    `json:"risk,omitempty"`
	Metrics                 *MetricsResults                 `json:"metrics,omitempty"`
}

// Translate moves information from the particular sub-struct, IE
// AboutYMLResults or LicenseResults into a generic, Data struct
func (u *UntranslatedResults) Translate() *TranslatedResults {
	var tr TranslatedResults
	// There is an argument to be made that the following "if" clauses
	// could be simplified with introspection since they all do
	// basically the same thing. I've (dmiles) chosen to writ it all
	// out in the name of explicit, easily-readable code.
	if u.AboutYML != nil {
		tr.Type = "about_yml"
		tr.Data = *u.AboutYML
	}
	if u.Buildsystem != nil {
		tr.Type = "buildsystems"
		tr.Data = *u.Buildsystem
	}
	if u.Community != nil {
		tr.Type = "community"
		tr.Data = *u.Community
	}
	if u.Coverage != nil {
		tr.Type = "coverage"
		tr.Data = *u.Coverage
	}
	if u.Dependency != nil {
		tr.Type = "dependency"
		tr.Data = *u.Dependency
	}
	if u.Difference != nil {
		tr.Type = "difference"
		tr.Data = *u.Difference
	}
	if u.Ecosystem != nil {
		tr.Type = "ecosystems"
		tr.Data = *u.Ecosystem
	}
	if u.ExternalVulnerabilities != nil {
		tr.Type = "external_vulnerability"
		tr.Data = *u.ExternalVulnerabilities
	}
	if u.License != nil {
		tr.Type = "license"
		tr.Data = *u.License
	}
	if u.Metrics != nil {
		tr.Type = "metrics"
		tr.Data = *u.Metrics
	}
	if u.Risk != nil {
		tr.Type = "risk"
		tr.Data = *u.Risk
	}
	if u.Secret != nil {
		tr.Type = "secrets"
		tr.Data = *u.Secret
	}
	if u.Virus != nil {
		tr.Type = "virus"
		u.Virus.ClamavDetails = *u.VirusDetails
		tr.Data = *u.Virus
	}
	if u.Vulnerability != nil {
		tr.Type = "vulnerability"
		tr.Data = *u.Vulnerability
	}
	return &tr
}

// TranslatedResults represents a result of a specific type that has been
// translated for use in reports
type TranslatedResults struct {
	Type string      `json:"type" xml:"type"`
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

type translatedResults struct {
	Type    string          `json:"type"`
	RawData json.RawMessage `json:"data"`
}

// UnmarshalJSON is a custom JSON unmarshaller implementation for the standard
// go json package to know how to properly interpret ScanSummaryResults from
// JSON.
func (r *TranslatedResults) UnmarshalJSON(b []byte) error {
	var tr translatedResults
	err := json.Unmarshal(b, &tr)
	if err != nil {
		return err
	}

	r.Type = tr.Type

	switch strings.ToLower(tr.Type) {
	case "about_yml":
		var a AboutYMLResults
		err := json.Unmarshal(tr.RawData, &a)
		if err != nil {
			return fmt.Errorf("failed to unmarshall about yml results: %v", err)
		}

		r.Data = a
	case "buildsystems":
		var b BuildsystemResults
		err := json.Unmarshal(tr.RawData, &b)
		if err != nil {
			return fmt.Errorf("failed to unmarshall buildsystems results: %v", err)
		}

		r.Data = b
	case "community":
		var c CommunityResults
		err := json.Unmarshal(tr.RawData, &c)
		if err != nil {
			// Note: Could be a slice, needs to be fixed
			if strings.Contains(err.Error(), "cannot unmarshal array") {
				var sliceOfCommunityResults []CommunityResults
				err := json.Unmarshal(tr.RawData, &sliceOfCommunityResults)
				if err == nil {
					c = sliceOfCommunityResults[0]
					break
				}
			}
			return fmt.Errorf("failed to unmarshall community results: %v", err)
		}

		r.Data = c
	case "coverage", "external_coverage":
		var c CoverageResults
		err := json.Unmarshal(tr.RawData, &c)
		if err != nil {
			return fmt.Errorf("failed to unmarshall coverage results: %v", err)
		}

		r.Data = c
	case "dependency":
		var d DependencyResults
		err := json.Unmarshal(tr.RawData, &d)
		if err != nil {
			return fmt.Errorf("failed to unmarshall dependency results: %v", err)
		}

		r.Data = d
	case "ecosystems":
		var e EcosystemResults
		err := json.Unmarshal(tr.RawData, &e)
		if err != nil {
			return fmt.Errorf("failed to unmarshall ecosystems results: %v", err)
		}

		r.Data = e
	case "license":
		var l LicenseResults
		err := json.Unmarshal(tr.RawData, &l)
		if err != nil {
			return fmt.Errorf("failed to unmarshall license results: %v", err)
		}

		r.Data = l
	case "metrics":
		var b MetricsResults
		err := json.Unmarshal(tr.RawData, &b)
		if err != nil {
			return fmt.Errorf("failed to unmarshall metrics results: %v", err)
		}

		r.Data = b
	case "risk":
		var b RiskResults
		err := json.Unmarshal(tr.RawData, &b)
		if err != nil {
			return fmt.Errorf("failed to unmarshall risk results: %v", err)
		}

		r.Data = b
	case "secrets":
		var b SecretResults
		err := json.Unmarshal(tr.RawData, &b)
		if err != nil {
			return fmt.Errorf("failed to unmarshall secrets results: %v", err)
		}

		r.Data = b
	case "virus", "clamav":
		var v VirusResults
		err := json.Unmarshal(tr.RawData, &v)
		if err != nil {
			return fmt.Errorf("failed to unmarshall virus results: %v", err)
		}

		r.Data = v
	case "vulnerability":
		var v VulnerabilityResults
		err := json.Unmarshal(tr.RawData, &v)
		if err != nil {
			return fmt.Errorf("failed to unmarshall vulnerability results: %v", err)
		}

		r.Data = v
	case "external_vulnerability":
		var v ExternalVulnerabilitiesResults
		err := json.Unmarshal(tr.RawData, &v)
		if err != nil {
			return fmt.Errorf("failed to unmarshall external vulnerabilities results: %v", err)
		}

		r.Data = v
	case "difference":
		var v DifferenceResults
		err := json.Unmarshal(tr.RawData, &v)
		if err != nil {
			return fmt.Errorf("failed to unmarshall difference results: %v", err)
		}

		r.Data = v
	default:
		return fmt.Errorf("unsupported results type found: %v", tr.Type)
	}

	return nil
}

// UnmarshalJSON is a custom JSON unmarshaller implementation for the standard
// go json package to know how to properly interpret ScanSummaryResults from
// JSON.
func (u *UntranslatedResults) UnmarshalJSON(b []byte) error {
	// first look for results in the proper translated format
	// e.g. CommunityResults
	tr := &translatedResults{}
	err := json.Unmarshal(b, tr)
	if err != nil {
		// we have received invalid stringified json
		return fmt.Errorf("unable to unmarshal json")
	}

	// if there is a type and it is `community`
	// parse the data out
	if tr.Type == "community" {
		c := &CommunityResults{}
		err = json.Unmarshal(tr.RawData, c)
		if err != nil {
			return err
		}
		u.Community = c
		return nil
	}

	// it is not translated and not community
	// ur2 is required to keep the parser from
	// recursing here
	type ur2 UntranslatedResults
	err = json.Unmarshal(b, (*ur2)(u))
	if err != nil {
		// we have received invalid stringified json
		return fmt.Errorf("unable to unmarshal json - %v", err.Error())
	}

	return nil
}

// AboutYMLResults represents the data collected from the AboutYML scan.  It
// includes a message and whether or not the About YML file found was valid or
// not.
type AboutYMLResults struct {
	Message string `json:"message" xml:"message"`
	Valid   bool   `json:"valid" xml:"valid"`
	Content string `json:"content" xml:"content"`
}

// Compiler represents the data for individual compilers or interpreters found
type Compiler struct {
	Name    string `json:"name" xml:"name"`
	Version string `json:"version" xml:"version"`
}

// Image represents the data for individual docker images found
type Image struct {
	Name    string `json:"name" xml:"name"`
	Version string `json:"version" xml:"version"`
}

// Dockerfile represents the data collected from a Dockerfile
type Dockerfile struct {
	Images       []Image                   `json:"images" xml:"images"`
	Dependencies []dependencies.Dependency `json:"dependencies" xml:"dependencies"`
}

// BuildsystemResults represents the data collected from an buildsystems scan.  It
// include the name and version of any compiler found
type BuildsystemResults struct {
	Compilers  []Compiler `json:"compilers" xml:"compilers"`
	Dockerfile Dockerfile `json:"docker_file" xml:"docker_file"`
}

// CommunityResults represents the data collected from a community scan.  It
// represents all known data regarding the open community of a software project
type CommunityResults struct {
	CommittersTotalCount int       `json:"committers_total_count" xml:"committers_total_count"`
	Name                 string    `json:"name" xml:"name"`
	URL                  string    `json:"url" xml:"url"`
	CommitsLastAt        time.Time `json:"commits_last_at" xml:"commits_last_at"`
	OldNames             []string  `json:"old_names" xml:"old_names"`
	StarsTotalCount      int       `json:"stars_total_count" xml:"stars_total_count"`
	NameChanged          bool      `json:"name_changed" xml:"name_changed"`
}

// CoverageResults represents the data collected from a code coverage scan.  It
// includes the value of the code coverage seen for the project.
type CoverageResults struct {
	Value float64 `json:"value" xml:"value"`
}

// Dependency represents data for an individual requirement resolution
type Dependency struct {
	LatestVersion string          `json:"latest_version" xml:"latest_version"`
	Org           string          `json:"org" xml:"org"`
	Name          string          `json:"name" xml:"name"`
	Type          string          `json:"type" xml:"type"`
	Package       string          `json:"package" xml:"package"`
	Version       string          `json:"version" xml:"version"`
	Scope         string          `json:"scope" xml:"scope"`
	Requirement   string          `json:"requirement" xml:"requirement"`
	File          string          `json:"file" xml:"file"`
	DepMeta       *DependencyMeta `json:"dependency_counts,omitempty" xml:"dependency_counts"`
	OutdatedMeta  *OutdatedMeta   `json:"outdated_version,omitempty" xml:"outdated_version"`
	Dependencies  []Dependency    `json:"dependencies" xml:"dependencies"`
}

// OutdatedMeta is used to represent the number of versions behind a dependcy is
type OutdatedMeta struct {
	MajorBehind int `json:"major_behind" xml:"major_behind"`
	MinorBehind int `json:"minor_behind" xml:"minor_behind"`
	PatchBehind int `json:"patch_behind" xml:"patch_behind"`
}

// DependencyMeta represents data for a summary of all dependencies resolved
type DependencyMeta struct {
	FirstDegreeCount     int `json:"first_degree_count" xml:"first_degree_count"`
	NoVersionCount       int `json:"no_version_count" xml:"no_version_count"`
	TotalUniqueCount     int `json:"total_unique_count" xml:"total_unique_count"`
	UpdateAvailableCount int `json:"update_available_count" xml:"update_available_count"`
	VulnerableCount      int `json:"vulnerable_count" xml:"vulnerable_count"`
}

// DependencyResults represents the data collected from a dependency scan.  It
// includes a list of the dependencies seen and meta data counts about those
// dependencies seen.
type DependencyResults struct {
	Dependencies []Dependency   `json:"dependencies" xml:"dependencies"`
	Meta         DependencyMeta `json:"meta" xml:"meta"`
}

// DifferenceResults represents the checksum of a project.  It includes a checksum
// and flag indicating if there was a difference detected within that last 5 scans
type DifferenceResults struct {
	Checksum   string `json:"checksum" xml:"checksum"`
	Difference bool   `json:"difference" xml:"difference"`
}

// EcosystemResults represents the data collected from an ecosystems scan.  It
// include the name of the ecosystem and the number of lines seen for the given
// ecosystem.
type EcosystemResults struct {
	Ecosystems map[string]int `json:"ecosystems" xml:"ecosystems"`
}

// MarshalJSON meets the marshaller interface to custom wrangle an ecosystem
// result into the json shape
func (e EcosystemResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Ecosystems)
}

// UnmarshalJSON meets the unmarshaller interface to custom wrangle the
// ecosystem scan into an ecosystem result
func (e *EcosystemResults) UnmarshalJSON(b []byte) error {
	var m map[string]int
	err := json.Unmarshal(b, &m)
	if err != nil {
		return fmt.Errorf("failed to unmarshal ecosystem result: %v", err.Error())
	}

	e.Ecosystems = m
	return nil
}

// ExternalVulnerabilitiesResults represents the data collected from an external
// vulnerability scan.  It includes the number of each vulnerability criticality
// seen within the project.
type ExternalVulnerabilitiesResults struct {
	Critical int `json:"critical" xml:"critical"`
	High     int `json:"high" xml:"high"`
	Medium   int `json:"medium" xml:"medium"`
	Low      int `json:"low" xml:"low"`
}

// LicenseResults represents the data collected from a license scan.  It
// includes the name and type of each license seen within the project.
type LicenseResults struct {
	*License `json:"license" xml:"license"`
}

// License represents a name and slice of types of licenses seen in a given file
type License struct {
	Name string        `json:"name" xml:"name"`
	Type []LicenseType `json:"type" xml:"type"`
}

// LicenseType represents a type of license such as MIT, Apache 2.0, etc
type LicenseType struct {
	Name       string  `json:"name" xml:"name"`
	Confidence float32 `json:"confidence"`
}

// FileNotes contains data related to file discoveries
type FileNotes map[string][]string

// ClamavDetails contains data related to the virus scan engine
type ClamavDetails struct {
	ClamavVersion   string `json:"clamav_version" xml:"clamav_version"`
	ClamavDbVersion string `json:"clamav_db_version" xml:"clamav_db_version"`
}

// MetricsResults is a slice of
type MetricsResults struct {
	Metrics risk.Metrics `json:"metrics" xml:"metrics"`
}

// MarshalJSON meets the marshaller interface to custom wrangle a risk
// result into the json shape
func (e MetricsResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Metrics)
}

// UnmarshalJSON meets the unmarshaller interface to custom wrangle the
// risk scan into an risk result
func (e *MetricsResults) UnmarshalJSON(b []byte) error {
	var s risk.Metrics
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal secrets result: %v", err.Error())
	}

	e.Metrics = s
	return nil
}

// RiskResults is a slice of
type RiskResults struct {
	Risk risk.EntityOverview `json:"risk" xml:"risk"`
}

// MarshalJSON meets the marshaller interface to custom wrangle a risk
// result into the json shape
func (e RiskResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Risk)
}

// UnmarshalJSON meets the unmarshaller interface to custom wrangle the
// risk scan into an risk result
func (e *RiskResults) UnmarshalJSON(b []byte) error {
	var s risk.EntityOverview
	err := json.Unmarshal(b, &s)
	if err != nil {
		s = risk.EntityOverview{}
	}

	e.Risk = s
	return nil
}

// Secret derived struct for results specific data
type Secret struct {
	secrets.Secret
	File string `json:"file" xml:"file"`
}

// SecretResults contains secrets finding data
type SecretResults struct {
	Secrets []Secret `json:"secrets" xml:"secrets"`
}

// MarshalJSON meets the marshaller interface to custom wrangle an ecosystem
// result into the json shape
func (e SecretResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Secrets)
}

// UnmarshalJSON meets the unmarshaller interface to custom wrangle the
// ecosystem scan into an ecosystem result
func (e *SecretResults) UnmarshalJSON(b []byte) error {
	var s []Secret
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal secrets result: %v", err.Error())
	}

	e.Secrets = s
	return nil
}

// VirusResults represents the data collected from a virus scan.  It includes
// information of the viruses seen and the virus scanner used.
type VirusResults struct {
	KnownViruses       int           `json:"known_viruses" xml:"known_viruses"`
	EngineVersion      string        `json:"engine_version" xml:"engine_version"`
	ScannedDirectories int           `json:"scanned_directories" xml:"scanned_directories"`
	ScannedFiles       int           `json:"scanned_files" xml:"scanned_files"`
	InfectedFiles      int           `json:"infected_files" xml:"infected_files"`
	DataScanned        string        `json:"data_scanned" xml:"data_scanned"`
	DataRead           string        `json:"data_read" xml:"data_read"`
	Time               string        `json:"time" xml:"time"`
	FileNotes          FileNotes     `json:"file_notes" xml:"file_notes"`
	ClamavDetails      ClamavDetails `json:"clam_av_details" xml:"clam_av_details"`
}

//VulnerabilityResults represents the data collected from a vulnerability scan.  It includes
// information of the vulnerabilities seen.
type VulnerabilityResults struct {
	Vulnerabilities []VulnerabilityResultsProduct `json:"vulnerabilities" xml:"vulnerabilities"`
	Meta            struct {
		VulnerabilityCount int    `json:"vulnerability_count" xml:"vulnerability_count"`
		ResolvedTo         string `json:"resolved_to" xml:"resolved_to"`
	} `json:"meta" xml:"meta"`
}

// VulnerabilityResultsProduct represents the data about a product collected from
// a vulnerability scan.  Vulnerabilities are linked to products.
type VulnerabilityResultsProduct struct {
	ID              int                                 `json:"id" xml:"id"`
	ExternalID      string                              `json:"external_id" xml:"external_id"`
	SourceID        int                                 `json:"source_id" xml:"source_id"`
	Title           string                              `json:"title" xml:"title"`
	Name            string                              `json:"name" xml:"name"`
	Org             string                              `json:"org" xml:"org"`
	Version         string                              `json:"version" xml:"version"`
	Up              interface{}                         `json:"up" xml:"up"`
	Edition         interface{}                         `json:"edition" xml:"edition"`
	Aliases         []string                            `json:"aliases" xml:"aliases"`
	CreatedAt       time.Time                           `json:"created_at" xml:"created_at"`
	UpdatedAt       time.Time                           `json:"updated_at" xml:"updated_at"`
	References      interface{}                         `json:"references" xml:"references"`
	Part            interface{}                         `json:"part" xml:"part"`
	Language        interface{}                         `json:"language" xml:"language"`
	Vulnerabilities []VulnerabilityResultsVulnerability `json:"vulnerabilities" xml:"vulnerabilities"`
	Query           Dependency                          `json:"query" xml:"query"`
}

// VulnerabilityResultsVulnerability wrapper
type VulnerabilityResultsVulnerability struct {
	vulnerabilities.Vulnerability
	Dependencies []VulnerabilityResultsProduct `json:"dependencies" xml:"dependencies"`
}

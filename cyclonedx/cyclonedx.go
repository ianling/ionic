package cyclonedx

import (
	"fmt"
	"github.com/CycloneDX/cyclonedx-go"
	"github.com/google/uuid"
	"github.com/ion-channel/ionic/aliases"
	"github.com/ion-channel/ionic/projects"
	"github.com/ion-channel/ionic/util"
	"strings"
)

// FromJSONString takes a CycloneDX SoftwareList in JSON format, as a string, and returns a BOM object
func FromJSONString(sbomContents string) (*cyclonedx.BOM, error) {
	bom, err := fromString(sbomContents, cyclonedx.BOMFileFormatJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CycloneDX SoftwareList in JSON format: %s", err.Error())
	}

	return bom, nil
}

// FromXMLString takes a CycloneDX SoftwareList in XML format, as a string, and returns a BOM object
func FromXMLString(sbomContents string) (*cyclonedx.BOM, error) {
	bom, err := fromString(sbomContents, cyclonedx.BOMFileFormatXML)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CycloneDX SoftwareList in XML format: %s", err.Error())
	}

	return bom, nil
}

func fromString(sbomContents string, fileType cyclonedx.BOMFileFormat) (*cyclonedx.BOM, error) {
	// need a stream-like object to decode
	sbomContentsStream := strings.NewReader(sbomContents)

	bom := new(cyclonedx.BOM)
	decoder := cyclonedx.NewBOMDecoder(sbomContentsStream, fileType)
	if err := decoder.Decode(bom); err != nil {
		return nil, err
	}

	return bom, nil
}

func projectFromComponent(component cyclonedx.Component) projects.Project {
	var projectAliases []aliases.Alias
	if len(component.Name) > 0 || len(component.Publisher) > 0 || len(component.Version) > 0 {
		projectAliases = []aliases.Alias{{
			Name:    component.Name,
			Org:     component.Publisher,
			Version: component.Version,
		}}
	}

	var source, branch, projectType string
	projectType = "source_unavailable" // default to source_unavailable, if we find a source we'll use it

	if component.ExternalReferences != nil && len(*component.ExternalReferences) > 0 {
		for _, externalRef := range *component.ExternalReferences {
			if externalRef.Type == cyclonedx.ERTypeVCS {
				projectType = "git"
				source, branch = util.ParseGitURL(externalRef.URL)
			}
		}
	}

	tempUUID := uuid.New().String()
	project := projects.Project{
		ID:      &tempUUID,
		Name:    &component.Name,
		Type:    &projectType,
		Source:  &source,
		Branch:  &branch,
		Active:  true,
		Monitor: true,
		Aliases: projectAliases,
		PURL:    component.PackageURL,
		CPE:     component.CPE,
	}

	return project
}

// ProjectsFromCycloneDX parses components from a CycloneDX SoftwareList into Projects.
func ProjectsFromCycloneDX(sbom *cyclonedx.BOM, includeDependencies bool) ([]projects.Project, error) {
	// pre-allocate a slice of projects
	estimatedSbomLength := 1
	if sbom.Components != nil {
		estimatedSbomLength += len(*sbom.Components)
	}

	sbomProjects := make([]projects.Project, 0, estimatedSbomLength)

	if sbom.Metadata != nil && sbom.Metadata.Component != nil {
		sbomProjects = append(sbomProjects, projectFromComponent(*sbom.Metadata.Component))
	}

	if includeDependencies && estimatedSbomLength > 1 {
		// get all the components in the SoftwareList
		for _, component := range *sbom.Components {
			project := projectFromComponent(component)

			// don't include duplicates
			// (e.g. if two dependencies share a transitive dependency, only count the transitive dependency once)
			if projects.ProjectSliceContains(sbomProjects, project) {
				continue
			}

			sbomProjects = append(sbomProjects, project)
		}
	}

	return sbomProjects, nil
}

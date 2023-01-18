package spdx

import (
	"fmt"
	"github.com/ion-channel/ionic/util"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/ion-channel/ionic/aliases"
	"github.com/ion-channel/ionic/projects"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/spdxlib"
)

type packageInfo struct {
	Name             string
	Version          string
	DownloadLocation string
	Description      string
	Organization     string
	CPE              string
	PURL             string
}

// packageInfoFromPackage takes either an spdx.Package2_1 or spdx.Package2_2 and returns a packageInfo object.
// This is used to convert SPDX packages to version-agnostic representations of the data we need.
func packageInfoFromPackage(spdxPackage interface{}) packageInfo {
	var name, version, downloadLocation, description, organization, cpe, purl string

	switch spdxPackage.(type) {
	case spdx.Package2_1:
		packageTyped := spdxPackage.(spdx.Package2_1)
		name = packageTyped.PackageName
		version = packageTyped.PackageVersion
		downloadLocation = packageTyped.PackageDownloadLocation
		description = packageTyped.PackageDescription

		if packageTyped.PackageSupplier != nil && packageTyped.PackageSupplier.SupplierType == "Organization" {
			organization = packageTyped.PackageSupplier.Supplier
		}

		for _, externalRef := range packageTyped.PackageExternalReferences {
			if externalRef.Category == "SECURITY" && externalRefIsCPE(externalRef.RefType) {
				cpe = externalRef.Locator
			} else if externalRef.Category == "PACKAGE-MANAGER" && externalRefIsPURL(externalRef.RefType) {
				purl = externalRef.Locator
			}
		}
	case spdx.Package2_2:
		packageTyped := spdxPackage.(spdx.Package2_2)
		name = packageTyped.PackageName
		version = packageTyped.PackageVersion
		downloadLocation = packageTyped.PackageDownloadLocation
		description = packageTyped.PackageDescription

		if packageTyped.PackageSupplier != nil && packageTyped.PackageSupplier.SupplierType == "Organization" {
			organization = packageTyped.PackageSupplier.Supplier
		}

		for _, externalRef := range packageTyped.PackageExternalReferences {
			if externalRefIsCPE(externalRef.RefType) {
				cpe = externalRef.Locator
			} else if externalRefIsPURL(externalRef.RefType) {
				purl = externalRef.Locator
			}
		}
	}

	return packageInfo{
		Name:             name,
		Version:          version,
		DownloadLocation: downloadLocation,
		Description:      description,
		Organization:     organization,
		CPE:              cpe,
		PURL:             purl,
	}
}

// ProjectsFromSPDX parses packages from an SPDX Document (v2.1 or v2.2) into Projects.
// The given document must be of the type *spdx.Document2_1 or *spdx.Document2_2.
// A package in the document must have a valid, resolveable PackageDownloadLocation in order to create a project
func ProjectsFromSPDX(doc interface{}, includeDependencies bool) ([]projects.Project, error) {
	// use a SPDX-version-agnostic container for tracking package info
	packageInfos := []packageInfo{}

	switch doc.(type) {
	case *spdx.Document2_1:
		docTyped := doc.(*spdx.Document2_1)

		if includeDependencies {
			// just get all of the packages
			for _, spdxPackage := range docTyped.Packages {
				packageInfos = append(packageInfos, packageInfoFromPackage(*spdxPackage))
			}
		} else {
			// get only the top-level packages
			topLevelPkgIDs, err := spdxlib.GetDescribedPackageIDs2_1(docTyped)
			if err != nil {
				return nil, fmt.Errorf("failed to retrieve described packages from SPDX 2.1 document: %s", err.Error())
			}

			for _, spdxPackage := range docTyped.Packages {
				var isTopLevelPackage bool
				for _, pkgID := range topLevelPkgIDs {
					if pkgID == spdxPackage.PackageSPDXIdentifier {
						isTopLevelPackage = true
						break
					}
				}

				if !isTopLevelPackage {
					continue
				}

				packageInfos = append(packageInfos, packageInfoFromPackage(*spdxPackage))
			}
		}
	case *spdx.Document2_2:
		docTyped := doc.(*spdx.Document2_2)

		if includeDependencies {
			// just get all of the packages
			for _, spdxPackage := range docTyped.Packages {
				packageInfos = append(packageInfos, packageInfoFromPackage(*spdxPackage))
			}
		} else {
			// get only the top-level packages
			topLevelPkgIDs, err := spdxlib.GetDescribedPackageIDs2_2(docTyped)
			if err != nil {
				return nil, fmt.Errorf("failed to retrieve described packages from SPDX 2.2 document: %s", err.Error())
			}

			for _, spdxPackage := range docTyped.Packages {
				var isTopLevelPackage bool
				for _, pkgID := range topLevelPkgIDs {
					if pkgID == spdxPackage.PackageSPDXIdentifier {
						isTopLevelPackage = true
						break
					}
				}

				if !isTopLevelPackage {
					continue
				}

				packageInfos = append(packageInfos, packageInfoFromPackage(*spdxPackage))
			}
		}
	default:
		return nil, fmt.Errorf("wrong document type given, need *spdx.Document2_1 or *spdx.Document2_2")
	}

	projs := []projects.Project{}
	for ii := range packageInfos {
		pkg := packageInfos[ii]
		// info we need to parse out of the SoftwareList
		var ptype, source, branch string

		tmpID := uuid.New().String()

		if pkg.DownloadLocation == "" || pkg.DownloadLocation == "NOASSERTION" || pkg.DownloadLocation == "NONE" {
			ptype = "source_unavailable"
		} else if strings.Contains(pkg.DownloadLocation, "git") {
			ptype = "git"

			// SPDX spec says that git URLs can look like "git+https://github.com/..."
			// we need to strip off the "git+"
			if strings.Index(pkg.DownloadLocation, "git+") == 0 {
				source = pkg.DownloadLocation[4:]
			} else {
				source = pkg.DownloadLocation
			}

			source, branch = util.ParseGitURL(source)
		} else {
			source = pkg.DownloadLocation
			ptype = "artifact"
		}

		proj := projects.Project{
			ID:          &tmpID,
			Branch:      &branch,
			Description: &pkg.Description,
			Type:        &ptype,
			Source:      &source,
			Name:        &pkg.Name,
			Active:      true,
			Monitor:     true,
			CPE:         pkg.CPE,
			PURL:        pkg.PURL,
		}

		// check if version, org, or name are not empty strings
		if len(pkg.Version) > 0 || len(pkg.Organization) > 0 || len(pkg.Name) > 0 {
			v := pkg.Version
			if pkg.Version == "NOASSERTION" {
				v = ""
			}
			proj.Aliases = []aliases.Alias{{
				Name:    pkg.Name,
				Org:     pkg.Organization,
				Version: v,
			}}
		}

		// make sure we don't already have an equivalent project
		if projects.ProjectSliceContains(projs, proj) {
			continue
		}

		projs = append(projs, proj)

	}

	return projs, nil
}

// externalRefIsCPE returns true if the given external reference type refers to a CPE.
func externalRefIsCPE(externalRefType string) bool {
	return externalRefType == "cpe23type" ||
		externalRefType == "cpe22type" ||
		externalRefType == "http://spdx.org/rdf/references/cpe23Type" ||
		externalRefType == "http://spdx.org/rdf/references/cpe22Type"
}

// externalRefIsPURL returns true if the given external reference type refers to a PURL.
func externalRefIsPURL(externalRefType string) bool {
	return externalRefType == "purl" || externalRefType == "http://spdx.org/rdf/references/purl"
}

// Helper function to parse email from SPDX Creator info
// SPDX email comes in the form Creator: Person: My Name (myname@mail.com)
// returns empty string if no email is found
func parseCreatorEmail(creatorPersons []string) string {
	if len(creatorPersons) > 0 {
		re := regexp.MustCompile(`\((.*?)\)`)
		email := re.FindStringSubmatch(creatorPersons[0])
		if len(email) > 0 && email != nil {
			return email[1]
		}
	}
	return ""
}

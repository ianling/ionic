package dependencies

import (
	"time"
)

const (
	// GetLatestVersionForDependencyEndpoint - returns the latest single version for a dependency
	GetLatestVersionForDependencyEndpoint = "v1/dependency/getLatestVersionForDependency"
	// GetVersionsForDependencyEndpoint - returns the list of known versions for a dependency
	GetVersionsForDependencyEndpoint = "v1/dependency/getVersionsForDependency"
	// ResolveDependenciesInFileEndpoint - given a dependency file and ecosystem name returns the full tree of known dependencies
	ResolveDependenciesInFileEndpoint = "v1/dependency/resolveDependenciesInFile"
	// ResolveFromFileEndpoint - given a dependency file and ecosystem name returns the full tree of known dependencies.
	// Only supports Gemfile.lock.
	ResolveFromFileEndpoint = "v1/dependency/resolveFromFile"
	// ResolveDependencySearchEndpoint is a string representation of the current endpoint for searching dependencies
	ResolveDependencySearchEndpoint = "v1/dependency/search"
	// GetDependencyVersions is a string representation of the current endpoint for returns the list of known versions for a dependency
	GetDependencyVersions = "v1/dependency/getVersions"
)

// Dependency represents all the known information for a dependency object
// within the Ion Channel API
type Dependency struct {
	Name            string       `json:"name,omitempty"`
	Version         string       `json:"version"`
	LatestVersion   string       `json:"latest_version"`
	Org             string       `json:"org"`
	Type            string       `json:"type"`
	Package         string       `json:"package"`
	Scope           string       `json:"scope"`
	Requirement     string       `json:"requirement"`
	Dependencies    []Dependency `json:"dependencies"`
	Confidence      float32      `json:"confidence"`
	CreatedAt       time.Time    `json:"created_at,omitempty"`
	UpdatedAt       time.Time    `json:"updated_at,omitempty"`
	OutdatedVersion OutdatedMeta `json:"outdated_version"`
	Matches         []string     `json:"matches,omitempty" xml:"matches,omitempty"`
}

type Metrics struct {
	ID                     string    `json:"id"`
	DependenciesTotalCount int       `json:"dependencies_total_count"`
	License                string    `json:"license_yn"`
	OrgPackageCount        int       `json:"org_package_count"`
	PrevVersionCount       int       `json:"prev_version_count"`
	UpdatedAt              time.Time `json:"updated_at"`
}

// OutdatedMeta is used to represent the number of versions behind a dependency is
type OutdatedMeta struct {
	MajorBehind int `json:"major_behind" xml:"major_behind"`
	MinorBehind int `json:"minor_behind" xml:"minor_behind"`
	PatchBehind int `json:"patch_behind" xml:"patch_behind"`
}

// Meta represents all the known meta information for a dependency set
// within the Ion Channel API
type Meta struct {
	// {"first_degree_count":13,"no_version_count":0,"total_unique_count":62,"update_available_count":12}
	FirstDegreeCount     int `json:"first_degree_count"`
	NoVersionCount       int `json:"no_version_count"`
	TotalUniqueCount     int `json:"total_unique_count"`
	UpdateAvailableCount int `json:"update_available_count"`
}

// DependencyResolutionResponse represents all the known information
// for a dependency object within the Ion Channel API
type DependencyResolutionResponse struct {
	Dependencies []Dependency `json:"dependencies,omitempty"`
	Meta         Meta         `json:"meta"`
}

// DependencyResolutionRequest options for creating a resolution request
// for a dependency file of a ecosystem type
type DependencyResolutionRequest struct {
	Ecosystem string
	File      string
	Flatten   bool
}

package products

import (
	"time"
)

const (
	// GetProductEndpoint is a string representation of the current endpoint for getting product
	GetProductEndpoint = "v1/vulnerability/getProducts"
	// ProductSearchEndpoint is a string representation of the current endpoint for product search
	ProductSearchEndpoint = "v1/product/search"
	// ProductGetProductEndpoint is a string representation of the current endpoint for getting a product
	ProductGetProductEndpoint = "v1/product/getProduct"
	// GetProductVersionsEndpoint is a string representation of the current endpoint for getting a product's versions
	GetProductVersionsEndpoint = "v1/product/getProductVersions"
)

// Product represents a software product within the system for identification
// across multiple sources
type Product struct {
	ID                 int           `json:"id" xml:"id"`
	Name               string        `json:"name" xml:"name"`
	Org                string        `json:"org" xml:"org"`
	Version            string        `json:"version" xml:"version"`
	Up                 string        `json:"up" xml:"up"`
	Edition            string        `json:"edition" xml:"edition"`
	Aliases            interface{}   `json:"aliases" xml:"aliases"`
	CreatedAt          time.Time     `json:"created_at" xml:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at" xml:"updated_at"`
	Title              string        `json:"title" xml:"title"`
	References         []interface{} `json:"references" xml:"references"`
	Part               string        `json:"part" xml:"part"`
	Language           string        `json:"language" xml:"language"`
	ExternalID         string        `json:"external_id" xml:"external_id"`
	Sources            []Source      `json:"source" xml:"source"`
	Confidence         float64       `json:"confidence" xml:"confidence"`
	VulnerabilityCount int           `json:"vulnerability_count" xml:"vulnerability_count"`
	Mttr               *int64        `json:"mttr_seconds" xml:"mttr_seconds"`
	Vulnerabilities    []interface{} `json:"vulnerabilities" xml:"vulnerabilities"`
	Matches            []string      `json:"matches,omitempty" xml:"matches,omitempty"`
}

type Count struct {
	Count  int    `json:"count"`
	Source string `json:"source"`
}

type MonthlyCounts struct {
	Count []Count `json:"count"`
	Month string  `json:"month"`
}

type Metrics struct {
	ID                                      string          `json:"id"`
	VulnTotalCount                          []Count         `json:"vuln_total_count"`
	CriticalSeverityVulnTotalCount          []Count         `json:"critical_severity_vuln_total_count"`
	HighSeverityVulnTotalCount              []Count         `json:"high_severity_vuln_total_count"`
	MediumSeverityVulnTotalCount            []Count         `json:"medium_severity_vuln_total_count"`
	LowSeverityVulnTotalCount               []Count         `json:"low_severity_vuln_total_count"`
	VulnMonthlyCount                        []MonthlyCounts `json:"vuln_monthly_count"`
	CriticalSeverityVulnMonthlyCount        []MonthlyCounts `json:"critical_severity_vuln_monthly_count"`
	HighSeverityVulnMonthlyCount            []MonthlyCounts `json:"high_severity_vuln_monthly_count"`
	MediumSeverityVulnMonthlyCount          []MonthlyCounts `json:"medium_severity_vuln_monthly_count"`
	LowSeverityVulnMonthlyCount             []MonthlyCounts `json:"low_severity_vuln_monthly_count"`
	AverageSeverity                         float64         `json:"average_severity"`
	AverageVulnMonthlyCount                 float64         `json:"average_vuln_monthly_count"`
	AverageCriticalSeverityVulnMonthlyCount float64         `json:"average_critical_severity_vuln_monthly_count"`
	AverageHighSeverityVulnMonthlyCount     float64         `json:"average_high_severity_vuln_monthly_count"`
	AverageMediumSeverityVulnMonthlyCount   float64         `json:"average_medium_severity_vuln_monthly_count"`
	AverageLowSeverityVulnMonthlyCount      float64         `json:"average_low_severity_vuln_monthly_count"`
	OrgProductsCount                        float64         `json:"org_products_count"`
	PrevVersionCount                        float64         `json:"prev_version_count"`
	UpdatedAt                               time.Time       `json:"updated_at"`
}

// Source represents information about where the product data came from
type Source struct {
	ID           int       `json:"id" xml:"id"`
	Name         string    `json:"name" xml:"name"`
	Description  string    `json:"description" xml:"description"`
	CreatedAt    time.Time `json:"created_at" xml:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" xml:"updated_at"`
	Attribution  string    `json:"attribution" xml:"attribution"`
	License      string    `json:"license" xml:"license"`
	CopyrightURL string    `json:"copyright_url" xml:"copyright_url"`
}

// SoftwareEntity represents information about a product as well as
// other info, like Git repository, committer counts, etc
type SoftwareEntity struct {
	Product    *Product             `json:"product,omitempty" xml:"product"`
	Github     *Github              `json:"github,omitempty" xml:"github,omitempty"`
	Package    *Package             `json:"package,omitempty" xml:"package,omitempty"`
	Confidence float64              `json:"confidence" xml:"confidence"`
	Scores     []ProductSearchScore `json:"scores,omitempty" xml:"scores"`
}

// ProductSearchScore represents the TF;IDF score for a given search result
// and a given search term
type ProductSearchScore struct {
	Term  string  `json:"term" xml:"term"`
	Score float64 `json:"score" xml:"score"`
}

// Github represents information from Github about a given repository
type Github struct {
	URI            string `json:"uri" xml:"uri"`
	CommitterCount uint   `json:"committer_count" xml:"committer_count"`
}

// Package represents information about a package from one of
// our supported package management systems like pypi, npm or rubygems
type Package struct {
	Name    string `json:"name" xml:"name"`
	Version string `json:"version" xml:"version"`
	Type    string `json:"type" xml:"type"`
}

// ProductSearchQuery collects all the various searching options that
// the productSearchEndpoint supports for use in a POST request
type ProductSearchQuery struct {
	SearchType        string   `json:"search_type" xml:"search_type"`
	SearchStrategy    string   `json:"search_strategy" xml:"search_strategy"`
	ProductIdentifier string   `json:"product_identifier" xml:"product_identifier"`
	Version           string   `json:"version" xml:"version"`
	Vendor            string   `json:"vendor" xml:"vendor"`
	Terms             []string `json:"terms" xml:"terms"`
}

// IsValid checks some of the constraints on the ProductSearchQuery to
// help the programmer determine if productSearchEndpoint will accept it
func (p *ProductSearchQuery) IsValid() bool {
	if len(p.SearchStrategy) > 0 {
		if p.SearchType == "concatenated" || p.SearchType == "deconcatenated" {
			return true
		}
	}
	return false
}

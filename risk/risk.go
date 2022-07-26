package risk

import (
	"time"

	"github.com/ion-channel/ionic/community"
	"github.com/ion-channel/ionic/products"
)

const (
	// GetScoresEnpoint location of the getscores endpoint
	GetScoresEnpoint = "v1/score/getScores"

	// EcosystemScope constant with name of scope for ecosystem
	EcosystemScope = "ecosystem"

	// SupplyChainScope constant with name of scope for supply chain
	SupplyChainScope = "supply_chain"

	// TechnologyScope constant with name of scope for technology
	TechnologyScope = "technology"
)

type Metrics struct {
	ID      string   `json:"id"`
	Metrics []Metric `json:"metrics"`
}

type Metric struct {
	Name     string         `json:"name"`
	Value    interface{}    `json:"value"`
	Bindings []ScoreBinding `json:"bindings"`
}

type IntMetric struct {
	Metric
	Value int `json:"value"`
}

type FloatMetric struct {
	Metric
	Value float64 `json:"value"`
}

type DateMetric struct {
	Metric
	Value time.Time `json:"value"`
}

type MonthlyCountMetric struct {
	Metric
	Value []community.MonthlyCount `json:"value"`
}

type MonthlyFloatMetric struct {
	Metric
	Value []community.MonthlyFloat `json:"value"`
}

type SourceMonthlyCountMetric struct {
	Metric
	Value []products.MonthlyCounts `json:"value"`
}

type SourceCountMetric struct {
	Metric
	Value []products.Count `json:"value"`
}

// ScoreBinding a mapping from metric to which scope it falls into.
type ScoreBinding struct {
	Metric    string `json:"metric"`
	Scope     string `json:"scope"`
	Category  string `json:"category"`
	Attribute string `json:"attribute"`
	Source    string `json:"source"`
}

// Scores top level struct for modeling the score tree
type Scores struct {
	Name   string  `json:"name"`
	Value  float64 `json:"value"`
	Scopes []Scope `json:"scopes"`
}

// Scope second tier struct for modeling the score tree
type Scope struct {
	Name       string     `json:"name"`
	Value      float64    `json:"value"`
	Categories []Category `json:"-"`
}

// GetScope returns a Scope value based on the name supplied
func (s *Scores) GetScope(name string) *Scope {
	var scope *Scope
	for _, sco := range s.Scopes {
		if sco.Name == name {
			scope = &sco
			break
		}
	}

	if scope == nil {
		scope = &Scope{Name: name}
	}
	return scope
}

// Category third tier struct for modeling the score tree
type Category struct {
	Name       string      `json:"name"`
	Value      float64     `json:"value"`
	Attributes []Attribute `json:"-"`
}

// Attribute leaf tier struct for modeling the score tree
type Attribute struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type RiskTag struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type MetricMetadata struct {
	Name           string         `json:"name"`
	Definition     string         `json:"definition"`
	Bindings       []ScoreBinding `json:"bindings"`
	RiskTags       []RiskTag      `json:"risk_tags"`
	RelatedMetrics []string       `json:"related_metrics"`
	GraphYN        bool           `json:"graph_yn"`
}

// MetricPoint defines the data needed for points on a single risk point
type MetricPoint struct {
	Name   string `json:"name" xml:"name"`
	Points int    `json:"points" xml:"points"`
}

// MetricPoints defines the data needed for points on a single risk point
type MetricPoints struct {
	Metrics     []MetricPoint `json:"metrics" xml:"metrics"`
	ProcessedAt time.Time     `json:"processed_at" xml:"processed_at"`
}

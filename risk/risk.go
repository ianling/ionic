package risk

import (
	"encoding/json"
	"time"
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
	Name         string           `json:"name"`
	Bindings     []ScoreBinding   `json:"bindings"`
	Severity     string           `json:"severity"`
	SeverityRank int              `json:"severity_rank"`
	Value        *json.RawMessage `json:"value"`
	Type         string           `json:"type"`
	Sources      []string         `json:"sources"`
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

type RelatedMetric struct {
	Name         string `json:"name"`
	InternalName string `json:"internal_name"`
}

type MetricMetadata struct {
	Name           string          `json:"name"`
	InternalName   string          `json:"internal_name"`
	Definition     string          `json:"definition"`
	Scopes         []string        `json:"scopes"`
	RiskTags       []RiskTag       `json:"risk_tags"`
	RelatedMetrics []RelatedMetric `json:"related_metrics"`
	GraphYN        bool            `json:"graph_yn"`
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

type EntityOverview struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Org           string          `json:"org"`
	Version       string          `json:"version"`
	Sources       []EntitySource  `json:"sources"`
	Scores        Scores          `json:"score"`
	EntitySummary string          `json:"summary"`
	RiskTags      []EntityRiskTag `json:"risk_tags"`
}

type EntityRiskTag struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
}

type EntitySource struct {
	Type   string   `json:"type"`
	Source []string `json:"source"`
	ID     string   `json:"id"`
}

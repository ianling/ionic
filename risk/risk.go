package risk

import "time"

const (
	// GetScoresEnpoint location of the getscores endpoint
	GetScoresEnpoint = "v1/score/getScores"

	// EcosystemScope constant with name of scope for ecosystem
	EcosystemScope = "ecosystem"

	// SupplyChainScope constant with name of scope for supply chain
	SupplyChainScope = "supply chain"

	// TechnologyScope constant with name of scope for technology
	TechnologyScope = "technology"
)

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

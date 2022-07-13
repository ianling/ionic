package risk

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

type Metrics struct {
	ID      string   `json:"id"`
	Metrics []Metric `json:"metrics"`
}

type Metric struct {
	Name     string         `json:"name"`
	Value    interface{}    `json:"value"`
	Bindings []ScoreBinding `json:"bindings"`
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

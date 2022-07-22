package software_lists

type RiskScope struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Risk struct {
	Score  int         `json:"score"`
	Scopes []RiskScope `json:"scopes"`
}

type Compliance struct {
	Passing int `json:"passing"`
	Failing int `json:"failing"`
}

type Resolution struct {
	Resolved          int `json:"resolved"`
	PartiallyResolved int `json:"partially_resolved"`
	Unresolved        int `json:"unresolved"`
}

type Metrics struct {
	Risk       Risk       `json:"risk"`
	Compliance Compliance `json:"compliance"`
	Resolution Resolution `json:"resolution"`
}

package rulesets

import (
	"strings"
	"time"

	"github.com/ion-channel/ionic/scans"
)

const (
	// GetProjectHistoryEndpoint is a string representation of the current endpoint for getting a project's history
	GetProjectHistoryEndpoint = "v1/ruleset/getProjectHistory"
)

// AppliedRulesetSummary identifies the rule set applied to an analysis of a
// project and the result of their evaluation
type AppliedRulesetSummary struct {
	ProjectID             string                 `json:"project_id"`
	TeamID                string                 `json:"team_id"`
	AnalysisID            string                 `json:"analysis_id"`
	RulesetID             string                 `json:"ruleset_id"`
	RulesetName           string                 `json:"ruleset_name"`
	RuleEvaluationSummary *RuleEvaluationSummary `json:"rule_evaluation_summary"`
	CreatedAt             time.Time              `json:"created_at"`
	UpdatedAt             time.Time              `json:"updated_at"`
}

// SummarizeEvaluation returns the calculated risk and passing values for the
// AppliedRulsetSummary. Only if the RuleEvalutionSummary has passed, will it
// return low risk and passing.
func (ar *AppliedRulesetSummary) SummarizeEvaluation() (string, bool) {
	if ar.RuleEvaluationSummary != nil && strings.ToLower(ar.RuleEvaluationSummary.Summary) == "pass" {
		return "low", true
	}

	return "high", false
}

// RuleEvaluationSummary represents the ruleset and the scans that were
// evaluated with the ruleset
type RuleEvaluationSummary struct {
	RulesetName string             `json:"ruleset_name"`
	Summary     string             `json:"summary"`
	Risk        string             `json:"risk"`
	Passed      bool               `json:"passed"`
	Ruleresults []scans.Evaluation `json:"ruleresults"`
}

// ProjectPassFailHistory represents a summary of one day's analysis results for a project.
type ProjectPassFailHistory struct {
	TeamID        string    `json:"team_id"`
	ProjectID     string    `json:"project_id"`
	// AnalysisID is the ID of the last analysis ran on a particular date.
	AnalysisID    string    `json:"analysis_id"`
	// Status is true if the last analysis ran on a particular date was passing, otherwise it is false
	Status        bool      `json:"pass"`
	// FailCount is the total number of the project's analyses that failed on a particular date
	FailCount     int       `json:"fail_count"`
	// PassCount is the total number of the project's analyses that passed on a particular date
	PassCount     int       `json:"pass_count"`
	// CreatedAt is the date this data is about
	CreatedAt     time.Time `json:"created_at"`
	// StatusFlipped indicates whether or not the project's status flipped from passing to failing, or vice versa,
	// on this day, or if this day's status differs from the previous day's.
	StatusFlipped bool      `json:"status_flipped"`
}

// ProjectRulesetHistory represents history of a project's ruleset changing
type ProjectRulesetHistory struct {
	OldRulesetID   string    `json:"old_ruleset_id"`
	OldRulesetName string    `json:"old_ruleset_name"`
	NewRulesetID   string    `json:"new_ruleset_id"`
	NewRulesetName string    `json:"new_ruleset_name"`
	UserID         string    `json:"user_id"`
	UserName       string    `json:"user_name"`
	CreatedAt      time.Time `json:"created_at"`
}

// ProjectAudit represents a projects history agregated by date
type ProjectAudit struct {
	Date           time.Time               `json:"date"`
	PassFail       *ProjectPassFailHistory `json:"project_pass_fail,omitempty"`
	RulesetHistory []ProjectRulesetHistory `json:"ruleset_history,omitempty"`
}

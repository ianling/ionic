package rulesets

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ion-channel/ionic/rules"
)

const (
	// CreateRuleSetEndpoint is a string representation of the current endpoint for creating ruleset
	CreateRuleSetEndpoint = "v1/ruleset/createRuleset"
	// GetAppliedRuleSetEndpoint is a string representation of the current endpoint for getting applied ruleset
	GetAppliedRuleSetEndpoint = "v1/ruleset/getAppliedRulesetForProject"
	// GetBatchAppliedRulesetEndpoint is a string representation of the current endpoint for getting batched applied rulesets
	GetBatchAppliedRulesetEndpoint = "v1/ruleset/getAppliedRulesets"
	// GetRuleSetEndpoint is a string representation of the current endpoint for getting ruleset
	GetRuleSetEndpoint = "v1/ruleset/getRuleset"
	// GetRuleSetsEndpoint is a string representation of the current endpoint for getting rulesets (plural)
	GetRuleSetsEndpoint = "v1/ruleset/getRulesets"
	// GetDefaultRuleSetsEndpoint is the endpoint for getting all the global default rulesets
	GetDefaultRuleSetsEndpoint = "v1/ruleset/getDefaultRulesets"
	// RulesetsGetRulesEndpoint is a string representation of the current endpoint for getting rules.
	RulesetsGetRulesEndpoint = "v1/ruleset/getRules"
	// RulesetsGetRulesetNames is a string representation of the current endpoint for getting ruleset names.
	RulesetsGetRulesetNames = "v1/ruleset/getRulesetNames"
	// GetRulesetAnalysesStatuses is a string representation of the current endpoint for getting analyses statuses.
	GetRulesetAnalysesStatuses = "v1/ruleset/getAnalysesStatuses"
)

// AppliedRulesetRequest represents a request for an applied ruleset result
type AppliedRulesetRequest struct {
	ProjectID  string `json:"project_id"`
	TeamID     string `json:"team_id"`
	SummaryID  string `json:"summary_id"`
	AnalysisID string `json:"analysis_id"`
}

// CreateRuleSetOptions struct for creating a ruleset
type CreateRuleSetOptions struct {
	Name        string   `json:"name"`
	Description string   `json:"description" default:" "`
	TeamID      string   `json:"team_id"`
	RuleIDs     []string `json:"rule_ids"`
}

// RuleSet is a collection of rules
type RuleSet struct {
	ID          string        `json:"id"`
	TeamID      string        `json:"team_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	RuleIDs     []string      `json:"rule_ids"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Rules       []rules.Rule  `json:"rules"`
	Deprecated  bool          `json:"has_deprecated_rules"`
	IsUsed      bool          `json:"has_projects_assigned"`
	DeletedAt   *sql.NullTime `json:"deleted_at,omitempty"`
	DeletedBy   string        `json:"deleted_by,omitempty"`
}

// NameForID represents the data object for ruleset name and its ID
type NameForID struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	TeamID string `json:"team_id"`
}

// Status is a representation of an analysis status
type Status struct {
	AnalysisID string `json:"analysis_id"`
	ProjectID  string `json:"project_id"`
	Status     string `json:"status"`
}

// String returns a JSON formatted string of the ruleset object
func (r RuleSet) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf("failed to format ruleset: %v", err.Error())
	}
	return string(b)
}

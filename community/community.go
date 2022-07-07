package community

import (
	"time"
)

const (
	// GetRepoEndpoint is a string representation of the current endpoint for getting repo
	GetRepoEndpoint = `v1/repo/getRepo`
	// GetReposInCommonEndpoint is a string representation of the current endpoint for getting repos
	GetReposInCommonEndpoint = `/v1/repo/getReposInCommon`
	// GetReposForActorEndpoint is a string representation of the current endpoint for getting repos
	GetReposForActorEndpoint = `v1/repo/getReposForActor`
	// SearchRepoEndpoint is a string representation of the current endpoint for searching repo
	SearchRepoEndpoint = `v1/repo/search`
	// GetMetricsEndpoint is a string representation of the current endpoint for getting metrics
	GetMetricsEndpoint = `v1/repo/getMetrics`
)

// Repo is a representation of a github repo and corresponding metrics about
// that repo pulled from github
type Repo struct {
	ID            string     `json:"id" xml:"id"`
	Name          string     `json:"name" xml:"name"`
	URL           string     `json:"url" xml:"url"`
	Committers    int        `json:"committers" xml:"committers"`
	TotalActors   int        `json:"total_actors,omitempty" xml:"total_actors,omitempty"`
	Confidence    float64    `json:"confidence" xml:"confidence"`
	OldNames      []string   `json:"old_names" xml:"old_names"`
	DefaultBranch string     `json:"default_branch,omitempty" xml:"default_branch,omitempty"`
	MasterBranch  string     `json:"master_branch,omitempty" xml:"master_branch,omitempty"`
	Stars         int        `json:"stars" xml:"stars"`
	Matches       []string   `json:"matches,omitempty" xml:"matches,omitempty"`
	CommittedAt   time.Time  `json:"committed_at" xml:"committed_at"`
	UpdatedAt     time.Time  `json:"updated_at" xml:"updated_at"`
	CreatedAt     *time.Time `json:"created_at" xml:"created_at"`
}

// Metrics is a set of data points that represents the measure of a softwares
// community health
type Metrics struct {
	ID                                        string                                  `json:"id" xml:"id"`
	Name                                      string                                  `json:"name" xml:"name"`
	Committers                                int                                     `json:"committers" xml:"committers"`
	TotalActors                               int                                     `json:"total_actors,omitempty" xml:"total_actors,omitempty"`
	CommittersMonthlyCount                    *[]MonthlyCount                         `json:"committers_monthly_count" xml:"committers_monthly_count"`
	ReleasesTotalCount                        *int                                    `json:"releases_total_count" xml:"releases_total_count"`
	ReleasesMonthlyCount                      *[]MonthlyCount                         `json:"releases_monthly_count" xml:"releases_monthly_count"`
	ReleasesLastAt                            *time.Time                              `json:"releases_last_at" xml:"releases_last_at"`
	PullRequestsTotalCount                    *int                                    `json:"pull_requests_total_count" xml:"pull_requests_total_count"`
	PullRequestsLastAt                        *time.Time                              `json:"pull_requests_last_at" xml:"pull_requests_last_at"`
	PullRequestsMonthlyCount                  *[]MonthlyCount                         `json:"pull_requests_monthly_count" xml:"pull_requests_monthly_count"`
	IssuesLastAt                              *time.Time                              `json:"issues_last_at" xml:"issues_last_at"`
	IssuesOpenMonthlyCount                    *[]MonthlyCount                         `json:"issues_open_monthly_count" xml:"issues_open_monthly_count"`
	IssuesClosedMonthlyCount                  *[]MonthlyCount                         `json:"issues_closed_monthly_count" xml:"issues_closed_monthly_count"`
	IssuesClosedMttrMonthly                   *[]MonthlyMttr                          `json:"issues_closed_mttr_monthly" xml:"issues_closed_mttr_monthly"`
	IssuesClosedMttr                          *float64                                `json:"issues_closed_mttr" xml:"issues_closed_mttr"`
	CommitsTotalCount                         *int                                    `json:"commits_total_count" xml:"commits_total_count"`
	CommitsMonthlyCount                       *[]MonthlyCount                         `json:"commits_monthly_count" xml:"commits_monthly_count"`
	ActorsMonthlyCount                        *[]MonthlyCount                         `json:"actors_monthly_count" xml:"actors_monthly_count"`
	ActionsTotalCount                         *int                                    `json:"actions_total_count" xml:"actions_total_count"`
	ActionsLastAt                             *time.Time                              `json:"actions_last_at" xml:"actions_last_at"`
	ActionsFirstAt                            *time.Time                              `json:"actions_first_at" xml:"actions_first_at"`
	ActionsMonthlyCount                       *[]MonthlyCount                         `json:"actions_monthly_count" xml:"actions_monthly_count"`
	ContributingActorsTotalCount              *int                                    `json:"contributing_actors_total_count" xml:"contributing_actors_total_count"`
	ContributingActorsMonthlyCount            *[]MonthlyCount                         `json:"contributing_actors_monthly_count" xml:"contributing_actors_monthly_count"`
	ContributingActionsTotalCount             *int                                    `json:"contributing_actions_total_count" xml:"contributing_actions_total_count"`
	ContributingActionsLastAt                 *time.Time                              `json:"contributing_actions_last_at" xml:"contributing_actions_last_at"`
	ContributingActionsMonthlyCount           *[]MonthlyCount                         `json:"contributing_actions_monthly_count" xml:"contributing_actions_monthly_count"`
	NewActorsMonthlyCount                     *[]MonthlyCount                         `json:"new_actors_monthly_count" xml:"new_actors_monthly_count"`
	MedianWorkingHour                         *int                                    `json:"median_working_hour" xml:"median_working_hour"`
	AverageMonthlyActions                     *float64                                `json:"average_monthly_actions" xml:"average_monthly_actions"`
	AverageMonthlyActors                      *float64                                `json:"average_monthly_actors" xml:"average_monthly_actors"`
	AverageMonthlyCommits                     *float64                                `json:"average_monthly_commits" xml:"average_monthly_commits"`
	AverageMonthlyCommitters                  *float64                                `json:"average_monthly_committers" xml:"average_monthly_committers"`
	AverageMonthlyContributingActions         *float64                                `json:"average_monthly_contributing_actions" xml:"average_monthly_contributing_actions"`
	AverageMonthlyContributingActors          *float64                                `json:"average_monthly_contributing_actors" xml:"average_monthly_contributing_actors"`
	AverageMonthlyGrowthNewActors             *float64                                `json:"average_monthly_growth_new_actors" xml:"average_monthly_growth_new_actors"`
	AverageMonthlyGrowthNewContributingActors *float64                                `json:"average_monthly_growth_new_contributing_actors" xml:"average_monthly_growth_new_contributing_actors"`
	AverageMonthlyIssuesClosed                *float64                                `json:"average_monthly_issues_closed" xml:"average_monthly_issues_closed"`
	AverageMonthlyIssuesOpen                  *float64                                `json:"average_monthly_issues_open" xml:"average_monthly_issues_open"`
	AverageMonthlyNewActors                   *float64                                `json:"average_monthly_new_actors" xml:"average_monthly_new_actors"`
	AverageMonthlyPassiveActions              *float64                                `json:"average_monthly_passive_actions" xml:"average_monthly_passive_actions"`
	AverageMonthlyPassiveActors               *float64                                `json:"average_monthly_passive_actors" xml:"average_monthly_passive_actors"`
	AverageMonthlyPrComments                  *float64                                `json:"average_monthly_pr_comments" xml:"average_monthly_pr_comments"`
	AverageMonthlyPullRequests                *float64                                `json:"average_monthly_pull_requests" xml:"average_monthly_pull_requests"`
	AverageMonthlyReleases                    *float64                                `json:"average_monthly_releases" xml:"average_monthly_releases"`
	ContributingActorActivityRate             *float64                                `json:"contributing_actor_activity_rate" xml:"contributing_actor_activity_rate"`
	MedianWorkingHourDiversityIndex           *float64                                `json:"median_working_hour_diversity_index" xml:"median_working_hour_diversity_index"`
	PassiveActionsTotalCount                  *int                                    `json:"passive_actions_total_count" xml:"passive_actions_total_count"`
	PassiveActorsTotalCount                   *int                                    `json:"passive_actors_total_count" xml:"passive_actors_total_count"`
	PrCommentTotalCount                       *int                                    `json:"pr_comment_total_count" xml:"pr_comment_total_count"`
	ActorRoleDiversityIndex                   *float64                                `json:"actor_role_diversity_index" xml:"actor_role_diversity_index"`
	ContributingActorActivityRateMonthly      *[]MonthlyContributingActorActivityRate `json:"contributing_actor_activity_rate_monthly" xml:"contributing_actor_activity_rate_monthly"`
	PassiveActorsMonthlyCount                 *[]MonthlyCount                         `json:"passive_actors_monthly_count" xml:"passive_actors_monthly_count"`
	PassiveActionsLastAt                      *time.Time                              `json:"passive_actions_last_at" xml:"passive_actions_last_at"`
	PassiveActionsMonthlyCount                *[]MonthlyCount                         `json:"passive_actions_monthly_count" xml:"passive_actions_monthly_count"`
	NewContributingActorsMonthlyCount         *[]MonthlyCount                         `json:"new_contributing_actors_monthly_count" xml:"new_contributing_actors_monthly_count"`
	PrCommentMonthlyCount                     *[]MonthlyCount                         `json:"pr_comment_monthly_count" xml:"pr_comment_monthly_count"`
}

// MonthlyCount defines the data needed for month and count
type MonthlyCount struct {
	Month string  `json:"month" xml:"month"`
	Count float32 `json:"count" xml:"count"`
}

// MonthlyMttr defines the data needed for month and mttr
type MonthlyMttr struct {
	Month string  `json:"month" xml:"month"`
	Mttr  float32 `json:"mttr" xml:"mttr"`
}

// MonthlyMttr defines the data needed for month and mttr
type MonthlyContributingActorActivityRate struct {
	Month                         string  `json:"month" xml:"month"`
	ContributingActorActivityRate float32 `json:"contributing_actor_activity_rate" xml:"contributing_actor_activity_rate"`
}

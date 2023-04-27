package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/ion-channel/ionic/pagination"

	"github.com/ion-channel/ionic/teams"
)

// CreateTeamOptions represents all the values that can be provided for a team
// at the time of creation
type CreateTeamOptions struct {
	Name           string `json:"name"`
	OrganizationID string `json:"organization_id"`
	POCName        string `json:"poc_name"`
	POCEmail       string `json:"poc_email"`
}

// CreateTeam takes a create team options, validates the minimum info is
// present, and makes the calls to create the team. It returns the team created
// and any errors it encounters with the API.
func (ic *IonClient) CreateTeam(opts CreateTeamOptions, token string) (*teams.Team, error) {
	//no empty or whitespace-only names
	if len(strings.TrimSpace(opts.Name)) == 0 {
		return nil, fmt.Errorf("name cannot be empty or whitespace")
	}

	b, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)

	b, err = ic.Post(teams.TeamsCreateTeamEndpoint, token, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create team: %v", err.Error())
	}

	var t teams.Team
	err = json.Unmarshal(b, &t)
	if err != nil {
		return nil, fmt.Errorf("failed to parse team from response: %v", err.Error())
	}

	return &t, nil
}

// GetTeam takes a team id and returns the Ion Channel representation of that
// team.  An error is returned for client communications and unmarshalling
// errors.
func (ic *IonClient) GetTeam(id, token string) (*teams.Team, error) {
	params := url.Values{}
	params.Set("someid", id)

	b, _, err := ic.Get(teams.TeamsGetTeamEndpoint, token, params, nil, pagination.Pagination{})
	if err != nil {
		return nil, fmt.Errorf("failed to get team: %v", err.Error())
	}

	var team teams.Team
	err = json.Unmarshal(b, &team)
	if err != nil {
		return nil, fmt.Errorf("cannot parse team: %v", err.Error())
	}

	return &team, nil
}

// GetTeams returns the Ion Channel representation of that
// team.  An error is returned for client communications and unmarshalling
// errors.
func (ic *IonClient) GetTeams(token string) ([]teams.Team, error) {
	b, _, err := ic.Get(teams.TeamsGetTeamsEndpoint, token, nil, nil, pagination.Pagination{})
	if err != nil {
		return nil, fmt.Errorf("failed to get teams: %v", err.Error())
	}

	var ts []teams.Team
	err = json.Unmarshal(b, &ts)
	if err != nil {
		return nil, fmt.Errorf("cannot parse teams: %v", err.Error())
	}

	return ts, nil
}

// UpdateTeam takes a team ID and updates fields related to that team.
func (ic *IonClient) UpdateTeam(id, name, contactName, contactEmail, defaultDeployKey, token string) (*teams.Team, error) {
	params := url.Values{}
	params.Set("id", id)

	teamRequest := teams.Team{
		Name:             name,
		POCName:          contactName,
		POCEmail:         contactEmail,
		DefaultDeployKey: defaultDeployKey,
	}

	body, err := json.Marshal(teamRequest)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(body)

	response, err := ic.Put(teams.TeamsUpdateTeamEndpoint, token, params, *buffer, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update team: %w", err)
	}

	var team teams.Team
	err = json.Unmarshal(response, &team)
	if err != nil {
		return nil, fmt.Errorf("cannot parse team: %w", err)
	}

	return &team, nil
}

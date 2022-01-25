package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/teamusers"
)

// CreateTeamUserOptions represents all the values that can be provided for a team
// user at the time of creation
type CreateTeamUserOptions struct {
	Status    string `json:"status"`
	Role      string `json:"role"`
	TeamID    string `json:"team_id"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CreateTeamUser takes a create team options, validates the minimum info is
// present, and makes the calls to create the team. It returns the team created
// and any errors it encounters with the API.
func (ic *IonClient) CreateTeamUser(opts CreateTeamUserOptions, token string) (*teamusers.TeamUser, error) {
	b, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	b, err = ic.Post(teamusers.TeamsCreateTeamUserEndpoint, token, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create team user: %v", err.Error())
	}

	var tu teamusers.TeamUser
	err = json.Unmarshal(b, &tu)
	if err != nil {
		return nil, fmt.Errorf("failed to parse team user from response: %v", err.Error())
	}

	return &tu, nil
}

// UpdateTeamUser takes a teamUser object in the desired state and then makes the calls to update the teamUser.
// It returns the update teamUser and any errors it encounters with the API.
func (ic *IonClient) UpdateTeamUser(teamuser *teamusers.TeamUser, token string) (*teamusers.TeamUser, error) {
	params := url.Values{}
	params.Set("someid", teamuser.ID)

	b, err := json.Marshal(teamuser)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	b, err = ic.Put(teamusers.TeamsUpdateTeamUserEndpoint, token, params, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update team user: %v", err.Error())
	}

	var tu teamusers.TeamUser
	err = json.Unmarshal(b, &tu)
	if err != nil {
		return nil, fmt.Errorf("failed to parse team user from response: %v", err.Error())
	}

	return &tu, nil
}

// DeleteTeamUser takes a teamUser object and then makes the call to delete the teamUser.
// It returns any errors it encounters with the API.
func (ic *IonClient) DeleteTeamUser(teamuser *teamusers.TeamUser, token string) error {
	params := url.Values{}
	params.Set("someid", teamuser.ID)

	_, err := json.Marshal(teamuser)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	_, err = ic.Delete(teamusers.TeamsDeleteTeamUserEndpoint, token, params, nil)
	if err != nil {
		return fmt.Errorf("failed to delete team user: %v", err.Error())
	}

	return nil
}

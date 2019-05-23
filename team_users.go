package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/teamusers"
)

const (
	teamsCreateTeamUserEndpoint = "v1/teamUsers/createTeamUser"
	teamsGetTeamUserEndpoint    = "v1/teamUsers/getTeamUser"
	teamsUpdateTeamUserEndpoint = "v1/teamUsers/updateTeamUser"
	teamsDeleteTeamUserEndpoint = "v1/teamUsers/deleteTeamUser"
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
	b, err = ic.Post(teamsCreateTeamUserEndpoint, token, nil, *buff, nil)
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

// GetTeamUser takes a team id and returns the Ion Channel representation of that
// team.  An error is returned for client communications and unmarshalling
// errors.
func (ic *IonClient) GetTeamUser(teamID, userID, token string) (*teamusers.TeamUser, error) {
	params := &url.Values{}
	params.Set("team_id", teamID)
	params.Set("user_id", userID)

	b, err := ic.Get(teamsGetTeamUserEndpoint, token, params, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get team: %v", err.Error())
	}

	var teamU teamusers.TeamUser
	err = json.Unmarshal(b, &teamU)
	if err != nil {
		return nil, fmt.Errorf("cannot parse team: %v", err.Error())
	}

	return &teamU, nil
}

// UpdateTeamUser takes a teamUser object in the desired state and then makes the calls to update the teamUser.
// It returns the update teamUser and any errors it encounters with the API.
func (ic *IonClient) UpdateTeamUser(teamuser *teamusers.TeamUser, token string) (*teamusers.TeamUser, error) {
	params := &url.Values{}
	params.Set("someid", teamuser.ID)

	b, err := json.Marshal(teamuser)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	b, err = ic.Put(teamsUpdateTeamUserEndpoint, token, params, *buff, nil)
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
// Once the delete call has been made, a GetTeamUser call is made to validate the deletion.
// It returns a string or any errors it encounters with the API.
func (ic *IonClient) DeleteTeamUser(teamuser *teamusers.TeamUser, token string) (json.RawMessage, error) {
	params := &url.Values{}
	params.Set("someid", teamuser.ID)

	_, err := json.Marshal(teamuser)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	response, err := ic.Delete(teamsDeleteTeamUserEndpoint, token, params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete team user: %v", err.Error())
	}

	params = &url.Values{}
	params.Set("team_id", teamuser.TeamID)
	params.Set("user_id", teamuser.UserID)

	t, err := ic.Get(teamsGetTeamUserEndpoint, token, params, nil, nil)
	if err == nil {
		var teamU teamusers.TeamUser
		err = json.Unmarshal(t, &teamU)
		if err != nil {
			return nil, fmt.Errorf("cannot parse team: %v", err.Error())
		}
		return nil, fmt.Errorf("failed to validate team user deletion: %v", t)
	}
	return response, nil
}
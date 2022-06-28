package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ion-channel/ionic/pagination"
	"net/url"

	"github.com/ion-channel/ionic/errors"
	"github.com/ion-channel/ionic/requests"
)

const (
	// UsersCreateUserEndpoint is a string representation of the current endpoint for creating users
	UsersCreateUserEndpoint = "v1/users/createUser"
	// UsersGetSelfEndpoint is a string representation of the current endpoint for get user self
	UsersGetSelfEndpoint = "v1/users/getSelf"
	// UsersGetUserEndpoint is a string representation of the current endpoint for getting user
	UsersGetUserEndpoint = "v1/users/getUser"
	// UsersGetUsers is a string representation of the current endpoint for getting users
	UsersGetUsers = "v1/users/getUsers"
	// UsersGetUserNames is a string representation of the current endpoint for getting users
	UsersGetUserNames = "v1/users/getUserNames"
	// UsersUpdatePreferencesEndpoint is the current endpoint for updating a user's preferences
	UsersUpdatePreferencesEndpoint = "v1/users/userPreferences"
)

// String returns a JSON formatted string of the user object
func (u User) String() string {
	b, err := json.Marshal(u)
	if err != nil {
		return fmt.Sprintf("failed to format user: %v", err.Error())
	}
	return string(b)
}

// IsMemberOfOrganization takes a team id and returns true if user is a member of that team.
// Deprecated: Check permissions instead.
func (u User) IsMemberOfOrganization(id string) bool {
	for _, role := range u.Organizations {
		if role.Organization.ID == id {
			return true
		}
	}

	return false
}

// IsAdminOfOrganization takes a team id and returns true if user is an admin of that team.
// Deprecated: Check permissions instead.
func (u User) IsAdminOfOrganization(id string) bool {
	for _, role := range u.Organizations {
		// "admin" is for backwards compatibility with the old role system.
		// It can be removed when that role is no longer used.
		if role.Organization.ID == id && (role.Role == "admin" || role.Role == OrganizationRoleManager || role.Role == OrganizationRoleOwner) {
			return true
		}
	}

	return false
}

// IsMemberOfTeam takes a team id and returns true if user is a member of that team.
func (u User) IsMemberOfTeam(id string) bool {
	for _, role := range u.Teams {
		if role.TeamID == id {
			return true
		}
	}

	return false
}

// IsAdminOfTeam takes a team id and returns true if user is an admin of that team.
func (u User) IsAdminOfTeam(id string) bool {
	for _, role := range u.Teams {
		if role.TeamID == id && role.Role == "admin" {
			return true
		}
	}

	return false
}

type CreateUserOptions struct {
	Email                string `json:"email"`
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

// CreateUser takes an email, username, and password.  The username and password
// are not required, and can be left blank if so chosen.  It will return the
// instantiated user object from the API or an error if it encounters one with
// the API.
func (ic *IonClient) CreateUser(opts CreateUserOptions, token string) (User, error) {
	if opts.Email == "" {
		return User{}, fmt.Errorf("create user: email is required")
	}

	b, err := json.Marshal(opts)
	if err != nil {
		return User{}, errors.Prepend("create user: failed to marshal request", err)
	}

	buff := bytes.NewBuffer(b)

	b, err = ic.Post(UsersCreateUserEndpoint, token, nil, *buff, nil)
	if err != nil {
		return User{}, errors.Prepend("create user", err)
	}

	var u User
	err = json.Unmarshal(b, &u)
	if err != nil {
		return User{}, errors.Prepend("create user: failed to unmarshal user", err)
	}

	return u, nil
}

// GetSelf returns the user object associated with the bearer token provided.
// An error is returned if the client cannot talk to the API or the returned
// user object is nil or blank
func (ic *IonClient) GetSelf(token string) (User, error) {
	b, _, err := ic.Get(UsersGetSelfEndpoint, token, nil, nil, pagination.Pagination{})
	if err != nil {
		return User{}, errors.Prepend("get self", err)
	}

	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		return user, errors.Prepend("get self: failed unmarshaling user", err)
	}

	return user, nil
}

// GetUser returns the user object associated with the bearer token provided.
// An error is returned if the client cannot talk to the API or the returned
// user object is nil or blank
func (ic *IonClient) GetUser(id, token string) (User, error) {
	params := url.Values{}
	params.Set("id", id)

	b, _, err := ic.Get(UsersGetUserEndpoint, token, params, nil, pagination.Pagination{})
	if err != nil {
		return User{}, errors.Prepend("get user", err)
	}

	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		return User{}, errors.Prepend("get user: failed unmarshaling user", err)
	}

	return user, nil
}

// GetUsers requests and returns all users for a given installation
func (ic *IonClient) GetUsers(token string) ([]User, error) {
	b, _, err := ic.Get(UsersGetUsers, token, nil, nil, pagination.Pagination{})
	if err != nil {
		return nil, errors.Prepend("get users", err)
	}

	var us []User
	err = json.Unmarshal(b, &us)
	if err != nil {
		return nil, errors.Prepend("get users: failed unmarshaling users", err)
	}

	return us, nil
}

type NameAndID struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetUserNames takes slice of ids and teamID and returns user names with their ids
func (ic *IonClient) GetUserNames(ids []string, teamID, token string) ([]NameAndID, error) {
	params := url.Values{}
	params.Set("team_id", teamID)

	byIDs := requests.ByIDs{
		IDs: ids,
	}

	b, err := json.Marshal(byIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	r, err := ic.Post(UsersGetUserNames, token, params, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get user names: %v", err.Error())
	}

	var s []NameAndID
	err = json.Unmarshal(r, &s)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user names: %v", err.Error())
	}

	return s, nil
}

// UpdateOwnUserPreferences takes a Preferences object and returns any errors that occurred while updating
// your preferences.
func (ic *IonClient) UpdateOwnUserPreferences(preferences Preferences, token string) error {
	return ic.UpdateUserPreferences("", preferences, token)
}

// UpdateUserPreferences takes a user ID and a Preferences object and returns any errors that occurred while updating
// the user's preferences.
func (ic *IonClient) UpdateUserPreferences(userID string, preferences Preferences, token string) error {
	params := url.Values{}
	params.Set("user_id", userID)

	b, err := json.Marshal(preferences)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)
	_, err = ic.Post(UsersUpdatePreferencesEndpoint, token, params, *buff, nil)
	if err != nil {
		return fmt.Errorf("failed to update user preferences: %v", err.Error())
	}

	return nil
}

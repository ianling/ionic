package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ion-channel/ionic/pagination"

	"github.com/ion-channel/ionic/users"

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
func (ic *IonClient) CreateUser(opts CreateUserOptions, token string) (users.User, error) {
	if opts.Email == "" {
		return users.User{}, fmt.Errorf("create user: email is required")
	}

	b, err := json.Marshal(opts)
	if err != nil {
		return users.User{}, errors.Prepend("create user: failed to marshal request", err)
	}

	buff := bytes.NewBuffer(b)

	b, err = ic.Post(UsersCreateUserEndpoint, token, nil, *buff, nil)
	if err != nil {
		return users.User{}, errors.Prepend("create user", err)
	}

	var u users.User
	err = json.Unmarshal(b, &u)
	if err != nil {
		return users.User{}, errors.Prepend("create user: failed to unmarshal user", err)
	}

	return u, nil
}

// GetSelf returns the user object associated with the bearer token provided.
// An error is returned if the client cannot talk to the API or the returned
// user object is nil or blank
func (ic *IonClient) GetSelf(token string) (users.User, error) {
	b, _, err := ic.Get(UsersGetSelfEndpoint, token, nil, nil, pagination.Pagination{})
	if err != nil {
		return users.User{}, errors.Prepend("get self", err)
	}

	var user users.User
	err = json.Unmarshal(b, &user)
	if err != nil {
		return user, errors.Prepend("get self: failed unmarshaling user", err)
	}

	return user, nil
}

// GetUser returns the user object associated with the bearer token provided.
// An error is returned if the client cannot talk to the API or the returned
// user object is nil or blank
func (ic *IonClient) GetUser(id, token string) (users.User, error) {
	params := url.Values{}
	params.Set("id", id)

	b, _, err := ic.Get(UsersGetUserEndpoint, token, params, nil, pagination.Pagination{})
	if err != nil {
		return users.User{}, errors.Prepend("get user", err)
	}

	var user users.User
	err = json.Unmarshal(b, &user)
	if err != nil {
		return users.User{}, errors.Prepend("get user: failed unmarshaling user", err)
	}

	return user, nil
}

// GetUsers requests and returns all users for a given installation
func (ic *IonClient) GetUsers(token string) ([]users.User, error) {
	b, _, err := ic.Get(UsersGetUsers, token, nil, nil, pagination.Pagination{})
	if err != nil {
		return nil, errors.Prepend("get users", err)
	}

	var us []users.User
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
func (ic *IonClient) UpdateOwnUserPreferences(preferences users.Preferences, token string) error {
	return ic.UpdateUserPreferences("", preferences, token)
}

// UpdateUserPreferences takes a user ID and a Preferences object and returns any errors that occurred while updating
// the user's preferences.
func (ic *IonClient) UpdateUserPreferences(userID string, preferences users.Preferences, token string) error {
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

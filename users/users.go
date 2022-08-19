package users

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/ion-channel/ionic/organizations"
)

type User struct {
	ID                string                               `json:"id"`
	Email             string                               `json:"email"`
	Username          string                               `json:"username"`
	CreatedAt         time.Time                            `json:"created_at"`
	UpdatedAt         time.Time                            `json:"updated_at"`
	LastActiveAt      time.Time                            `json:"last_active_at"`
	Status            UserStatus                           `json:"status"`
	ExternallyManaged bool                                 `json:"externally_managed"`
	Metadata          *string                              `json:"metadata"`
	SysAdmin          bool                                 `json:"sys_admin"`
	System            bool                                 `json:"system"`
	Organizations     []organizations.UserOrganizationRole `json:"organizations"`
	Teams             []UserTeamRole                       `json:"teams"`
}

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
		if role.Organization.ID == id && (role.Role == "admin" || role.Role == organizations.OrganizationRoleManager || role.Role == organizations.OrganizationRoleOwner) {
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

type UserTeamRole struct {
	Role   string `json:"role"`
	TeamID string `json:"team_id"`
}

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusDisabled UserStatus = "disabled"
)

var AllUserStatus = []UserStatus{
	UserStatusActive,
	UserStatusDisabled,
}

func (e UserStatus) IsValid() bool {
	switch e {
	case UserStatusActive, UserStatusDisabled:
		return true
	}
	return false
}

func (e UserStatus) String() string {
	return string(e)
}

func (e *UserStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserStatus", str)
	}
	return nil
}

func (e UserStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

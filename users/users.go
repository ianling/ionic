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

type Preferences struct {
	Flip                bool                        `json:"flip"`
	NotificationChannel NotificationChannelOption   `json:"notification_channel"`
	Frequency           NotificationFrequencyOption `json:"frequency"`
}

type NotificationChannelOption string

const (
	NotificationChannelOptionEmail NotificationChannelOption = "email"
)

var AllNotificationChannelOption = []NotificationChannelOption{
	NotificationChannelOptionEmail,
}

func (e NotificationChannelOption) IsValid() bool {
	switch e {
	case NotificationChannelOptionEmail:
		return true
	}
	return false
}

func (e NotificationChannelOption) String() string {
	return string(e)
}

func (e *NotificationChannelOption) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NotificationChannelOption(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NotificationChannelOption", str)
	}
	return nil
}

func (e NotificationChannelOption) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NotificationFrequencyOption string

const (
	NotificationFrequencyOptionDaily NotificationFrequencyOption = "daily"
)

var AllNotificationFrequencyOption = []NotificationFrequencyOption{
	NotificationFrequencyOptionDaily,
}

func (e NotificationFrequencyOption) IsValid() bool {
	switch e {
	case NotificationFrequencyOptionDaily:
		return true
	}
	return false
}

func (e NotificationFrequencyOption) String() string {
	return string(e)
}

func (e *NotificationFrequencyOption) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NotificationFrequencyOption(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NotificationFrequencyOption", str)
	}
	return nil
}

func (e NotificationFrequencyOption) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Permission string

const (
	// Grants all permissions over all objects in the system.
	// Reserved for System Admin role. Cannot be granted to or inherited by other roles.
	PermissionAllPermissions Permission = "ALL_PERMISSIONS"
	// Grants the ability to modify an organization's settings, including billing information.
	PermissionOrganizationModify Permission = "ORGANIZATION_MODIFY"
	// Grants the ability to view basic information about an organization, including its name,
	// the date and time it was created, etc.
	PermissionOrganizationView Permission = "ORGANIZATION_VIEW"
	// Grants the ability to invite users to an organization.
	PermissionOrganizationUserCreate Permission = "ORGANIZATION_USER_CREATE"
	// Grants the ability to change the roles of existing users within an organization.
	PermissionOrganizationUserModify Permission = "ORGANIZATION_USER_MODIFY"
	// Grants the ability to remove users from an organization.
	PermissionOrganizationUserRemove Permission = "ORGANIZATION_USER_REMOVE"
	// Grants the ability to view a list of all the organization's members.
	PermissionOrganizationUserView Permission = "ORGANIZATION_USER_VIEW"
	// Grants the ability to create a new software list and add it to an organization's software inventory.
	PermissionOrganizationSoftwareListCreate Permission = "ORGANIZATION_SOFTWARE_LIST_CREATE"
	// Grants the ability to modify an organization's existing software lists,
	// including adding, modifying, and removing components from individual software lists.
	PermissionOrganizationSoftwareListModify Permission = "ORGANIZATION_SOFTWARE_LIST_MODIFY"
	// Grants the ability to view an organization's software lists and any components they contain.
	// This also includes the ability to export an SBOM for the software lists and view risk scoring data.
	PermissionOrganizationSoftwareListView Permission = "ORGANIZATION_SOFTWARE_LIST_VIEW"
	// Grants the ability to remove software lists from an organization's software inventory.
	PermissionOrganizationSoftwareListRemove Permission = "ORGANIZATION_SOFTWARE_LIST_REMOVE"
)

var AllPermission = []Permission{
	PermissionAllPermissions,
	PermissionOrganizationModify,
	PermissionOrganizationView,
	PermissionOrganizationUserCreate,
	PermissionOrganizationUserModify,
	PermissionOrganizationUserRemove,
	PermissionOrganizationUserView,
	PermissionOrganizationSoftwareListCreate,
	PermissionOrganizationSoftwareListModify,
	PermissionOrganizationSoftwareListView,
	PermissionOrganizationSoftwareListRemove,
}

func (e Permission) IsValid() bool {
	switch e {
	case PermissionAllPermissions, PermissionOrganizationModify, PermissionOrganizationView, PermissionOrganizationUserCreate, PermissionOrganizationUserModify, PermissionOrganizationUserRemove, PermissionOrganizationUserView, PermissionOrganizationSoftwareListCreate, PermissionOrganizationSoftwareListModify, PermissionOrganizationSoftwareListView, PermissionOrganizationSoftwareListRemove:
		return true
	}
	return false
}

func (e Permission) String() string {
	return string(e)
}

func (e *Permission) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Permission(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Permission", str)
	}
	return nil
}

func (e Permission) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
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

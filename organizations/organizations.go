package organizations

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CreateOrganizationRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Organization struct {
	ID        string               `json:"id"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	DeletedAt *time.Time           `json:"deleted_at"`
	Name      string               `json:"name"`
	Members   []OrganizationMember `json:"members"`
}

type OrganizationMember struct {
	ID        string           `json:"id"`
	UserID    string           `json:"user_id"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	RoleID    string           `json:"role_id"`
	Role      OrganizationRole `json:"role"`
	CreatedAt time.Time        `json:"created_at"`
	JoinedAt  *time.Time       `json:"joined_at"`
	DeletedAt *time.Time       `json:"deleted_at"`
}

type OrganizationMemberUpdate struct {
	UserID    string     `json:"user_id"`
	RoleID    *string    `json:"role_id"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UpdateOrganizationMembersInput struct {
	OrgID   string                     `json:"org_id"`
	Members []OrganizationMemberUpdate `json:"members"`
}

type UserOrganizationRole struct {
	RoleID       string           `json:"role_id"`
	Role         OrganizationRole `json:"role"`
	Description  string           `json:"description"`
	Organization Organization     `json:"organization"`
}

type OrganizationRole string

const (
	OrganizationRoleOwner   OrganizationRole = "Owner"
	OrganizationRoleManager OrganizationRole = "Manager"
	OrganizationRoleMember  OrganizationRole = "Member"
)

var AllOrganizationRole = []OrganizationRole{
	OrganizationRoleOwner,
	OrganizationRoleManager,
	OrganizationRoleMember,
}

func (e OrganizationRole) IsValid() bool {
	switch e {
	case OrganizationRoleOwner, OrganizationRoleManager, OrganizationRoleMember:
		return true
	}
	return false
}

func (e OrganizationRole) String() string {
	return string(e)
}

func (e *OrganizationRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrganizationRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrganizationRole", str)
	}
	return nil
}

func (e OrganizationRole) MarshalGQL(w io.Writer) {
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

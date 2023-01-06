package organizations

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CreateOrganizationRequest struct {
	Name string `json:"name"`
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

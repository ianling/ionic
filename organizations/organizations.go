package organizations

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	// OrganizationsCreateEndpoint is the endpoint for creating an organization
	OrganizationsCreateEndpoint = "v1/organizations/createOrganization"
	// OrganizationsGetEndpoint is the endpoint for getting an organization
	OrganizationsGetEndpoint = "v1/organizations/getOrganization"
	// OrganizationsGetBulkEndpoint is the endpoint for getting organizations
	OrganizationsGetBulkEndpoint = "v1/organizations/getOrganizations"
	// OrganizationsUpdateEndpoint is the endpoint for updating an organization
	OrganizationsUpdateEndpoint = "v1/organizations/updateOrganization"
	// OrganizationsDisableEndpoint is the endpoint for disabling an organization
	OrganizationsDisableEndpoint = "v1/organizations/disableOrganization"
	// OrganizationsAddMemberEndpoint is the endpoint for adding an existing user as a member of an organization
	OrganizationsAddMemberEndpoint = "v1/organizations/addMember"
)

// Organization is a logical collection of teams.
type Organization struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name"`
}

// OrganizationRole is the type for constants that enumerate the different roles a user can have in an organization.
type OrganizationRole string

const (
	// OrganizationRoleAdmin is the administrator role in an organization
	OrganizationRoleAdmin = "admin"
	// OrganizationRoleMember is the regular member role in an organization
	OrganizationRoleMember = "member"
)

// String returns a JSON formatted string of the team object
func (o Organization) String() string {
	b, err := json.Marshal(o)
	if err != nil {
		return fmt.Sprintf("failed to format organization: %v", err.Error())
	}

	return string(b)
}

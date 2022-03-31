package organizations

import (
	"time"
)

const (
	// OrganizationsCreateEndpoint is the endpoint for creating an organization
	OrganizationsCreateEndpoint = "v1/organizations/createOrganization"
	// OrganizationsGetOwnEndpoint is the endpoint for getting the organizations the user belongs to
	OrganizationsGetOwnEndpoint = "v1/organizations/getOwnOrganizations"
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

type (
	// Organization is a logical collection of teams.
	// Users can be members of an Organization, which grants them access to all the Organization's teams.
	Organization struct {
		ID        string               `json:"id"`
		CreatedAt time.Time            `json:"created_at"`
		UpdatedAt time.Time            `json:"updated_at"`
		DeletedAt *time.Time           `json:"deleted_at"`
		Name      string               `json:"name"`
		Members   []OrganizationMember `json:"members"`
	}

	// OrganizationMember represents a particular user's role in an organization.
	OrganizationMember struct {
		UserID string           `json:"user_id"`
		Role   OrganizationRole `json:"role"`
	}

	// UserOrganizationRole represents a particular user's role in an organization in a standalone form,
	// containing information about both the user and the organization they belong to.
	UserOrganizationRole struct {
		OrganizationMember
		Organization Organization `json:"organization"`
	}
)

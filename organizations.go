package ionic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ion-channel/ionic/organizations"
	"github.com/ion-channel/ionic/pagination"
	"github.com/ion-channel/ionic/requests"
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
	// OrganizationsUpdateMembersEndpoint is the endpoint for altering existing members of an organization
	OrganizationsUpdateMembersEndpoint = "v1/organizations/updateMembers"
)

// CreateOrganizationOptions represents all the values that can be provided for an organization
// at the time of creation
type CreateOrganizationOptions struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

// CreateOrganization takes a create team options, validates the minimum info is
// present, and makes the calls to create the team. It returns the ID of the created organization
// and any errors it encounters with the API.
func (ic *IonClient) CreateOrganization(opts CreateOrganizationOptions, token string) (*organizations.Organization, error) {
	//no empty or whitespace-only names
	if len(strings.TrimSpace(opts.Name)) == 0 {
		return nil, fmt.Errorf("name cannot be empty or whitespace")
	}

	b, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)

	b, err = ic.Post(OrganizationsCreateEndpoint, token, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create organization: %v", err.Error())
	}

	var org organizations.Organization
	err = json.Unmarshal(b, &org)
	if err != nil {
		return nil, fmt.Errorf("failed to parse organization from response: %v", err.Error())
	}

	return &org, nil
}

// GetOwnOrganizations takes a token and returns a list of organizations the user belongs to.
func (ic *IonClient) GetOwnOrganizations(token string) (*[]organizations.UserOrganizationRole, error) {
	resp, _, err := ic.Get(OrganizationsGetOwnEndpoint, token, nil, nil, pagination.Pagination{})
	if err != nil {
		return nil, fmt.Errorf("failed to get own organizations: %v", err.Error())
	}

	var orgs []organizations.UserOrganizationRole
	err = json.Unmarshal(resp, &orgs)
	if err != nil {
		return nil, fmt.Errorf("cannot parse own organizations: %v", err.Error())
	}

	return &orgs, nil
}

// GetOrganization takes an organization id and returns the Ion Channel representation of that organization.
func (ic *IonClient) GetOrganization(id, token string) (*organizations.Organization, error) {
	b, _, err := ic.Get(fmt.Sprintf("%s/%s", OrganizationsGetEndpoint, id), token, nil, nil, pagination.Pagination{})
	if err != nil {
		return nil, fmt.Errorf("failed to get organization: %v", err.Error())
	}

	var organization organizations.Organization
	err = json.Unmarshal(b, &organization)
	if err != nil {
		return nil, fmt.Errorf("cannot parse organization: %v", err.Error())
	}

	return &organization, nil
}

// GetOrganizations takes one or more IDs and returns those organizations.
func (ic *IonClient) GetOrganizations(ids requests.ByIDs, token string) (*[]organizations.Organization, error) {
	b, err := json.Marshal(ids)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)

	resp, err := ic.Post(OrganizationsGetBulkEndpoint, token, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get organizations: %v", err.Error())
	}

	var orgs []organizations.Organization
	err = json.Unmarshal(resp, &orgs)
	if err != nil {
		return nil, fmt.Errorf("cannot parse organizations: %v", err.Error())
	}

	return &orgs, nil
}

// UpdateOrganization takes an organization ID, and the fields to update, returns the updated organization.
func (ic *IonClient) UpdateOrganization(id string, name string, token string) (*organizations.Organization, error) {
	req := struct {
		Name string `json:"name"`
	}{Name: name}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)

	resp, err := ic.Put(fmt.Sprintf("%s/%s", OrganizationsUpdateEndpoint, id), token, nil, *buff, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update organization: %v", err.Error())
	}

	var org organizations.Organization
	err = json.Unmarshal(resp, &org)
	if err != nil {
		return nil, fmt.Errorf("cannot parse organization: %v", err.Error())
	}

	return &org, nil
}

// DisableOrganization takes an organization ID and returns any errors that occurred.
func (ic *IonClient) DisableOrganization(id string, token string) error {
	_, err := ic.Delete(fmt.Sprintf("%s/%s", OrganizationsDisableEndpoint, id), token, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to disable organization: %v", err.Error())
	}

	return nil
}

// AddMemberToOrganization takes an organization ID, a user ID, and a role, and returns any errors that occurred.
func (ic *IonClient) AddMemberToOrganization(organizationID string, userID string, roleID string, token string) error {
	req := struct {
		UserID string `json:"user_id"`
		RoleID string `json:"role"`
	}{
		UserID: userID,
		RoleID: roleID,
	}

	b, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)

	_, err = ic.Post(fmt.Sprintf("%s/%s", OrganizationsAddMemberEndpoint, organizationID), token, nil, *buff, nil)
	if err != nil {
		return fmt.Errorf("failed to add member to organization: %v", err.Error())
	}

	return nil
}

// UpdateOrganizationMembers takes an organization ID and a slice of UpdateOrganizationMemberInput, and returns any errors that occurred.
func (ic *IonClient) UpdateOrganizationMembers(organizationID string, usersToUpdate []organizations.OrganizationMemberUpdate, token string) error {
	b, err := json.Marshal(usersToUpdate)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err.Error())
	}

	buff := bytes.NewBuffer(b)

	_, err = ic.Put(fmt.Sprintf("%s/%s", OrganizationsUpdateMembersEndpoint, organizationID), token, nil, *buff, nil)
	if err != nil {
		return fmt.Errorf("failed to update organization members: %v", err.Error())
	}

	return nil
}

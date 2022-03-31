package organizations

// OrganizationRole is the type for constants that enumerate the different roles a user can have in an organization.
type OrganizationRole string

const (
	// OrganizationRoleOwner is the Owner role in an organization
	OrganizationRoleOwner OrganizationRole = "owner"
	// OrganizationRoleManager is the Manager role in an organization
	OrganizationRoleManager OrganizationRole = "manager"
	// OrganizationRoleMember is the Member role in an organization
	OrganizationRoleMember OrganizationRole = "member"
	// OrganizationRoleSystemAdmin denotes that the user is a system admin and therefore has full access to the organization
	OrganizationRoleSystemAdmin OrganizationRole = "system admin"
)

package teams

// TeamRole is the type for constants that enumerate the different roles a user can have in a Team.
type TeamRole string

const (
	// TeamRoleAdmin is the Admin role in a Team
	TeamRoleAdmin TeamRole = "admin"
	// TeamRoleMember is the Member role in a Team
	TeamRoleMember TeamRole = "member"
	// TeamRoleSystemAdmin denotes that the user is a system admin and therefore has full access to the Team
	TeamRoleSystemAdmin TeamRole = "system admin"
)

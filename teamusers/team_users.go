package teamusers

import (
	"time"
)

const (
	// TeamsCreateTeamUserEndpoint is a string representation of the current endpoint for creating team user
	TeamsCreateTeamUserEndpoint = "v1/teamUsers/createTeamUser"
	// TeamsUpdateTeamUserEndpoint is a string representation of the current endpoint for updating team user
	TeamsUpdateTeamUserEndpoint = "v1/teamUsers/updateTeamUser"
	// TeamsDeleteTeamUserEndpoint is a string representation of the current endpoint for deleting team user
	TeamsDeleteTeamUserEndpoint = "v1/teamUsers/deleteTeamUser"
	// TeamsGetTeamUserEndpoint is a string representation of the current endpoint for getting team users
	TeamsGetTeamUserEndpoint = "v1/teamUsers/getTeamUsers"
)

// TeamUser is a representation of an Ion Channel Team User relationship within the system
type TeamUser struct {
	ID        string    `json:"id"`
	TeamID    string    `json:"team_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Status    string    `json:"status"`
	Role      string    `json:"role"`
}

// TeamUserRole contains information about a user's role on a team.
type TeamUserRole struct {
	ID           string    `json:"id"`
	TeamID       string    `json:"team_id"`
	UserID       string    `json:"user_id"`
	Role         string    `json:"role"`
	Status       string    `json:"status"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	LastActiveAt time.Time `json:"last_active_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	InvitedAt    time.Time `json:"invited_at"`
}

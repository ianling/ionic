package aliases

import "time"

//Alias map user defined project names to a Common Platform Enumeration (CPE).
type Alias struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Org       string    `json:"org"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   string    `json:"version"`
}

const (
	// AddAliasEndpoint allows the user to attach an alias to a project.  Consists of org, name, version. Requires team id and project id.
	AddAliasEndpoint = "v1/project/addAlias"
)

// Equal checks if a given Alias is equivalent to another Alias, based on the name, org, and version information
// they contain. Returns true if they are equivalent.
func (a Alias) Equal(x Alias) bool {
	if a.Name == x.Name && a.Version == x.Version && a.Org == x.Org {
		return true
	}

	return false
}

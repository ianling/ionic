package projects

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ion-channel/ionic/aliases"
	"github.com/ion-channel/ionic/rulesets"
	"github.com/ion-channel/ionic/tags"
)

const (
	validEmailRegex     = `(?i)^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	validGitURIRegex    = `^(?:(?:http|ftp|gopher|mailto|mid|cid|news|nntp|prospero|telnet|rlogin|tn3270|wais|svn|git|rsync)+\+ssh\:\/\/|git\+https?:\/\/|git\@|(?:http|ftp|gopher|mailto|mid|cid|news|nntp|prospero|telnet|rlogin|tn3270|wais|svn|git|rsync|ssh|file|s3)+s?:\/\/)[^\s]+$`
	validDockerURIRegex = `[a-z0-9]+(?:[._-][a-z0-9]+)*`
)

const (
	// CreateProjectEndpoint is a string representation of the current endpoint for creating project
	CreateProjectEndpoint = "v1/project/createProject"
	// CreateProjectsFromCSVEndpoint is a string representation of the current endpoint for creating projects from CSV
	CreateProjectsFromCSVEndpoint = "v1/project/createProjectsCSV"
	// GetProjectEndpoint is a string representation of the current endpoint for getting project
	GetProjectEndpoint = "v1/project/getProject"
	// GetProjectByURLEndpoint is a string representation of the current endpoint for getting project by URL
	GetProjectByURLEndpoint = "v1/project/getProjectByUrl"
	// GetProjectsEndpoint is a string representation of the current endpoint for getting projects
	GetProjectsEndpoint = "v1/project/getProjects"
	// UpdateProjectEndpoint is a string representation of the current endpoint for updating project
	UpdateProjectEndpoint = "v1/project/updateProject"
	// GetUsedRulesetIdsEndpoint is a string representation of the current endpoint for getting list of a team's in use rulesets
	GetUsedRulesetIdsEndpoint = "v1/project/getUsedRulesetIds"
	// GetProjectsNamesEndpoint is a string representation of the current endpoint for getting projects alaised names and versions
	GetProjectsNamesEndpoint = "v1/project/getProjectsNames"
)

var (
	// ErrInvalidProject is returned when a given project does not pass the
	// standards for a project
	ErrInvalidProject = fmt.Errorf("project has invalid fields")
)

// Project is a representation of a project within the Ion Channel system
type Project struct {
	ID               *string                          `json:"id,omitempty"`
	TeamID           *string                          `json:"team_id,omitempty"`
	RulesetID        *string                          `json:"ruleset_id,omitempty"`
	Name             *string                          `json:"name,omitempty"`
	Type             *string                          `json:"type,omitempty"`
	Source           *string                          `json:"source,omitempty"`
	Branch           *string                          `json:"branch,omitempty"`
	Description      *string                          `json:"description,omitempty"`
	Active           bool                             `json:"active"`
	Draft            bool                             `json:"draft"`
	ChatChannel      string                           `json:"chat_channel"`
	CreatedAt        time.Time                        `json:"created_at"`
	UpdatedAt        time.Time                        `json:"updated_at"`
	DeployKey        string                           `json:"deploy_key"`
	Monitor          bool                             `json:"should_monitor"`
	MonitorFrequency string                           `json:"monitor_frequency"`
	POCName          string                           `json:"poc_name"`
	POCEmail         string                           `json:"poc_email"`
	Username         string                           `json:"username"`
	Password         string                           `json:"password"`
	KeyFingerprint   string                           `json:"key_fingerprint"`
	Private          bool                             `json:"private"`
	Aliases          []aliases.Alias                  `json:"aliases"`
	Tags             []tags.Tag                       `json:"tags"`
	RulesetHistory   []rulesets.ProjectRulesetHistory `json:"ruleset_history"`
	SoftwareListID   string                           `json:"sbom_id"`
	ComponentID      string                           `json:"sbom_entry_id"`
	CPE              string                           `json:"cpe"`
	PURL             string                           `json:"purl"`
}

// RulesetID represents a ruleset ID
type RulesetID struct {
	RulesetID string `json:"ruleset_id"`
}

// Name represents a single project name and id
type Name struct {
	ID          string `json:"project_id"`
	Name        string `json:"name"`
	ProductName string `json:"product_name"`
	Version     string `json:"version"`
	Org         string `json:"org"`
}

// String returns a JSON formatted string of the project object
func (p Project) String() string {
	b, err := json.Marshal(p)
	if err != nil {
		return fmt.Sprintf("failed to format project: %v", err.Error())
	}
	return string(b)
}

// ProjectReachable checks if the artifact URL is reachable
func (p *Project) ProjectReachable(client http.Client) (map[string]string, error) {
	invalidFields := make(map[string]string)
	var projErr error
	if p.Type != nil {
		switch strings.ToLower(*p.Type) {
		case "artifact":
			u, err := url.Parse(*p.Source)
			if err != nil {
				invalidFields["source"] = fmt.Sprintf("source must be a valid url: %v", err.Error())
				projErr = ErrInvalidProject
			}

			if u != nil {
				res, err := client.Head(u.String())
				if err != nil {
					invalidFields["source"] = "source failed to return a response"
					projErr = ErrInvalidProject
				}

				if res != nil && res.StatusCode == http.StatusNotFound {
					invalidFields["source"] = "source returned a not found"
					projErr = ErrInvalidProject
				}
			}
		case "git", "svn", "s3":
			r := regexp.MustCompile(validGitURIRegex)
			if p.Source != nil && !r.MatchString(strings.ToLower(*p.Source)) {
				invalidFields["source"] = "source must be a valid uri"
				projErr = ErrInvalidProject
			}
		case "docker":
			r := regexp.MustCompile(validDockerURIRegex)
			if p.Source != nil && !r.MatchString(*p.Source) {
				invalidFields["source"] = "source must be a docker image name"
				projErr = ErrInvalidProject
			}
		case "source_unavailable":
			if p.Source != nil && len(*p.Source) > 0 {
				invalidFields["source"] = "source cannot be specified for this project type"
				projErr = ErrInvalidProject
			}
		default:
			invalidFields["type"] = "invalid type value"
			projErr = ErrInvalidProject
		}
	}

	return invalidFields, projErr
}

// ValidateRequiredFields verifies the project contains the fields required
func (p *Project) ValidateRequiredFields() (map[string]string, error) {
	invalidFields := make(map[string]string)
	var projErr error

	if p.TeamID == nil {
		invalidFields["team_id"] = "missing team id"
		projErr = ErrInvalidProject
	}

	if p.RulesetID == nil {
		invalidFields["ruleset_id"] = "missing ruleset id"
		projErr = ErrInvalidProject
	}

	if p.Name == nil {
		invalidFields["name"] = "missing name"
		projErr = ErrInvalidProject
	}

	if p.Type == nil {
		invalidFields["type"] = "missing type"
		projErr = ErrInvalidProject
	}

	if p.Source == nil && (p.Type != nil && strings.ToLower(*p.Type) != "source_unavailable") {
		invalidFields["source"] = "missing source"
		projErr = ErrInvalidProject
	}

	if p.Branch == nil && p.Type != nil && strings.ToLower(*p.Type) == "git" {
		invalidFields["branch"] = "missing branch"
		projErr = ErrInvalidProject
	}

	if p.Description == nil {
		invalidFields["description"] = "missing description"
		projErr = ErrInvalidProject
	}

	p.POCEmail = strings.TrimSpace(p.POCEmail)

	r := regexp.MustCompile(validEmailRegex)
	if p.POCEmail != "" && !r.MatchString(p.POCEmail) {
		invalidFields["poc_email"] = "invalid email supplied"
		projErr = ErrInvalidProject
	}

	if p.Type != nil {
		switch strings.ToLower(*p.Type) {
		case "artifact":
			_, err := url.Parse(*p.Source)
			if err != nil {
				invalidFields["source"] = fmt.Sprintf("source must be a valid url: %v", err.Error())
				projErr = ErrInvalidProject
			}
		case "git", "svn", "s3":
			r := regexp.MustCompile(validGitURIRegex)
			if p.Source != nil && !r.MatchString(strings.ToLower(*p.Source)) {
				invalidFields["source"] = "source must be a valid uri"
				projErr = ErrInvalidProject
			}
		case "docker":
			r := regexp.MustCompile(validDockerURIRegex)
			if p.Source != nil && !r.MatchString(*p.Source) {
				invalidFields["source"] = "source must be a docker image name"
				projErr = ErrInvalidProject
			}
		case "source_unavailable":
			if p.Source != nil && len(*p.Source) > 0 {
				invalidFields["source"] = "source cannot be specified for this project type"
				projErr = ErrInvalidProject
			}
		default:
			invalidFields["type"] = "invalid type value"
			projErr = ErrInvalidProject
		}
	}

	return invalidFields, projErr
}

// Validate takes an http client; returns a slice of fields as a string and
// an error. The fields will be a list of fields that did not pass the
// validation. An error will only be returned if any of the fields fail their
// validation.
// Since this also checks for project reachability, ValidateRequiredFields
// can be used to skip that check.
func (p *Project) Validate(client http.Client) (map[string]string, error) {
	invalidFields := make(map[string]string)
	var projErr error

	if p.TeamID == nil {
		invalidFields["team_id"] = "missing team id"
		projErr = ErrInvalidProject
	}

	if p.RulesetID == nil {
		invalidFields["ruleset_id"] = "missing ruleset id"
		projErr = ErrInvalidProject
	}

	if p.Name == nil {
		invalidFields["name"] = "missing name"
		projErr = ErrInvalidProject
	}

	if p.Type == nil {
		invalidFields["type"] = "missing type"
		projErr = ErrInvalidProject
	}

	if p.Source == nil && (p.Type != nil && strings.ToLower(*p.Type) != "source_unavailable") {
		invalidFields["source"] = "missing source"
		projErr = ErrInvalidProject
	}

	if p.Branch == nil && p.Type != nil && strings.ToLower(*p.Type) == "git" {
		invalidFields["branch"] = "missing branch"
		projErr = ErrInvalidProject
	}

	if p.Description == nil {
		invalidFields["description"] = "missing description"
		projErr = ErrInvalidProject
	}

	p.POCEmail = strings.TrimSpace(p.POCEmail)

	r := regexp.MustCompile(validEmailRegex)
	if p.POCEmail != "" && !r.MatchString(p.POCEmail) {
		invalidFields["poc_email"] = "invalid email supplied"
		projErr = ErrInvalidProject
	}

	if p.Type != nil {
		switch strings.ToLower(*p.Type) {
		case "artifact":
			u, err := url.Parse(*p.Source)
			if err != nil {
				invalidFields["source"] = fmt.Sprintf("source must be a valid url: %v", err.Error())
				projErr = ErrInvalidProject
			}

			if u != nil {
				res, err := client.Head(u.String())
				if err != nil {
					invalidFields["source"] = "artifact source failed to return a response"
					projErr = ErrInvalidProject
				}

				if res != nil && res.StatusCode == http.StatusNotFound {
					invalidFields["source"] = fmt.Sprintf("artifact source of %v returned a 404 (not found)", *p.Source)
					projErr = ErrInvalidProject
				}
			}
		case "git", "svn", "s3":
			r := regexp.MustCompile(validGitURIRegex)
			if p.Source != nil && !r.MatchString(strings.ToLower(*p.Source)) {
				invalidFields["source"] = "source must be a valid uri"
				projErr = ErrInvalidProject
			}
		case "docker":
			r := regexp.MustCompile(validDockerURIRegex)
			if p.Source != nil && !r.MatchString(*p.Source) {
				invalidFields["source"] = "source must be a docker image name"
				projErr = ErrInvalidProject
			}
		case "source_unavailable":
			if p.Source != nil && len(*p.Source) > 0 {
				invalidFields["source"] = "source cannot be specified for this project type"
				projErr = ErrInvalidProject
			}
		default:
			invalidFields["type"] = "invalid type value"
			projErr = ErrInvalidProject
		}
	}

	return invalidFields, projErr
}

// Filter represents the available fields to filter a get project request
// with.
type Filter struct {
	// ID filters on a single ID
	ID *string `sql:"id"`
	// IDs filters on one or more IDs
	IDs            *[]string `sql:"id"`
	TeamID         *string   `sql:"team_id"`
	SoftwareListID *string   `sql:"sbom_id"`
	ComponentIDs   *[]string `sql:"sbom_entry_id"`
	Source         *string   `sql:"source"`
	Type           *string   `sql:"type"`
	Active         *bool     `sql:"active"`
	Draft          *bool     `sql:"draft"`
	Monitor        *bool     `sql:"should_monitor"`
}

// ParseParam takes a param string, breaks it apart, and repopulates it into a
// struct for further use. Any invalid or incomplete interpretations of a field
// will be ignored and only valid entries put into the struct.
func ParseParam(param string) Filter {
	pf := Filter{}

	fvs := strings.Split(param, ",")
	for ii := range fvs {
		parts := strings.SplitN(fvs[ii], ":", 2)
		if len(parts) != 2 {
			continue
		}

		name := parts[0]
		value := parts[1]

		field := reflect.ValueOf(&pf).Elem().FieldByNameFunc(func(n string) bool { return strings.ToLower(n) == strings.ToLower(name) })
		if !field.IsValid() {
			continue
		}

		kind := field.Type().Kind()
		if kind == reflect.Ptr {
			kind = field.Type().Elem().Kind()
		}

		switch kind {
		case reflect.String:
			field.Set(reflect.ValueOf(&value))
		case reflect.Bool:
			value, err := strconv.ParseBool(value)
			if err != nil {
				continue
			}

			field.Set(reflect.ValueOf(&value))
		case reflect.Slice:
			value := strings.Split(value, " ")
			field.Set(reflect.ValueOf(&value))
		default:
			// shouldn't ever happen, but just in case
			continue
		}
	}

	return pf
}

// Param converts the non nil fields of the Project Filter into a string usable
// for URL query params.
func (pf *Filter) Param() string {
	ps := make([]string, 0)

	// get the fields and values of the underlying Filter (Elem dereferences the pointer)
	fields := reflect.TypeOf(pf).Elem()
	values := reflect.ValueOf(pf).Elem()

	for i := 0; i < fields.NumField(); i++ {
		value := values.Field(i)

		if value.IsNil() {
			continue
		}

		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		name := fields.Field(i).Name

		switch value.Kind() {
		case reflect.String:
			ps = append(ps, fmt.Sprintf("%v:%v", name, value.String()))
		case reflect.Bool:
			ps = append(ps, fmt.Sprintf("%v:%v", name, value.Bool()))
		case reflect.Slice:
			// elements of the slice are separated by spaces:
			// example: IDs:abc def ghi
			sliceLen := value.Len()
			if sliceLen == 0 {
				continue
			}

			valueStr := strings.Join(value.Interface().([]string), " ")
			valueStr = fmt.Sprintf("%v:%v", name, valueStr)

			ps = append(ps, valueStr)
		}
	}

	return strings.Join(ps, ",")
}

// ProjectSliceContains checks if a given []Project contains a Project matching the given Project.
// This is used to prevent duplicate Projects from appearing in a slice of Projects.
// Returns true if the given Project has an alias that matches an alias of any other project in the slice.
func ProjectSliceContains(projects []Project, projectToFind Project) bool {
	for _, project := range projects {
		for _, alias := range project.Aliases {
			for _, aliasToFind := range projectToFind.Aliases {
				if alias.Equal(aliasToFind) {
					return true
				}
			}
		}
	}

	return false
}

// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package ionic

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Compliance struct {
	Passing int `json:"passing"`
	Failing int `json:"failing"`
}

type Component struct {
	ID            string                `json:"id"`
	SbomID        string                `json:"sbom_id"`
	Name          string                `json:"name"`
	Version       string                `json:"version"`
	Org           string                `json:"org"`
	Status        ComponentStatus       `json:"status"`
	SearchResults SearchResults         `json:"search_results"`
	Suggestions   []ComponentSuggestion `json:"suggestions"`
	CreatedAt     time.Time             `json:"created_at"`
	UpdatedAt     time.Time             `json:"updated_at"`
	DeletedAt     *time.Time            `json:"deleted_at"`
	ErrorMessage  *string               `json:"error_message"`
}

type ComponentSuggestion struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CreateOrganizationRequest struct {
	Name string `json:"name"`
}

type CreateSoftwareListRequest struct {
	Name             string  `json:"name"`
	OrgID            string  `json:"org_id"`
	Version          *string `json:"version"`
	SupplierName     *string `json:"supplier_name"`
	ContactName      *string `json:"contact_name"`
	ContactEmail     *string `json:"contact_email"`
	RulesetID        *string `json:"ruleset_id"`
	MonitorFrequency *string `json:"monitor_frequency"`
}

type Metrics struct {
	Risk       Risk       `json:"risk"`
	Compliance Compliance `json:"compliance"`
	Resolution Resolution `json:"resolution"`
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
	UserID    string           `json:"user_id"`
	Username  string           `json:"username"`
	Role      OrganizationRole `json:"role"`
	CreatedAt time.Time        `json:"created_at"`
	DeletedAt *time.Time       `json:"deleted_at"`
}

type OrganizationMemberUpdate struct {
	UserID    string            `json:"user_id"`
	Role      *OrganizationRole `json:"role"`
	DeletedAt *time.Time        `json:"deleted_at"`
}

type PackageSearchResult struct {
	SearchResult
	Purl string `json:"purl"`
}

type Preferences struct {
	Flip                bool                        `json:"flip"`
	NotificationChannel NotificationChannelOption   `json:"notification_channel"`
	Frequency           NotificationFrequencyOption `json:"frequency"`
}

type ProductSearchResult struct {
	SearchResult
	Cpe string `json:"cpe"`
}

type RepoSearchResult struct {
	SearchResult
	RepoURL string `json:"repo_url"`
}

type Resolution struct {
	Resolved          int `json:"resolved"`
	PartiallyResolved int `json:"partially_resolved"`
	Unresolved        int `json:"unresolved"`
}

type Risk struct {
	Score  *int        `json:"score"`
	Scopes []RiskScope `json:"scopes"`
}

type RiskScope struct {
	Name  string `json:"name"`
	Value *int   `json:"value"`
}

type SearchResult struct {
	ID                    string  `json:"id"`
	Confidence            float64 `json:"confidence"`
	IsUserInput           bool    `json:"is_user_input"`
	Selected              bool    `json:"selected"`
	AutomaticallySelected bool    `json:"automatically_selected"`
	Name                  string  `json:"name"`
	Org                   string  `json:"org"`
	Version               string  `json:"version"`
}

type SearchResults struct {
	Package []PackageSearchResult `json:"package"`
	Repo    []RepoSearchResult    `json:"repo"`
	Product []ProductSearchResult `json:"product"`
}

type SoftwareInventory struct {
	ID            string         `json:"id"`
	Organization  Metrics        `json:"organization"`
	SoftwareLists []SoftwareList `json:"softwareLists"`
}

type SoftwareList struct {
	ID               string             `json:"id"`
	Name             string             `json:"name"`
	Version          string             `json:"version"`
	Supplier         string             `json:"supplier"`
	ContactName      string             `json:"contact_name"`
	ContactEmail     string             `json:"contact_email"`
	MonitorFrequency string             `json:"monitor_frequency"`
	Status           SoftwareListStatus `json:"status"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
	DeletedAt        *time.Time         `json:"deleted_at"`
	EntryCount       *int               `json:"entry_count"`
	Metrics          Metrics            `json:"metrics"`
	Entries          []Component        `json:"entries"`
	TeamID           string             `json:"team_id"`
	OrgID            string             `json:"org_id"`
	RulesetID        string             `json:"ruleset_id"`
}

type UpdateOrganizationMembersInput struct {
	OrgID   string                     `json:"org_id"`
	Members []OrganizationMemberUpdate `json:"members"`
}

type User struct {
	ID                string                 `json:"id"`
	Email             string                 `json:"email"`
	Username          string                 `json:"username"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
	LastActiveAt      time.Time              `json:"last_active_at"`
	Status            UserStatus             `json:"status"`
	ExternallyManaged bool                   `json:"externally_managed"`
	Metadata          *string                `json:"metadata"`
	SysAdmin          bool                   `json:"sys_admin"`
	System            bool                   `json:"system"`
	Organizations     []UserOrganizationRole `json:"organizations"`
	Teams             []UserTeamRole         `json:"teams"`
}

type UserOrganizationRole struct {
	Role         OrganizationRole `json:"role"`
	Organization Organization     `json:"organization"`
}

type UserTeamRole struct {
	Role   string `json:"role"`
	TeamID string `json:"team_id"`
}

type ComponentStatus string

const (
	ComponentStatusNoResolution      ComponentStatus = "no_resolution"
	ComponentStatusPartialResolution ComponentStatus = "partial_resolution"
	ComponentStatusResolved          ComponentStatus = "resolved"
	ComponentStatusErrored           ComponentStatus = "errored"
	ComponentStatusDeleted           ComponentStatus = "deleted"
)

var AllComponentStatus = []ComponentStatus{
	ComponentStatusNoResolution,
	ComponentStatusPartialResolution,
	ComponentStatusResolved,
	ComponentStatusErrored,
	ComponentStatusDeleted,
}

func (e ComponentStatus) IsValid() bool {
	switch e {
	case ComponentStatusNoResolution, ComponentStatusPartialResolution, ComponentStatusResolved, ComponentStatusErrored, ComponentStatusDeleted:
		return true
	}
	return false
}

func (e ComponentStatus) String() string {
	return string(e)
}

func (e *ComponentStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ComponentStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ComponentStatus", str)
	}
	return nil
}

func (e ComponentStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NotificationChannelOption string

const (
	NotificationChannelOptionEmail NotificationChannelOption = "email"
)

var AllNotificationChannelOption = []NotificationChannelOption{
	NotificationChannelOptionEmail,
}

func (e NotificationChannelOption) IsValid() bool {
	switch e {
	case NotificationChannelOptionEmail:
		return true
	}
	return false
}

func (e NotificationChannelOption) String() string {
	return string(e)
}

func (e *NotificationChannelOption) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NotificationChannelOption(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NotificationChannelOption", str)
	}
	return nil
}

func (e NotificationChannelOption) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NotificationFrequencyOption string

const (
	NotificationFrequencyOptionDaily NotificationFrequencyOption = "daily"
)

var AllNotificationFrequencyOption = []NotificationFrequencyOption{
	NotificationFrequencyOptionDaily,
}

func (e NotificationFrequencyOption) IsValid() bool {
	switch e {
	case NotificationFrequencyOptionDaily:
		return true
	}
	return false
}

func (e NotificationFrequencyOption) String() string {
	return string(e)
}

func (e *NotificationFrequencyOption) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NotificationFrequencyOption(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NotificationFrequencyOption", str)
	}
	return nil
}

func (e NotificationFrequencyOption) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
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

type SearchResultType string

const (
	SearchResultTypePackage SearchResultType = "package"
	SearchResultTypeProduct SearchResultType = "product"
	SearchResultTypeRepo    SearchResultType = "repo"
)

var AllSearchResultType = []SearchResultType{
	SearchResultTypePackage,
	SearchResultTypeProduct,
	SearchResultTypeRepo,
}

func (e SearchResultType) IsValid() bool {
	switch e {
	case SearchResultTypePackage, SearchResultTypeProduct, SearchResultTypeRepo:
		return true
	}
	return false
}

func (e SearchResultType) String() string {
	return string(e)
}

func (e *SearchResultType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SearchResultType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SearchResultType", str)
	}
	return nil
}

func (e SearchResultType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SoftwareListStatus string

const (
	SoftwareListStatusCreated          SoftwareListStatus = "created"
	SoftwareListStatusAutocompletedone SoftwareListStatus = "autocompletedone"
	SoftwareListStatusAllconfirmed     SoftwareListStatus = "allconfirmed"
)

var AllSoftwareListStatus = []SoftwareListStatus{
	SoftwareListStatusCreated,
	SoftwareListStatusAutocompletedone,
	SoftwareListStatusAllconfirmed,
}

func (e SoftwareListStatus) IsValid() bool {
	switch e {
	case SoftwareListStatusCreated, SoftwareListStatusAutocompletedone, SoftwareListStatusAllconfirmed:
		return true
	}
	return false
}

func (e SoftwareListStatus) String() string {
	return string(e)
}

func (e *SoftwareListStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SoftwareListStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SoftwareListStatus", str)
	}
	return nil
}

func (e SoftwareListStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusDisabled UserStatus = "disabled"
)

var AllUserStatus = []UserStatus{
	UserStatusActive,
	UserStatusDisabled,
}

func (e UserStatus) IsValid() bool {
	switch e {
	case UserStatusActive, UserStatusDisabled:
		return true
	}
	return false
}

func (e UserStatus) String() string {
	return string(e)
}

func (e *UserStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserStatus", str)
	}
	return nil
}

func (e UserStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

package software_lists

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

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

package searches

import (
	"time"

	"github.com/ion-channel/ionic/risk"
)

// Report represents all data in report from a search
// across multiple sources
type Report struct {
	Name           string              `json:"name" xml:"name"`
	Org            string              `json:"org" xml:"org"`
	Version        string              `json:"version" xml:"version"`
	Type           string              `json:"type" xml:"type"`
	Origin         string              `json:"origin" xml:"origin"`
	CreatedAt      time.Time           `json:"created_at" xml:"created_at"`
	UpdatedAt      time.Time           `json:"updated_at" xml:"updated_at"`
	Confidence     float64             `json:"confidence" xml:"confidence"`
	URL            string              `json:"url" xml:"url"`
	ExternalID     string              `json:"external_id" xml:"external_id"`
	ExternalURL    string              `json:"external_url" xml:"external_url"`
	Scores         risk.Scores         `json:"scores" xml:"scores"`
	Matches        []string            `json:"matches" xml:"matches"`
	VerifiedEntity risk.VerifiedEntity `json:"verified_entity" xml:"verified_entity"`
}

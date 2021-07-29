package searches

import (
	"time"
)

// Report represents all data in report from a search
// across multiple sources
type Report struct {
	Name        string    `json:"name" xml:"name"`
	Org         string    `json:"org" xml:"org"`
	Version     string    `json:"version" xml:"version"`
	Type        string    `json:"type" xml:"type"`
	CreatedAt   time.Time `json:"created_at" xml:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" xml:"updated_at"`
	Confidence  float64   `json:"confidence" xml:"confidence"`
	URL         string    `json:"url" xml:"url"`
	ExternalID  string    `json:"external_id" xml:"external_id"`
	ExternalURL string    `json:"external_url" xml:"external_url"`
}

package deliveries

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	// DeliveriesGetDestinationEndpoints returns all destinations for a team. Requires team id.
	DeliveriesGetDestinationsEndpoints = "/v1/teams/getDeliveryDestinations"
	// DeliveriesDeleteDestinationEndpoint markes a delivery destination as deleted. It requires a delivery destination id.
	DeliveriesDeleteDestinationEndpoint = "/v1/teams/deleteDeliveryDestination"
)

// Destination is a representation of a single location that a team can deliver results to.
type Destination struct {
	ID        string `json:"id"`
	TeamID    string `json:"team_id"`
	Location  string `json:"location"`
	Region    string `json:"region"`
	accessKey string
	secretKey string
	Name      string     `json:"name"`
	DestType  string     `json:"type"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// String returns a JSON formatted string of the delivery object
func (p Destination) String() string {
	b, err := json.Marshal(p)
	if err != nil {
		return fmt.Sprintf("failed to format delivery: %v", err.Error())
	}
	return string(b)
}

package risk

import "time"

// MetricPoint defines the data needed for points on a single risk point
type MetricPoint struct {
	Name   string `json:"name" xml:"name"`
	Points int    `json:"points" xml:"points"`
}

// MetricPoints defines the data needed for points on a single risk point
type MetricPoints struct {
	Metrics     []MetricPoint `json:"metrics" xml:"metrics"`
	ProcessedAt time.Time     `json:"processed_at" xml:"processed_at"`
}

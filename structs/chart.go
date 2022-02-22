package structs

import "time"

type ChartResult struct {
	Data        []Chart `json:"Data"`
	DataVersion int     `json:"DataVersion"`
}

type Chart struct {
	Close    float64   `json:"Close"`
	High     float64   `json:"High"`
	Interest int       `json:"Interest"`
	Low      float64   `json:"Low"`
	Open     float64   `json:"Open"`
	Time     time.Time `json:"Time"`
	Volume   int       `json:"Volume"`
}

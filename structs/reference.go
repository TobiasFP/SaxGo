package structs

import "time"

type ExchangeResult struct {
	AllDay           bool   `json:"AllDay"`
	CountryCode      string `json:"CountryCode"`
	Currency         string `json:"Currency"`
	ExchangeID       string `json:"ExchangeId"`
	ExchangeSessions []struct {
		EndTime   time.Time `json:"EndTime"`
		StartTime time.Time `json:"StartTime"`
		State     string    `json:"State"`
	} `json:"ExchangeSessions"`
	Mic                  string `json:"Mic"`
	Name                 string `json:"Name"`
	PriceSourceName      string `json:"PriceSourceName"`
	TimeZone             int    `json:"TimeZone"`
	TimeZoneAbbreviation string `json:"TimeZoneAbbreviation"`
	TimeZoneOffset       string `json:"TimeZoneOffset"`
}

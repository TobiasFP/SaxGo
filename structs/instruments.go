package structs

type InstrumentsResult struct {
	Data []Instrument `json:"Data"`
}
type Instrument struct {
	AssetType    string   `json:"AssetType"`
	CurrencyCode string   `json:"CurrencyCode"`
	Description  string   `json:"Description"`
	ExchangeId   string   `json:"ExchangeId"`
	GroupId      int      `json:"GroupId"`
	Identifier   int      `json:"Identifier"`
	SummaryType  string   `json:"SummaryType"`
	Symbol       string   `json:"Symbol"`
	TradableAs   []string `json:"TradableAs"`
}

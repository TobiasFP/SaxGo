package structs

import "time"

type OrderResult struct {
	OrderId string `json:"OrderId"`
}

type TradeOrder struct {
	Uic           int     `json:"Uic"`
	BuySell       string  `json:"BuySell"`
	AssetType     string  `json:"AssetType"`
	Amount        float64 `json:"Amount"`
	AmountType    string  `json:"AmountType"`
	OrderPrice    float64 `json:"OrderPrice"`
	OrderType     string  `json:"OrderType"`
	OrderRelation string  `json:"OrderRelation"`
	ManualOrder   bool    `json:"ManualOrder"`
	OrderDuration struct {
		DurationType string `json:"DurationType"`
	} `json:"OrderDuration"`
	AccountKey string `json:"AccountKey"`
}

type InfoPriceResult struct {
	AssetType              string                 `json:"AssetType"`
	InstrumentPriceDetails InstrumentPriceDetails `json:"InstrumentPriceDetails"`
	DisplayAndFormat       DisplayAndFormat       `json:"DisplayAndFormat"`
	PriceInfo              PriceInfo              `json:"PriceInfo"`
	PriceInfoDetails       PriceInfoDetails       `json:"PriceInfoDetails"`
	LastUpdated            time.Time              `json:"LastUpdated"`
	PriceSource            string                 `json:"PriceSource"`
	Quote                  Quote                  `json:"Quote"`
	Uic                    int                    `json:"Uic"`
}

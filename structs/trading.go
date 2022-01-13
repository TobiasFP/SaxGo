package structs

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

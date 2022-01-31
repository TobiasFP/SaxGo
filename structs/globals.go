package structs

type DisplayAndFormat struct {
	Currency    string  `json:"Currency"`
	Decimals    float64 `json:"Decimals"`
	Description string  `json:"Description"`
	Format      string  `json:"Format"`
	Symbol      string  `json:"Symbol"`
}

type PositionBase struct {
	AccountId                  string  `json:"AccountId"`
	Amount                     float64 `json:"Amount"`
	AssetType                  string  `json:"AssetType"`
	CanBeClosed                bool    `json:"CanBeClosed"`
	ClientId                   string  `json:"ClientId"`
	CloseConversionRateSettled bool    `json:"CloseConversionRateSettled"`
	ExecutionTimeOpen          string  `json:"ExecutionTimeOpen"`
	IsForceOpen                bool    `json:"IsForceOpen"`
	IsMarketOpen               bool    `json:"IsMarketOpen"`
	OpenPrice                  float64 `json:"OpenPrice"`
	SpotDate                   string  `json:"SpotDate"`
	Status                     string  `json:"Status"`
	Uic                        float64 `json:"Uic"`
	ValueDate                  string  `json:"ValueDate"`
	SourceOrderId              string  `json:"SourceOrderId"`
}

type PositionView struct {
	Ask                             float64 `json:"Ask"`
	Bid                             float64 `json:"Bid"`
	CalculationReliability          string  `json:"CalculationReliability"`
	CurrentPrice                    float64 `json:"CurrentPrice"`
	CurrentPriceDelayMinutes        float64 `json:"CurrentPriceDelayMinutes"`
	CurrentPriceType                string  `json:"CurrentPriceType"`
	Exposure                        float64 `json:"Exposure"`
	ExposureCurrency                string  `json:"ExposureCurrency"`
	ExposureInBaseCurrency          float64 `json:"ExposureInBaseCurrency"`
	InstrumentPriceDayPercentChange float64 `json:"InstrumentPriceDayPercentChange"`
	ProfitLossOnTrade               float64 `json:"ProfitLossOnTrade"`
	ProfitLossOnTradeInBaseCurrency float64 `json:"ProfitLossOnTradeInBaseCurrency"`
	SettlementInstruction           struct {
		ActualRolloverAmount            float64 `json:""`
		ActualSettlementAmount          float64 `json:""`
		Amount                          float64 `json:""`
		IsSettlementInstructionsAllowed bool    `json:""`
		Month                           float64 `json:""`
		SettlementType                  string  `json:""`
		Year                            float64 `json:""`
	} `json:"SettlementInstruction"`
	TradeCostsTotal               float64 `json:"TradeCostsTotal"`
	TradeCostsTotalInBaseCurrency float64 `json:"TradeCostsTotalInBaseCurrency"`
}

type Quote struct {
	Amount           int     `json:"Amount"`
	Ask              float64 `json:"Ask"`
	Bid              float64 `json:"Bid"`
	DelayedByMinutes int     `json:"DelayedByMinutes"`
	ErrorCode        string  `json:"ErrorCode"`
	MarketState      string  `json:"MarketState"`
	Mid              float64 `json:"Mid"`
	PriceSource      string  `json:"PriceSource"`
	PriceSourceType  string  `json:"PriceSourceType"`
	PriceTypeAsk     string  `json:"PriceTypeAsk"`
	PriceTypeBid     string  `json:"PriceTypeBid"`
}

type InstrumentPriceDetails struct {
	IsMarketOpen       bool   `json:"IsMarketOpen"`
	ShortTradeDisabled bool   `json:"ShortTradeDisabled"`
	ValueDate          string `json:"ValueDate"`
}

type PriceInfo struct {
	High          float64 `json:"High"`
	Low           float64 `json:"Low"`
	NetChange     float64 `json:"NetChange"`
	PercentChange float64 `json:"PercentChange"`
}
type PriceInfoDetails struct {
	AskSize        float64 `json:"AskSize"`
	BidSize        float64 `json:"BidSize"`
	LastClose      float64 `json:"LastClose"`
	LastTraded     float64 `json:"LastTraded"`
	LastTradedSize float64 `json:"LastTradedSize"`
	Open           float64 `json:"Open"`
	Volume         float64 `json:"Volume"`
}

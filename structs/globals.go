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
	ClientId                   float64 `json:"ClientId"`
	CloseConversionRateSettled bool    `json:"CloseConversionRateSettled"`
	ExecutionTimeOpen          string  `json:"ExecutionTimeOpen"`
	IsForceOpen                bool    `json:"IsForceOpen"`
	IsMarketOpen               bool    `json:"IsMarketOpen"`
	OpenPrice                  float64 `json:"OpenPrice"`
	SpotDate                   string  `json:"SpotDate"`
	Status                     string  `json:"Status"`
	Uic                        float64 `json:"Uic"`
	ValueDate                  string  `json:"ValueDate"`
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

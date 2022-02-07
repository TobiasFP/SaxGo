package structs

type Positions struct {
	Count float64    `json:"__count"`
	Data  []Position `json:"Data"`
}

type Position struct {
	NetPositionId    string           `json:"NetPositionId"`
	PositionId       int              `json:"PositionId"`
	PositionView     PositionView     `json:"PositionView"`
	PositionBase     PositionBase     `json:"PositionBase"`
	DisplayAndFormat DisplayAndFormat `json:"DisplayAndFormat"`
	Exchange         exchange         `json:"Exchange"`
}

type MyBalance struct {
	CalculationReliability string  `json:"CalculationReliability"`
	CashBalance            float64 `json:"CashBalance"`
	CashBlocked            float64 `json:"CashBlocked"`
	ChangesScheduled       bool    `json:"ChangesScheduled"`
	ClosedPositionsCount   float64 `json:"ClosedPositionsCount"`
	CollateralAvailable    float64 `json:"CollateralAvailable"`
	CollateralCreditValue  struct {
		Line           float64
		UtilizationPct float64
	} `json:"CollateralCreditValue"`
	CorporateActionUnrealizedAmounts float64 `json:"CorporateActionUnrealizedAmounts"`
	CostToClosePositions             float64 `json:"CostToClosePositions"`
	Currency                         string  `json:"Currency"`
	CurrencyDecimals                 float64 `json:"CurrencyDecimals"`
	InitialMargin                    struct {
		CollateralAvailable   float64 `json:"CollateralAvailable"`
		CollateralCreditValue struct {
			Line           float64 `json:"Line"`
			UtilizationPct float64 `json:"UtilizationPct"`
		} `json:"CollateralCreditValue"`
		MarginAvailable              float64 `json:"MarginAvailable"`
		MarginCollateralNotAvailable float64 `json:"MarginCollateralNotAvailable"`
		MarginUsedByCurrentPositions float64 `json:"MarginUsedByCurrentPositions"`
		MarginUtilizationPct         float64 `json:"MarginUtilizationPct"`
		NetEquityForMargin           float64 `json:"NetEquityForMargin"`
		OtherCollateralDeduction     float64 `json:"OtherCollateralDeduction"`
	} `json:"InitialMargin"`
	IsPortfolioMarginModelSimple     bool    `json:"IsPortfolioMarginModelSimple"`
	MarginAvailableForTrading        float64 `json:"MarginAvailableForTrading"`
	MarginCollateralNotAvailable     float64 `json:"MarginCollateralNotAvailable"`
	MarginExposureCoveragePct        float64 `json:"MarginExposureCoveragePct"`
	MarginNetExposure                float64 `json:"MarginNetExposure"`
	MarginUsedByCurrentPositions     float64 `json:"MarginUsedByCurrentPositions"`
	MarginUtilizationPct             float64 `json:"MarginUtilizationPct"`
	NetEquityForMargin               float64 `json:"NetEquityForMargin"`
	NetPositionsCount                float64 `json:"NetPositionsCount"`
	NonMarginPositionsValue          float64 `json:"NonMarginPositionsValue"`
	OpenIpoOrdersCount               float64 `json:"OpenIpoOrdersCount"`
	OpenPositionsCount               float64 `json:"OpenPositionsCount"`
	OptionPremiumsMarketValue        float64 `json:"OptionPremiumsMarketValue"`
	OrdersCount                      float64 `json:"OrdersCount"`
	OtherCollateral                  float64 `json:"OtherCollateral"`
	SettlementValue                  float64 `json:"SettlementValue"`
	TotalValue                       float64 `json:"TotalValue"`
	TransactionsNotBooked            float64 `json:"TransactionsNotBooked"`
	TriggerOrdersCount               float64 `json:"TriggerOrdersCount"`
	UnrealizedMarginClosedProfitLoss float64 `json:"UnrealizedMarginClosedProfitLoss"`
	UnrealizedMarginOpenProfitLoss   float64 `json:"UnrealizedMarginOpenProfitLoss"`
	UnrealizedMarginProfitLoss       float64 `json:"UnrealizedMarginProfitLoss"`
	UnrealizedPositionsValue         float64 `json:"UnrealizedPositionsValue"`
}

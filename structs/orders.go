package structs

type Orders struct {
	Count float64 `json:"__count"`
	Data  []Order `json:"Data"`
}

type displayandformat struct {
	Currency    string  `json:"Currency"`
	Decimals    float64 `json:"Decimals"`
	Description string  `json:"Description"`
	Format      string  `json:"Format"`
	Symbol      string  `json:"Symbol"`
}

type duration struct {
	DurationType string `json:"DurationType"`
}
type exchange struct {
	Description string `json:"Description"`
	ExchangeId  string `json:"ExchangeId"`
	IsOpen      bool   `json:"IsOpen"`
	TimeZoneId  string `json:"TimeZoneId"`
}

type Order struct {
	AccountId                string           `json:"AccountId"`
	AccountKey               string           `json:"AccountKey"`
	Amount                   float64          `json:"Amount"`
	Ask                      float64          `json:"Ask"`
	AssetType                string           `json:"AssetType"`
	Bid                      float64          `json:"Bid"`
	BuySell                  string           `json:"BuySell"`
	CalculationReliability   string           `json:"CalculationReliability"`
	ClientId                 string           `json:"ClientId"`
	ClientKey                string           `json:"ClientKey"`
	ClientName               string           `json:"ClientName"`
	ClientNote               string           `json:"ClientNote"`
	CorrelationKey           string           `json:"CorrelationKey"`
	CurrentPrice             float64          `json:"CurrentPrice"`
	CurrentPriceDelayMinutes float64          `json:"CurrentPriceDelayMinutes"`
	CurrentPriceType         string           `json:"CurrentPriceType"`
	DisplayAndFormat         displayandformat `json:"DisplayAndFormat"`
	DistanceToMarket         float64          `json:"DistanceToMarket"`
	Duration                 duration         `json:"Duration"`
	Exchange                 exchange         `json:"Exchange"`
	IpoSubscriptionFee       float64          `json:"IpoSubscriptionFee"`
	IsForceOpen              bool             `json:"IsForceOpen"`
	IsMarketOpen             bool             `json:"IsMarketOpen"`
	MarketPrice              float64          `json:"MarketPrice"`
	MarketState              string           `json:"MarketState"`
	MarketValue              float64          `json:"MarketValue"`
	NonTradableReason        string           `json:"NonTradableReason"`
	OpenOrderType            string           `json:"OpenOrderType"`
	OrderAmountType          string           `json:"OrderAmountType"`
	OrderId                  string           `json:"OrderId"`
	OrderRelation            string           `json:"OrderRelation"`
	OrderTime                string           `json:"OrderTime"`
	Price                    float64          `json:"Price"`
	RelatedOpenOrders        []string         `json:"RelatedOpenOrders"`
	Status                   string           `json:"Status"`
	TradingStatus            string           `json:"TradingStatus"`
	Uic                      int              `json:"Uic"`
}

package structs

import "time"

type User struct {
	ClientKey                         string   `json:"ClientKey"`
	Culture                           string   `json:"Culture"`
	Language                          string   `json:"Language"`
	LastLoginStatus                   string   `json:"LastLoginStatus"`
	LastLoginTime                     string   `json:"LastLoginTime"`
	LegalAssetTypes                   []string `json:"LegalAssetTypes"`
	MarketDataViaOpenApiTermsAccepted bool     `json:"MarketDataViaOpenApiTermsAccepted"`
	Name                              string   `json:"Name"`
	TimeZoneId                        float64  `json:"TimeZoneId"`
	UserId                            string   `json:"UserId"`
	UserKey                           string   `json:"UserKey"`
}

type AccountResult struct {
	Data []Account `json:"Data"`
}

type Account struct {
	AccountGroupKey                       string    `json:"AccountGroupKey"`
	AccountID                             string    `json:"AccountId"`
	AccountKey                            string    `json:"AccountKey"`
	AccountSubType                        string    `json:"AccountSubType"`
	AccountType                           string    `json:"AccountType"`
	Active                                bool      `json:"Active"`
	CanUseCashPositionsAsMarginCollateral bool      `json:"CanUseCashPositionsAsMarginCollateral"`
	CfdBorrowingCostsActive               bool      `json:"CfdBorrowingCostsActive"`
	ClientID                              string    `json:"ClientId"`
	ClientKey                             string    `json:"ClientKey"`
	CreationDate                          time.Time `json:"CreationDate"`
	Currency                              string    `json:"Currency"`
	CurrencyDecimals                      int       `json:"CurrencyDecimals"`
	DirectMarketAccess                    bool      `json:"DirectMarketAccess"`
	IndividualMargining                   bool      `json:"IndividualMargining"`
	IsCurrencyConversionAtSettlementTime  bool      `json:"IsCurrencyConversionAtSettlementTime"`
	IsMarginTradingAllowed                bool      `json:"IsMarginTradingAllowed"`
	IsShareable                           bool      `json:"IsShareable"`
	IsTrialAccount                        bool      `json:"IsTrialAccount"`
	LegalAssetTypes                       []string  `json:"LegalAssetTypes"`
	ManagementType                        string    `json:"ManagementType"`
	MarginCalculationMethod               string    `json:"MarginCalculationMethod"`
	MarginLendingEnabled                  string    `json:"MarginLendingEnabled"`
	PortfolioBasedMarginEnabled           bool      `json:"PortfolioBasedMarginEnabled"`
	Sharing                               []string  `json:"Sharing"`
	SupportsAccountValueProtectionLimit   bool      `json:"SupportsAccountValueProtectionLimit"`
	UseCashPositionsAsMarginCollateral    bool      `json:"UseCashPositionsAsMarginCollateral"`
}

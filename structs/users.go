package structs

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

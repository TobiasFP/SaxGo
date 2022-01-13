package saxgo

import (
	"context"
	"net/http"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

var (
	clientID           = ""
	clientSecret       = ""
	AccessToken        = ""
	RefreshToken       = ""
	AccountKey         = ""
	Time         int64 = 0
)

func getClient() (*http.Client, error) {
	conf := &oauth2.Config{
		RedirectURL:  "http://localhost/auth/callback",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://sim.logonvalidation.net/authorize",
			TokenURL: "https://sim.logonvalidation.net/token",
		},
	}
	ctx := context.Background()
	httpClient := &http.Client{Transport: nil}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	date := time.Unix(Time, 0)
	token := oauth2.Token{
		AccessToken:  AccessToken,
		TokenType:    "code",
		RefreshToken: RefreshToken,
		Expiry:       date,
	}
	token.TokenType = "Bearer"
	return conf.Client(ctx, &token), nil
}

func getSaxgoClient() (SaxoClient, error) {
	saxo := SaxoClient{
		Dev:          true,
		Saxoporturl:  Saxodevporturl,
		Saxorefurl:   Saxodevrefurl,
		Saxotradeurl: Saxodevtradeurl,
	}
	client, err := getClient()
	if err != nil {
		return saxo, err
	}

	saxo.Http = client
	return saxo, nil

}

func TestGetMyOrders(t *testing.T) {
	saxo, err := getSaxgoClient()
	if err != nil {
		t.Errorf(err.Error())
	}

	orders, err := saxo.GetMyOrders()
	if err != nil {
		t.Errorf(err.Error())
	}
	if orders.Data[0].AccountId != "16164583" {
		t.Errorf("got %q, wanted %q", orders.Data[0].AccountId, "16164583")
	}
}

func TestBuyStock(t *testing.T) {
	saxo, err := getSaxgoClient()
	if err != nil {
		t.Errorf(err.Error())
	}

	stockRes, err := saxo.BuyStock(18096309, 100, AccountKey)
	if err != nil {
		t.Errorf(err.Error())
	}

	if stockRes.OrderId == "" {
		t.Errorf("got %q, wanted it to not be empty", stockRes.OrderId)
	}

}

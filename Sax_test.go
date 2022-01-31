package saxgo

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/TobiasFP/SaxGo/testcredentials"
	"golang.org/x/oauth2"
)

func getClient() (*http.Client, error) {
	conf := &oauth2.Config{
		RedirectURL:  "http://localhost/auth/callback",
		ClientID:     testcredentials.ClientID,
		ClientSecret: testcredentials.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://sim.logonvalidation.net/authorize",
			TokenURL: "https://sim.logonvalidation.net/token",
		},
	}
	ctx := context.Background()
	httpClient := &http.Client{Transport: nil}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	date := time.Unix(testcredentials.Time, 0)
	token := oauth2.Token{
		AccessToken:  testcredentials.AccessToken,
		TokenType:    "code",
		RefreshToken: testcredentials.RefreshToken,
		Expiry:       date,
	}
	token.TokenType = "Bearer"
	return conf.Client(ctx, &token), nil
}

func getSaxgoClient() (SaxoClient, error) {
	saxo := SaxoClient{
		SaxoUrl:              SaxoSimUrl,
		SimConnectedWithLive: true,
	}

	client, err := getClient()
	if err != nil {
		return saxo, err
	}

	saxo.Http = client

	err = saxo.SetAccountKey()
	if err != nil {
		return saxo, err
	}
	return saxo, nil

}

func TestGetMyOrders(t *testing.T) {
	saxo, err := getSaxgoClient()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	orders, err := saxo.GetMyOrders()
	if err != nil {
		t.Errorf(err.Error())
	}
	if orders.Data[0].AccountId != "16164583" {
		t.Errorf("got %q, wanted %q", orders.Data[0].AccountId, "16164583")
	}
}
func TestGetMyPositions(t *testing.T) {
	saxo, err := getSaxgoClient()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	positions, err := saxo.GetMyPositions()
	if err != nil {
		t.Errorf(err.Error())
	}
	if positions.Data[0].PositionId != "5001843576" {
		t.Errorf("got %q, wanted %q", positions.Data[0].PositionId, "5001843576")
	}
}

func TestSetAccountKey(t *testing.T) {
	saxo, err := getSaxgoClient()
	if err != nil {
		t.Errorf(err.Error())
	}

	err = saxo.SetAccountKey()
	if err != nil {
		t.Errorf(err.Error())
	}
	if saxo.SaxoAccountKey == "" {
		t.Errorf("got %q, wanted it to not be empty", saxo.SaxoAccountKey)
	}
}

func TestSellOrder(t *testing.T) {
	saxo, err := getSaxgoClient()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	stockRes, err := saxo.SellOrder("5001025814")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if stockRes.OrderId == "" {
		t.Errorf("got %q, wanted it to not be empty", stockRes.OrderId)
	}
}

func TestBuyStock(t *testing.T) {
	saxo, err := getSaxgoClient()
	if err != nil {
		t.Errorf(err.Error())
	}

	stockRes, err := saxo.BuyStock(18096309, 100, "USD")
	if err != nil {
		t.Errorf(err.Error())
	}

	if stockRes.OrderId == "" {
		t.Errorf("got %q, wanted it to not be empty", stockRes.OrderId)
	}

}

func TestConvertCashAmountToStockAmount(t *testing.T) {
	uic := 49975 //Car xnas
	amount := 200.00
	expected := amount * 184.49
	saxo, err := getSaxgoClient()
	if err != nil {
		t.Errorf(err.Error())
	}
	stockAmount, err := saxo.ConvertCashAmountToStockAmount(amount, uic, "EUR")
	if err != nil {
		t.Errorf(err.Error())
	}
	if stockAmount != expected {
		t.Errorf("got %v, expected %v", stockAmount, expected)
	}
}

func TestGetStockPrice(t *testing.T) {
	uic := 49975
	expected := 184.49
	saxo, err := getSaxgoClient()
	if err != nil {
		t.Errorf(err.Error())
	}
	price, err := saxo.getStockPriceIncludingCostToBuy(uic, "EUR")
	if err == nil {
		if price != expected {
			t.Errorf("got %v, expected %v", price, expected)
		}
	} else {
		if saxo.isSim() && err.Error() == "stock prices are unavailable in Simulation mode" {
			return
		}

		t.Errorf("Expected market to be closed, but it seams open")

	}
}

func TestMarketOpen(t *testing.T) {

	expectedOpen := true

	saxo, err := getSaxgoClient()
	if err != nil {
		t.Errorf(err.Error())
	}

	ExchangeId := "NASDAQ"

	exchangeResult, err := saxo.Exchange(ExchangeId)
	if err != nil {
		t.Errorf(err.Error())
	}
	now := time.Now()
	for _, exchangeSession := range exchangeResult.ExchangeSessions {
		if exchangeSession.State == "Closed" && exchangeSession.StartTime.Before(now) && exchangeSession.EndTime.After(now) {
			expectedOpen = false
		}
	}

	isOpen, err := saxo.MarketOpen(ExchangeId)
	if err != nil {
		t.Errorf(err.Error())
	}

	if isOpen != expectedOpen {
		if expectedOpen {
			t.Errorf("Expected market to be open")
		} else {
			t.Errorf("Expected market to be closed")
		}
	}
}

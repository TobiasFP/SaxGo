package saxgo

import (
	"context"
	"net/http"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

var (
	clientID           = "83840fbc414641bd88ff0a779573d2dd"
	clientSecret       = "b83e613ca4d24389b4c883ba08c8081d"
	AccessToken        = "eyJhbGciOiJFUzI1NiIsIng1dCI6IjhGQzE5Qjc0MzFCNjNFNTVCNjc0M0QwQTc5MjMzNjZCREZGOEI4NTAifQ.eyJvYWEiOiI3Nzc3MCIsImlzcyI6Im9hIiwiYWlkIjoiMjM2NCIsInVpZCI6IkQwRm9rMUNrWUc2dXhkNlNwZ2w4NFE9PSIsImNpZCI6IkQwRm9rMUNrWUc2dXhkNlNwZ2w4NFE9PSIsImlzYSI6IkZhbHNlIiwidGlkIjoiNzM5MCIsInNpZCI6IjU0OWE2YzljZTVjNzQxNDdhYjYyMjg1ZGYxMjEwNDBmIiwiZGdpIjoiODQiLCJleHAiOiIxNjQzMTM3NjQ4Iiwib2FsIjoiMUYifQ.J81Cvk78jhAb3TWTyNtafJEoTxf0bUeIeiu9mJsRU-mNzrKaLaFzvfEkGYxUxJhtoLk5mY9MSQbW2GszdEkAQw"
	RefreshToken       = "cc2903a9-6969-49e4-a9df-770a2dcfd15a"
	Time         int64 = 1643137648
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

func TestBuyStock(t *testing.T) {
	saxo, err := getSaxgoClient()
	if err != nil {
		t.Errorf(err.Error())
	}

	stockRes, err := saxo.BuyStock(18096309, 100)
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
	stockAmount, _ := saxo.ConvertCashAmountToStockAmount(amount, uic, "EUR")
	if stockAmount != expected {
		t.Errorf("got %v, expected %v", stockAmount, expected)
	}
}

func TestGetStockPrice(t *testing.T) {
	uic := 49975 //Car xnas
	expected := 184.49
	saxo, err := getSaxgoClient()
	if err != nil {
		t.Errorf(err.Error())
	}
	price, err := saxo.getStockPrice(uic, "EUR")
	if err == nil {
		if price != expected {
			t.Errorf("got %v, expected %v", price, expected)
		}
	} else {
		if err.Error() != "market is closed" {
			t.Errorf("Expected market to be closed, but it seams open")
		}
	}
}

func TestIsMarketOpen(t *testing.T) {
	saxo, err := getSaxgoClient()
	if err != nil {
		t.Errorf(err.Error())
	}

	open, err := saxo.IsMarketOpen("NASDAQ")
	if err != nil {
		t.Errorf(err.Error())
	}
	if open == true {
		t.Errorf("Expected market to be closed, but it seams open")
	}

}

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
	AccessToken        = "eyJhbGciOiJFUzI1NiIsIng1dCI6IjhGQzE5Qjc0MzFCNjNFNTVCNjc0M0QwQTc5MjMzNjZCREZGOEI4NTAifQ.eyJvYWEiOiI3Nzc3MCIsImlzcyI6Im9hIiwiYWlkIjoiMjM2NCIsInVpZCI6IkQwRm9rMUNrWUc2dXhkNlNwZ2w4NFE9PSIsImNpZCI6IkQwRm9rMUNrWUc2dXhkNlNwZ2w4NFE9PSIsImlzYSI6IkZhbHNlIiwidGlkIjoiNzM5MCIsInNpZCI6IjI5ZTJiNTg2NDJmYjQyNDM5ODFlNmVkOTgzNjIwYzQyIiwiZGdpIjoiODQiLCJleHAiOiIxNjQyMDY3NzMzIiwib2FsIjoiMUYifQ.7FD-x_R_NZ4uNFiZgQNi40t_sIf6bimGgHkbly2re5Z2WiUz_-MyCWrDvcL5t95A1UetAFYEhyOSZqbJoUOyog"
	RefreshToken       = "1ed82007-8d69-4786-aa67-a63a23fd502b"
	Time         int64 = 1642067733
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

	err = saxo.setAccountKey()
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

	err = saxo.setAccountKey()
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

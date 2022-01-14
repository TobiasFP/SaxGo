package saxgo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/TobiasFP/SaxGo/structs"
)

var Saxodevporturl = "https://gateway.saxobank.com/sim/openapi/port/"
var Saxodevrefurl = "https://gateway.saxobank.com/sim/openapi/ref/"
var Saxodevtradeurl = "https://gateway.saxobank.com/sim/openapi/trade/"

type SaxoClient struct {
	Http           *http.Client
	Dev            bool
	Saxoporturl    string
	Saxorefurl     string
	Saxotradeurl   string
	SaxoAccountKey string
}

func (saxo *SaxoClient) SetAccountKey() error {
	var me structs.AccountResult
	resp, err := saxo.Http.Get(saxo.Saxoporturl + "v1/accounts/me")
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &me)
	if err != nil {
		return err
	}
	saxo.SaxoAccountKey = me.Data[0].AccountKey
	return nil
}

// https://www.developer.saxo/openapi/learn/orders-and-positions
func (saxo SaxoClient) GetMyOrders() (structs.Orders, error) {
	var orders structs.Orders

	resp, err := saxo.Http.Get(saxo.Saxoporturl + "v1/orders/me?fieldGroups=DisplayAndFormat")
	if err != nil {
		return orders, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return orders, err
	}
	err = json.Unmarshal(body, &orders)
	if err != nil {
		return orders, err
	}

	return orders, err
}

func (saxo SaxoClient) BuyStock(uic int, amount float64) (structs.OrderResult, error) {
	var order structs.OrderResult

	stock := structs.TradeOrder{
		Uic:           uic,
		BuySell:       "Buy",
		AssetType:     "Stock",
		Amount:        amount,
		AmountType:    "CashAmount",
		OrderType:     "Market",
		OrderRelation: "StandAlone",
		ManualOrder:   true,
		OrderDuration: struct {
			DurationType string "json:\"DurationType\""
		}{DurationType: "DayOrder"},
		AccountKey: saxo.SaxoAccountKey,
	}

	stockJson, err := json.Marshal(stock)
	if err != nil {
		return order, err
	}

	resp, err := saxo.Http.Post(saxo.Saxotradeurl+"v2/orders", "application/json", bytes.NewBuffer(stockJson))
	if err != nil {
		return order, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return order, err
	}

	err = json.Unmarshal(body, &order)
	if err != nil {
		return order, err
	}

	return order, err
}

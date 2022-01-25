package saxgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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

func (saxo SaxoClient) BuyStock(uic int, amount float64, currency string) (structs.OrderResult, error) {
	var order structs.OrderResult

	stockAmount, err := saxo.ConvertCashAmountToStockAmount(amount, uic, currency)
	if err != nil {
		return order, err
	}

	stock := structs.TradeOrder{
		Uic:           uic,
		BuySell:       "Buy",
		AssetType:     "Stock",
		Amount:        stockAmount,
		AmountType:    "Quantity",
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

func (saxo SaxoClient) ConvertCashAmountToStockAmount(cashAmount float64, stockUic int, currency string) (float64, error) {
	price, err := saxo.getStockPrice(stockUic, currency)
	if err != nil {
		return 0, err
	}
	return cashAmount / price, nil
}

func (saxo SaxoClient) getStockPrice(stockUic int, currency string) (float64, error) {
	var infoPrice structs.InfoPriceResult
	resp, err := saxo.Http.Get(saxo.Saxotradeurl + "v1/infoprices/?FieldGroups=PriceInfo,PriceInfoDetails,Commissions,InstrumentPriceDetails&AssetType=Stock&Amount=1&Uic=" + fmt.Sprint(stockUic))
	if err != nil {
		return 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(body, &infoPrice)
	if err != nil {
		return 0, err
	}

	if !infoPrice.InstrumentPriceDetails.IsMarketOpen {
		return 0, errors.New("market is closed")
	}

	return infoPrice.Quote.Ask, nil
}

func (saxo SaxoClient) IsMarketOpen(ExchangeId string) (bool, error) {
	var exchangeResult structs.ExchangeResult
	resp, err := saxo.Http.Get(saxo.Saxorefurl + "v1/exchanges/" + ExchangeId)
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(body, &exchangeResult)
	if err != nil {
		return false, err
	}

	if exchangeResult.AllDay {
		return true, nil
	}
	now := time.Now()
	for _, exchangeSession := range exchangeResult.ExchangeSessions {
		if exchangeSession.State == "AutomatedTrading" && exchangeSession.StartTime.Before(now) && exchangeSession.EndTime.After(now) {
			return true, nil
		}
	}

	return false, nil
}

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

var SaxoSimUrl = "https://gateway.saxobank.com/sim/openapi/"
var SaxoLiveUrl = "https://gateway.saxobank.com/openapi/"

// The SimConnectedWithLive is to indicate wether your sim account is connected to a live account.
// In this case, your sim account can obtain price info.
type SaxoClient struct {
	Http                 *http.Client
	SaxoUrl              string
	SaxoAccountKey       string
	SimConnectedWithLive bool
}

func (saxo *SaxoClient) SetAccountKey() error {
	var me structs.AccountResult
	resp, err := saxo.Http.Get(saxo.SaxoUrl + "port/v1/accounts/me")
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

	resp, err := saxo.Http.Get(saxo.SaxoUrl + "port/v1/orders/me?fieldGroups=DisplayAndFormat")
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

func (saxo SaxoClient) GetMyPositions() (structs.Positions, error) {
	var positions structs.Positions

	resp, err := saxo.Http.Get(saxo.SaxoUrl + "port/v1/positions/me?fieldGroups=DisplayAndFormat,ExchangeInfo,Greeks,PositionBase,PositionIdOnly,PositionView")
	if err != nil {
		return positions, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return positions, err
	}
	err = json.Unmarshal(body, &positions)
	return positions, err
}

func (saxo SaxoClient) SellOrder(orderID string) (structs.OrderResult, error) {
	var order structs.OrderResult

	orders, err := saxo.GetMyOrders()
	if err != nil {
		return order, err
	}
	for _, saxoOrder := range orders.Data {
		if saxoOrder.OrderId == orderID {
			return order, errors.New("order has not been converted into a position, so cannot sell")
		}
	}

	positions, err := saxo.GetMyPositions()
	if err != nil {
		return order, err
	}
	// This should be optimised to simply call the positions endpoint with a search instaed.
	for _, position := range positions.Data {
		if position.PositionBase.SourceOrderId == orderID {
			if position.PositionBase.Amount <= 0 {
				return order, errors.New("position is already sold, cannot resell")
			}
			order, err = saxo.SellStock(int(position.PositionBase.Uic), position.PositionBase.Amount)
			return order, err
		}
	}

	return order, errors.New("position or Uic not matching, cannot sell")
}

func (saxo SaxoClient) SellStock(uic int, amount float64) (structs.OrderResult, error) {
	var order structs.OrderResult

	stock := structs.TradeOrder{
		Uic:           uic,
		BuySell:       "Sell",
		AssetType:     "Stock",
		Amount:        amount,
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

	resp, err := saxo.Http.Post(saxo.SaxoUrl+"trade/v2/orders", "application/json", bytes.NewBuffer(stockJson))
	if err != nil {
		return order, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return order, err
	}

	restErr, err := structs.GetRestError(body)
	if err != nil {
		return order, errors.New(restErr.ErrorInfo.Message)
	}

	err = json.Unmarshal(body, &order)
	return order, err
}

func (saxo SaxoClient) BuyStock(uic int, amount float64, currency string) (structs.OrderResult, error) {
	var order structs.OrderResult
	stockAmount := amount
	if !saxo.isSim() {
		convertedAmount, err := saxo.ConvertCashAmountToStockAmount(amount, uic, currency)
		if err != nil {
			return order, err
		}

		stockAmount = convertedAmount
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

	resp, err := saxo.Http.Post(saxo.SaxoUrl+"trade/v2/orders", "application/json", bytes.NewBuffer(stockJson))
	if err != nil {
		return order, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return order, err
	}

	err = json.Unmarshal(body, &order)
	return order, err
}

func (saxo SaxoClient) ConvertCashAmountToStockAmount(cashAmount float64, stockUic int, currency string) (float64, error) {
	if saxo.isSim() {
		return 0, errors.New("stock prices are unavailable in Simulation mode")
	}
	price, err := saxo.getStockPriceIncludingCostToBuy(stockUic, currency)
	if err != nil {
		return 0, err
	}
	return cashAmount / price, nil
}

func (saxo SaxoClient) GetInfoPrice(stockUic int, currency string) (structs.InfoPriceResult, error) {
	var infoPrice structs.InfoPriceResult
	if saxo.isSim() {
		return infoPrice, errors.New("stock prices are unavailable in Simulation mode, without a connected live account")
	}

	resp, err := saxo.Http.Get(saxo.SaxoUrl + "trade/v1/infoprices/?FieldGroups=PriceInfo,PriceInfoDetails,Commissions,DisplayAndFormat,InstrumentPriceDetails&AssetType=Stock&Amount=1&Uic=" + fmt.Sprint(stockUic))
	if err != nil {
		return infoPrice, err
	}
	if infoPrice.DisplayAndFormat.Currency != currency {
		return infoPrice, errors.New("You ask for " + currency + " But we can only provide info for " + infoPrice.DisplayAndFormat.Currency)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return infoPrice, err
	}

	err = json.Unmarshal(body, &infoPrice)
	if err != nil {
		return infoPrice, err
	}
	return infoPrice, nil
}

func (saxo SaxoClient) getStockPriceIncludingCostToBuy(stockUic int, currency string) (float64, error) {
	infoprice, err := saxo.GetInfoPrice(stockUic, currency)
	if err != nil {
		return 0, err
	}
	return infoprice.Quote.Mid + infoprice.Commissions.CostBuy, nil
}

func (saxo SaxoClient) Exchange(ExchangeId string) (structs.ExchangeResult, error) {
	var exchangeResult structs.ExchangeResult
	resp, err := saxo.Http.Get(saxo.SaxoUrl + "ref/v1/exchanges/" + ExchangeId)
	if err != nil {
		return exchangeResult, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return exchangeResult, err
	}

	err = json.Unmarshal(body, &exchangeResult)
	return exchangeResult, err
}

func (saxo SaxoClient) MarketOpen(ExchangeId string) (bool, error) {
	exchangeResult, err := saxo.Exchange(ExchangeId)
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

func (saxo SaxoClient) isSim() bool {
	return !saxo.SimConnectedWithLive && saxo.SaxoUrl == SaxoSimUrl
}

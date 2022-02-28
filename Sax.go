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

func (saxo SaxoClient) CancelOrder(id string) (orders structs.Orders, err error) {
	resp, err := saxo.Http.Get(saxo.SaxoUrl + "trade/v2/orders/" + id + "/?AccountKey=" + saxo.SaxoAccountKey)
	if err != nil {
		return orders, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return orders, err
	}
	err = json.Unmarshal(body, &orders)
	return orders, err
}

func (saxo SaxoClient) GetOrderDetails(id string) (order structs.Order, err error) {
	resp, err := saxo.Http.Get(saxo.SaxoUrl + "port/v1/orders/" + id + "/details/?ClientKey=" + saxo.SaxoAccountKey)
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

func (saxo SaxoClient) GetMyPositions() (positions structs.Positions, err error) {
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

func (saxo SaxoClient) SellOrder(orderID string, orderPrice float64) (structs.OrderResult, error) {
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
			order, err = saxo.SellStock(int(position.PositionBase.Uic), position.PositionBase.Amount, "", orderPrice)
			return order, err
		}
	}

	return order, errors.New("position or Uic not matching, cannot sell")
}

// Sells stock.
// PositionId is optional, if PositionId == 0, the selling of a stock will be unrelated to an order
// orderPrice is optional, if orderPrice == 0, the selling of a stock will be done at market price decided by saxo.
func (saxo SaxoClient) SellStock(uic int, amount float64, positionId string, orderPrice float64) (structs.OrderResult, error) {
	stock := structs.TradeOrder{
		Uic:         uic,
		BuySell:     "Sell",
		AssetType:   "Stock",
		Amount:      amount,
		AmountType:  "Quantity",
		OrderPrice:  orderPrice,
		OrderType:   "Limit",
		ManualOrder: true,
		OrderDuration: struct {
			DurationType string "json:\"DurationType\""
		}{DurationType: "DayOrder"},
		AccountKey: saxo.SaxoAccountKey,
	}

	if positionId != "" {
		stock.PositionId = positionId
	}

	return saxo.Trade(stock)
}

func (saxo SaxoClient) BuyStock(uic int, stockAmount float64, orderPrice float64) (order structs.OrderResult, err error) {
	if stockAmount == 0 {
		return order, errors.New("cannot buy 0 shares. you try to invest too little")
	}

	stockOrder := structs.TradeOrder{
		Uic:         uic,
		BuySell:     "Buy",
		AssetType:   "Stock",
		Amount:      stockAmount,
		AmountType:  "Quantity",
		OrderPrice:  orderPrice,
		OrderType:   "Limit",
		ManualOrder: true,
		PositionId:  "",
		OrderDuration: struct {
			DurationType string "json:\"DurationType\""
		}{DurationType: "DayOrder"},
		AccountKey: saxo.SaxoAccountKey,
	}
	return saxo.Trade(stockOrder)
}

func (saxo SaxoClient) BuyCfdOnStock(uic int, cfdAmount float64) (order structs.OrderResult, err error) {
	if cfdAmount == 0 {
		return order, errors.New("cannot buy 0 shares. you try to invest too little")
	}

	cfdOrder := structs.TradeOrder{
		Uic:         uic,
		BuySell:     "Buy",
		AssetType:   "CfdOnStock",
		Amount:      cfdAmount,
		AmountType:  "Quantity",
		OrderType:   "Market",
		ManualOrder: true,
		PositionId:  "",
		OrderDuration: struct {
			DurationType string "json:\"DurationType\""
		}{DurationType: "DayOrder"},
		AccountKey: saxo.SaxoAccountKey,
	}
	return saxo.Trade(cfdOrder)
}

func (saxo SaxoClient) Trade(orderRequest structs.TradeOrder) (order structs.OrderResult, err error) {
	stockJson, err := json.Marshal(orderRequest)
	if err != nil {
		return order, err
	}
	resp, err := saxo.Http.Post(saxo.SaxoUrl+"trade/v2/orders", "application/json", bytes.NewBuffer(stockJson))
	if err != nil {
		return order, err
	}
	return orderResToOrderStruct(resp)
}

func orderResToOrderStruct(resp *http.Response) (order structs.OrderResult, err error) {

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return order, err
	}

	restErr, err := structs.GetRestError(body)
	if err != nil {
		return order, errors.New(restErr.FullMessage)
	}

	validationErr, err := structs.GetValidationError(body)
	if err != nil {
		return order, errors.New(validationErr.FullMessage)
	}

	err = json.Unmarshal(body, &order)
	return order, err
}

func (saxo SaxoClient) GetInfoPrice(stockUic int) (infoPrice structs.InfoPriceResult, err error) {
	if saxo.isSim() {
		return infoPrice, errors.New("stock prices are unavailable in Simulation mode, without a connected live account")
	}

	resp, err := saxo.Http.Get(saxo.SaxoUrl + "trade/v1/infoprices/?FieldGroups=PriceInfo,PriceInfoDetails,Commissions,DisplayAndFormat,InstrumentPriceDetails&AssetType=Stock&Amount=1&Uic=" + fmt.Sprint(stockUic))
	if err != nil {
		return infoPrice, err
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

func (saxo SaxoClient) GetStockPriceIncludingCostToBuy(stockUic int) (float64, error) {
	infoprice, err := saxo.GetInfoPrice(stockUic)
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

func (saxo SaxoClient) GetChart(assetType string, horizon int, stockUic int, date time.Time) (chartRes structs.ChartResult, err error) {
	if saxo.isSim() {
		return chartRes, errors.New("charts are unavailable in Simulation mode, without a connected live account")
	}
	resp, err := saxo.Http.Get(saxo.SaxoUrl + "/chart/v1/charts/?AssetType=" + assetType + "&Horizon=" + fmt.Sprint(horizon) + "&Mode=UpTo&Time=" + date.Format("2006-01-02T15:04:05.000000Z") + "&Uic=" + fmt.Sprint(stockUic))
	if err != nil {
		return chartRes, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return chartRes, err
	}

	err = json.Unmarshal(body, &chartRes)
	if err != nil {
		return chartRes, err
	}
	return chartRes, nil
}

func (saxo SaxoClient) isSim() bool {
	return !saxo.SimConnectedWithLive && saxo.SaxoUrl == SaxoSimUrl
}

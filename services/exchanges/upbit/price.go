package upbit

import (
	"businessLogics/model/price"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"net/http"
	"strings"
	"time"
)

type PriceResp struct {
	Market       string  `json:"market"`
	OpeningPrice float64 `json:"opening_price"`
	TradePrice   float64 `json:"trade_price"`
	TradeVolume  float64 `json:"acc_trade_volume_24h"`
	ChangeRate   float64 `json:"signed_change_rate"`
}
type PricesResp []PriceResp

func (pr PriceResp) splitMarket() (market, symbol string) {
	market, symbol, _ = strings.Cut(pr.Market, "-")
	return
}

func (pr PriceResp) convertFormat() price.Price {
	market, symbol := pr.splitMarket()

	exchangeRate := decimal.NewFromFloat(0.00083)
	priceFloat, _ := decimal.NewFromFloat(pr.TradePrice).Mul(exchangeRate).Float64()

	return price.Price{
		Symbol:        symbol + "/" + market,
		Price:         priceFloat,
		OriginalPrice: pr.TradePrice,
		Change24H:     pr.ChangeRate * 100,
		Volume24:      pr.TradeVolume * priceFloat,
		Logo:          "",
	}
}

type MarketResp struct {
	Market      string `json:"market"`
	KoreanName  string `json:"korean_name"`
	EnglishName string `json:"english_name"`
}
type MarketsResp []MarketResp

func (mr MarketsResp) marketNames(block int) []string {
	bufSize := len(mr)/block + 1
	buffer := make([]bytes.Buffer, bufSize)
	for i, market := range mr {
		buffer[i%bufSize].WriteString(market.Market + ",")
	}

	names := make([]string, bufSize)
	for i := range buffer {
		names[i] = strings.TrimRight(buffer[i].String(), ",")
	}

	return names
}

func Request(url string) (resp *http.Response, err error) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	resp, err = client.Do(req)
	if err != nil {
		return
	}

	return
}

func getMarkets() (markets MarketsResp, err error) {
	resp, err := Request("https://api.upbit.com/v1/market/all")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&markets)
	if err != nil {
		return
	}

	return
}

func GetPrices() {
	defer func() {
		time.Sleep(3 * time.Second)
		GetPrices()
	}()

	markets, err := getMarkets()
	if err != nil {
		return
	}

	marketNames := markets.marketNames(40)
	pricesResp := make(PricesResp, 0)

	for _, names := range marketNames {
		resp, err := Request(fmt.Sprintf("https://api.upbit.com/v1/ticker?markets=%s", names))
		if err != nil {
			return
		}
		defer resp.Body.Close()

		tmpPrices := make(PricesResp, 0)
		err = json.NewDecoder(resp.Body).Decode(&tmpPrices)
		if err != nil {
			return
		}

		pricesResp = append(pricesResp, tmpPrices...)
	}

	prices := make(price.Prices, 0)
	for _, p := range pricesResp {
		prices = append(prices, p.convertFormat())
	}

	err = price.SetAll("upbit", prices)
	if err != nil {
		return
	}
}

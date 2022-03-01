package httpclientserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type btc struct {
	Symbol      string `json:"symbol"`
	PriceChange string `json:"priceChange"`
	LastPrice   string `json:"lastPrice"`
}

func InitHttpClient() {
	resp, err := http.Get("https://api2.binance.com/api/v3/ticker/24hr")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var coinData []btc
	err = json.Unmarshal(body, &coinData)
	if err != nil {
		panic(err)
	}
	for _, v := range coinData {
		fmt.Printf("|%30s|%30s|%30s|\n", v.Symbol, v.PriceChange, v.LastPrice)
	}

}

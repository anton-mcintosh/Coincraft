
package main

import (
	"fmt"
	"net/http"
  "encoding/json"
)

type Money struct {
  Dollars float64
  Coins float64
  Buys int
  Sells int
}

type CoinResponse struct {
  Prices [][]interface{}`json:"prices"`
}
var data CoinResponse

func extractPrices (data CoinResponse) []float64 {
  var prices []float64
  for _, entry := range data.Prices {
    price := entry[1].(float64)
    prices = append(prices, price)
  }
  return prices
}
func main() {

  money := Money{1000, 0, 0, 0}

  var coin string
  var strategy string
  var trade_fee int
  
  fmt.Println("Pick a coin")
  fmt.Scanln(&coin)
  fmt.Println("Pick a strategy")
  fmt.Scanln(&strategy)
  fmt.Println("Specify a trade fee")
  fmt.Scanln(&trade_fee)
  url := "https://api.coingecko.com/api/v3/coins/" + coin + "/market_chart?vs_currency=usd&days=365"
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

  json.NewDecoder(res.Body).Decode(&data)

  prices := extractPrices(data)
  macd, signal := MACD(prices, 12, 26, 9)

  for i := 0; i < len(macd); i++ {
    if macd[i] > signal[i] {
      money.buy(1, prices[i])
    }
  }
}

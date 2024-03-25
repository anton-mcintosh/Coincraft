
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

func main() {

  var coin string
  
  fmt.Println("Pick a coin")
  fmt.Scanln(&coin)
  url := "https://api.coingecko.com/api/v3/coins/" + coin + "/market_chart?vs_currency=usd&days=365"
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

  var data CoinResponse
  json.NewDecoder(res.Body).Decode(&data)

  prices := extractPrices(data)
  macd, signal := MACD(prices, 12, 26, 9)

  for i := 0; i < len(macd); i++ {
    fmt.Println(macd[i], "- - - ", signal[i])
  }
}

func extractPrices(data CoinResponse) []float64 {
  var prices []float64
  for _, entry := range data.Prices {
    price := entry[1].(float64)
    prices = append(prices, price)
  }
  return prices
}

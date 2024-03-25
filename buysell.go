package main

func buy(s.Money, quantity int, price float64) {
  s.Dollars -= quantity * price
  s.Stock += quantity
  s.Buys++
}

func sell(s.Money, quantity int, price float64) {
  s.Dollars += quantity * price
  s.Stock -= quantity
  s.Sells++
}

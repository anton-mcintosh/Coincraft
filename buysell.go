package main

func (s *Money) buy(quantity int, price float64) {
  s.Dollars -= float64(quantity) * price
  s.Coins += float64(quantity)
  s.Buys++
}

func (s *Money) sell (quantity int, price float64) {
  s.Dollars += float64(quantity) * price
  s.Coins -= float64(quantity)
  s.Sells++
}

package main

func EMA(data []float64, period int) []float64 {
  var ema []float64
  var multiplier float64 = 2.0 / float64(period + 1)
  var sum float64 = 0.0
  for i := 0; i < period; i++ {
    sum += data[i]
    ema = append(ema, sum / float64(i + 1))
  }
  for i := period; i < len(data); i++ {
    sum = data[i] * multiplier + sum * (1 - multiplier)
    ema = append(ema, sum)
  }
  return ema
}

package main

func crossover(macd []float64, signal []float64) []float64 {
  cross := make([]float64, len(macd))
  for i := 0; i < len(macd); i++ {
    if macd[i] > signal[i] {
      cross[i] = 1
    } else {
      cross[i] = 0
    }
  }
  return cross
}

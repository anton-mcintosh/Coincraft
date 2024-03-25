package main

func MACD(data []float64, short int, long int, signal int) (macd []float64, signalLine []float64) {
  shortEMA := EMA(data, short)
  longEMA := EMA(data, long)
  macd = make([]float64, len(data))
  for i := 0; i < len(data); i++ {
    macd[i] = shortEMA[i] - longEMA[i]
  }
  signalLine = EMA(macd, signal)
  return macd, signalLine
}

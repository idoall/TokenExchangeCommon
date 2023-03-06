package commonstock

import "github.com/idoall/TokenExchangeCommon/commonstock/container/bst"

// Exponential Moving Average (EMA).
func Ema(period int, values []float64) []float64 {
	result := make([]float64, len(values))

	k := float64(2) / float64(1+period)

	for i, value := range values {
		if i > 0 {
			result[i] = (value * k) + (result[i-1] * float64(1-k))
		} else {
			result[i] = value
		}
	}

	return result
}

func Rma(period int, values []float64) []float64 {
	result := make([]float64, len(values))
	sum := float64(0)

	for i, value := range values {
		count := i + 1

		if i < period {
			sum += value
		} else {
			sum = (result[i-1] * float64(period-1)) + value
			count = period
		}

		result[i] = sum / float64(count)
	}

	return result
}

// Dema calculates the Double Exponential Moving Average (DEMA).
//
// DEMA = (2 * EMA(values)) - EMA(EMA(values))
//
// Returns dema.
func Dema(period int, values []float64) []float64 {
	ema1 := Ema(period, values)
	ema2 := Ema(period, ema1)

	dema := subtract(multiplyBy(ema1, 2), ema2)

	return dema
}

// Moving max for the given period.
func Max(period int, values []float64) []float64 {
	result := make([]float64, len(values))

	buffer := make([]float64, period)
	bst := bst.New()

	for i := 0; i < len(values); i++ {
		bst.Insert(values[i])

		if i >= period {
			bst.Remove(buffer[i%period])
		}

		buffer[i%period] = values[i]
		result[i] = bst.Max().(float64)
	}

	return result
}

// Moving min for the given period.
func Min(period int, values []float64) []float64 {
	result := make([]float64, len(values))

	buffer := make([]float64, period)
	bst := bst.New()

	for i := 0; i < len(values); i++ {
		bst.Insert(values[i])

		if i >= period {
			bst.Remove(buffer[i%period])
		}

		buffer[i%period] = values[i]
		result[i] = bst.Min().(float64)
	}

	return result
}

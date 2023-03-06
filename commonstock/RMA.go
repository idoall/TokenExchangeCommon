package commonstock

import (
	"time"

	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

// Rolling Moving Average (RMA).
type RMA struct {
	Period int //默认13
	data   []RMAPoint
	kline  []*commonmodels.Kline
}

type RMAPoint struct {
	Value float64
	Time  time.Time
}

// NewRMA new Func
func NewRMA(list []*commonmodels.Kline, period int) *RMA {
	m := &RMA{kline: list, Period: period}
	return m
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

// Calculation Func
func (e *RMA) Calculation() *RMA {
	sum := float64(0)

	for i, v := range e.kline {
		count := i + 1

		if i < e.Period {
			sum += v.Close
		} else {
			sum = (e.data[i-1].Value * float64(e.Period-1)) + v.Close
			count = e.Period
		}

		e.data = append(e.data, RMAPoint{
			Value: sum / float64(count),
			Time:  v.KlineTime,
		},
		)
	}

	return e
}

// GetPoints Func
func (e *RMA) GetPoints() []RMAPoint {
	return e.data
}

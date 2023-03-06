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

// Calculation Func
func (e *RMA) Calculation() *RMA {

	closeing := make([]float64, len(e.kline))
	for _, v := range e.kline {
		closeing = append(closeing, v.Close)
	}

	rams := Rma(e.Period, closeing)

	for i := 0; i < len(rams); i++ {
		e.data = append(e.data, RMAPoint{
			Time:  e.kline[i].KlineTime,
			Value: rams[i],
		})
	}

	return e
}

// GetPoints Func
func (e *RMA) GetPoints() []RMAPoint {
	return e.data
}

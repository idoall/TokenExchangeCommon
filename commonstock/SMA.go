package commonstock

import (
	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

// SMA struct
type SMA struct {
	Period int //默认计算几天的MA,KDJ一般是9，OBV是10、20、30
	points []SMAPoint
	kline  []*commonmodels.Kline
}

type SMAPoint struct {
	point
}

// NewSMA new Func
func NewSMA(list []*commonmodels.Kline, period int) *SMA {
	m := &SMA{kline: list, Period: period}
	return m
}

// Calculation Func
func (e *SMA) Calculation() *SMA {
	for i := 0; i < len(e.kline); i++ {
		var smaPointStruct SMAPoint
		if i > e.Period-1 {
			var sum float64
			for j := i; j >= (i - (e.Period - 1)); j-- {

				sum += e.kline[j].Close
			}
			smaPointStruct.Value = (+(sum / float64(e.Period)))
			// e.Value = append(e.Value, +(sum / e.Period))
		} else {
			smaPointStruct.Value = 0.0
			// e.Value = append(e.Value, 0.0)
		}

		smaPointStruct.Time = e.kline[i].KlineTime
		e.points = append(e.points, smaPointStruct)
	}
	return e
}

// GetPoints Func
func (e *SMA) GetPoints() []SMAPoint {
	return e.points
}

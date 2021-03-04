package commonstock

import (
	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

// MA struct
type MA struct {
	Period int //默认计算几天的MA,KDJ一般是9，OBV是10、20、30
	points []MAPoint
	kline  []*commonmodels.Kline
}

type MAPoint struct {
	point
}

// NewMA new Func
func NewMA(list []*commonmodels.Kline, period int) *MA {
	m := &MA{kline: list, Period: period}
	return m
}

// GetPoints return Point
func (e *MA) GetPoints() []MAPoint {
	return e.points
}

// Calculation Func
func (e *MA) Calculation() *MA {
	for i := 0; i < len(e.kline); i++ {
		if i < e.Period-1 {
			p := MAPoint{}
			p.Time = e.kline[i].KlineTime
			p.Value = 0.0
			e.points = append(e.points, p)
			continue
		}
		var sum float64
		for j := 0; j < e.Period; j++ {

			sum += e.kline[i-j].Close
		}

		p := MAPoint{}
		p.Time = e.kline[i].KlineTime
		p.Value = sum / float64(e.Period)
		e.points = append(e.points, p)
	}
	return e
}

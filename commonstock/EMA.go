package commonstock

import (
	"time"

	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

// EMA struct
type EMA struct {
	Period int //默认计算几天的EMA
	points []emaPoint
	kline  []*commonmodels.Kline
}

type emaPoint struct {
	point
}

// NewEMA new Func
func NewEMA(list []*commonmodels.Kline, period int) *EMA {
	m := &EMA{kline: list, Period: period}
	return m
}

// Calculation Func
func (e *EMA) Calculation() {
	for _, v := range e.kline {
		e.Add(v.KlineTime, v.Close)
	}
}

// GetPoints return Point
func (e *EMA) GetPoints() []emaPoint {
	return e.points
}

// Add adds a new Value to Ema
func (e *EMA) Add(timestamp time.Time, value float64) {
	p := emaPoint{}
	p.Time = timestamp

	//平滑指数，一般取作2/(N+1)
	alpha := 2.0 / (float64(e.Period) + 1.0)

	// fmt.Println(alpha)

	emaTminusOne := value
	if len(e.points) > 0 {
		emaTminusOne = e.points[len(e.points)-1].Value
	}

	emaT := alpha*value + (1-alpha)*emaTminusOne
	p.Value = emaT
	e.points = append(e.points, p)
}

package commonstock

import (
	"time"

	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

// ProjectionOscillator struct
type ProjectionOscillator struct {
	Period int //默认一般是50
	Smooth int // 默认一般是1
	data   []ProjectionOscillatorPoint
	kline  []*commonmodels.Kline
}

// ProjectionOscillatorPoint 投影振荡器策略
// 在po高于spo时提供买入操作，在po低于spo时提供卖出操作。
type ProjectionOscillatorPoint struct {
	Time time.Time
	Po   float64
	Spo  float64
}

// NewProjectionOscillator new Func
func NewProjectionOscillator(list []*commonmodels.Kline, period, smooth int) *ProjectionOscillator {
	m := &ProjectionOscillator{kline: list, Period: period, Smooth: smooth}
	return m
}

// Calculation Func
func (e *ProjectionOscillator) Calculation() *ProjectionOscillator {

	period := e.Period
	smooth := e.Smooth
	var high, low, closing []float64
	for _, v := range e.kline {
		high = append(high, v.High)
		low = append(low, v.Low)
		closing = append(closing, v.Close)
	}

	x := generateNumbers(0, float64(len(closing)), 1)
	mHigh, _ := MovingLeastSquare(period, x, high)
	mLow, _ := MovingLeastSquare(period, x, low)

	vHigh := add(high, multiply(mHigh, x))
	vLow := add(low, multiply(mLow, x))

	pu := Max(period, vHigh)
	pl := Min(period, vLow)

	po := divide(multiplyBy(subtract(closing, pl), 100), subtract(pu, pl))
	spo := Ema(smooth, po)

	for i := 0; i < len(po); i++ {
		e.data = append(e.data, ProjectionOscillatorPoint{
			Time: e.kline[i].KlineTime,
			Po:   po[i],
			Spo:  spo[i],
		})
	}
	return e
}

// GetPoints Func
func (e *ProjectionOscillator) GetPoints() []ProjectionOscillatorPoint {
	return e.data
}

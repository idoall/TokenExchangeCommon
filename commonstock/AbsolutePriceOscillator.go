package commonstock

import (
	"time"

	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

// AbsolutePriceOscillator struct
type AbsolutePriceOscillator struct {
	FastPeriod int // 默认一般是14
	SlowPeriod int // 默认一般是30
	data       []AbsolutePriceOscillatorPoint
	kline      []*commonmodels.Kline
}

// AbsolutePriceOscillatorPoint 绝对价格震荡指标 (APO)
// AbsolutePriceOscillator函数计算用于跟踪趋势的技术指标。APO 上穿零表示看涨，而下穿零表示看跌。正值表示上升趋势，负值表示下降趋势。
type AbsolutePriceOscillatorPoint struct {
	Time  time.Time
	Value float64
}

// NewAbsolutePriceOscillator new Func
func NewAbsolutePriceOscillator(list []*commonmodels.Kline) *AbsolutePriceOscillator {
	m := &AbsolutePriceOscillator{kline: list}
	return m
}

// Calculation Func
func (e *AbsolutePriceOscillator) Calculation() *AbsolutePriceOscillator {

	var closing []float64
	for _, v := range e.kline {
		closing = append(closing, v.Close)
	}

	fast := Ema(e.FastPeriod, closing)
	slow := Ema(e.SlowPeriod, closing)
	apo := subtract(fast, slow)

	for i := 0; i < len(apo); i++ {
		e.data = append(e.data, AbsolutePriceOscillatorPoint{
			Time:  e.kline[i].KlineTime,
			Value: apo[i],
		})
	}
	return e
}

// GetPoints Func
func (e *AbsolutePriceOscillator) GetPoints() []AbsolutePriceOscillatorPoint {
	return e.data
}

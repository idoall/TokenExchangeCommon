package commonstock

import (
	"time"

	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

// EMA struct
type EMA struct {
	Period int //默认计算几天的EMA
	points []EMAPoint
	kline  []*commonmodels.Kline
}

type EMAPoint struct {
	Value float64
	Time  time.Time
}

// NewEMA new Func
func NewEMA(list []*commonmodels.Kline, period int) *EMA {
	m := &EMA{kline: list, Period: period}
	return m
}

// Calculation Func
func (e *EMA) Calculation() *EMA {

	closeing := make([]float64, len(e.kline))
	for _, v := range e.kline {
		closeing = append(closeing, v.Close)
	}

	var emas = Ema(e.Period, closeing)

	for i := 0; i < len(emas); i++ {
		e.points = append(e.points, EMAPoint{
			Time:  e.kline[i].KlineTime,
			Value: emas[i],
		})
	}
	return e
}

// GetPoints return Point
func (e *EMA) GetPoints() []EMAPoint {
	return e.points
}

// Add adds a new Value to Ema
// 使用方法，先添加最早日期的数据,最后一条应该是当前日期的数据，结果与 AICoin 对比完全一致
// func (e *EMA) Add(timestamp time.Time, value float64) {
// 	p := EMAPoint{}
// 	p.Time = timestamp

// 	//平滑指数，一般取作2/(N+1)
// 	alpha := 2.0 / (float64(e.Period) + 1.0)

// 	// fmt.Println(alpha)

// 	emaTminusOne := value
// 	if len(e.points) > 0 {
// 		emaTminusOne = e.points[len(e.points)-1].Value
// 	}

// 	// 计算 EMA指数
// 	emaT := alpha*value + (1-alpha)*emaTminusOne
// 	p.Value = emaT
// 	e.points = append(e.points, p)
// }

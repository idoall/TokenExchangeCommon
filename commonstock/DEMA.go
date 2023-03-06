package commonstock

import (
	"time"

	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

// DEMA struct
type DEMA struct {
	Period int //默认计算几天的DEMA
	points []DEMAPoint
	kline  []*commonmodels.Kline
}

// DEMAPoint Dema函数计算给定期间的双指数移动平均线 (DEMA)。

// 双指数移动平均线 (DEMA) 是由 Patrick Mulloy 引入的技术指标。目的是减少技术交易者使用的价格图表中存在的噪音量。DEMA 使用两个指数移动平均线 (EMA) 来消除滞后。当价格高于平均水平时，它有助于确认上升趋势，当价格低于平均水平时，它有助于确认下降趋势。当价格超过平均线时，可能表示趋势发生变化。
type DEMAPoint struct {
	Value float64
	Time  time.Time
}

// NewDEMA new Func
func NewDEMA(list []*commonmodels.Kline, period int) *DEMA {
	m := &DEMA{kline: list, Period: period}
	return m
}

// Calculation Func
func (e *DEMA) Calculation() *DEMA {

	closeing := make([]float64, len(e.kline))
	for _, v := range e.kline {
		closeing = append(closeing, v.Close)
	}

	var DEMAs = Dema(e.Period, closeing)

	for i := 0; i < len(DEMAs); i++ {
		e.points = append(e.points, DEMAPoint{
			Time:  e.kline[i].KlineTime,
			Value: DEMAs[i],
		})
	}
	return e
}

// GetPoints return Point
func (e *DEMA) GetPoints() []DEMAPoint {
	return e.points
}

// Add adds a new Value to DEMA
// 使用方法，先添加最早日期的数据,最后一条应该是当前日期的数据，结果与 AICoin 对比完全一致
// func (e *DEMA) Add(timestamp time.Time, value float64) {
// 	p := DEMAPoint{}
// 	p.Time = timestamp

// 	//平滑指数，一般取作2/(N+1)
// 	alpha := 2.0 / (float64(e.Period) + 1.0)

// 	// fmt.Println(alpha)

// 	DEMATminusOne := value
// 	if len(e.points) > 0 {
// 		DEMATminusOne = e.points[len(e.points)-1].Value
// 	}

// 	// 计算 DEMA指数
// 	DEMAT := alpha*value + (1-alpha)*DEMATminusOne
// 	p.Value = DEMAT
// 	e.points = append(e.points, p)
// }

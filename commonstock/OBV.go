package commonstock

// OBV计算方法：
// 主公式：当日OBV=前一日OBV+今日成交量
// 1.基期OBV值为0，即该股上市的第一天，OBV值为0
// 2.若当日收盘价＞上日收盘价，则当日OBV=前一日OBV＋今日成交量
// 3.若当日收盘价＜上日收盘价，则当日OBV=前一日OBV－今日成交量
// 4.若当日收盘价＝上日收盘价，则当日OBV=前一日OBV

import (
	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

// OBV struct
type OBV struct {
	points []obvPoint
	kline  []*commonmodels.Kline
}

type obvPoint struct {
	point
}

// NewOBV new OBV
func NewOBV(list []*commonmodels.Kline) *OBV {
	m := &OBV{kline: list}
	return m
}

// Calculation Func
func (e *OBV) Calculation() *OBV {
	for i := 0; i < len(e.kline); i++ {
		item := e.kline[i]
		var value float64

		//由于OBV的计算方法过于简单化，所以容易受到偶然因素的影响，为了提高OBV的准确性，可以采取多空比率净额法对其进行修正。
		//多空比率净额= [（收盘价－最低价）－（最高价-收盘价）] ÷（ 最高价－最低价）×V
		// value = ((item.Close - item.Low) - (item.High - item.Close)) / (item.High - item.Close) * item.Vol

		if i-1 == -1 {
			value = 0
		} else if item.Close > e.kline[i-1].Close {
			value = e.points[i-1].Value + item.Vol
		} else if item.Close < e.kline[i-1].Close {
			value = e.points[i-1].Value - item.Vol
		} else if item.Close == e.kline[i-1].Close {
			value = e.points[i-1].Value
		}
		var p obvPoint
		p.Value = value
		p.Time = item.KlineTime
		e.points = append(e.points, p)
	}
	return e
}

// GetPoints Func
func (e *OBV) GetPoints() []obvPoint {
	return e.points
}

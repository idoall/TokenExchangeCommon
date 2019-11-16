package commonstock

import (
	"time"

	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

/*
1、计算移动平均值（EMA）
12日EMA的算式为
EMA（12）=前一日EMA（12）×11/13+今日收盘价×2/13
26日EMA的算式为
EMA（26）=前一日EMA（26）×25/27+今日收盘价×2/27
2、计算离差值（DIF）
DIF=今日EMA（12）－今日EMA（26）
3、计算DIF的9日EMA
根据离差值计算其9日的EMA，即离差平均值，是所求的MACD值。为了不与指标原名相混淆，此值又名
DEA或DEM。
今日DEA（MACD）=前一日DEA×8/10+今日DIF×2/10计算出的DIF和DEA的数值均为正值或负值。
用（DIF-DEA）×2即为MACD柱状图。
*/

// MACD is the main object
type MACD struct {
	PeriodShort  int //默认12
	PeriodSignal int //信号长度默认9
	PeriodLong   int //默认26
	points       []macdPoint
	kline        []*commonmodels.Kline
}

type macdPoint struct {
	Time time.Time
	DIF  float64
	DEA  float64
	MACD float64
}

// NewMACD new Func
// 使用方法，先添加最早日期的数据,最后一条应该是当前日期的数据，结果与 AICoin 对比完全一致
func NewMACD(list []*commonmodels.Kline) *MACD {
	m := &MACD{PeriodShort: 12, PeriodSignal: 9, PeriodLong: 26, kline: list}
	return m
}

// Calculation Func
func (e *MACD) Calculation() *MACD {

	emaShort := NewEMA(e.kline, e.PeriodShort).Calculation().GetPoints()
	emaLong := NewEMA(e.kline, e.PeriodLong).Calculation().GetPoints()
	//计算DIF
	for i := 0; i < len(e.kline); i++ {
		dif := emaShort[i].Value - emaLong[i].Value
		e.points = append(e.points, macdPoint{DIF: dif, Time: emaShort[i].Time})
	}

	//临时变量，用于计算DEA
	var difTempKline []*commonmodels.Kline
	for _, v := range e.points {
		difTempKline = append(difTempKline, &commonmodels.Kline{KlineTime: v.Time, Close: v.DIF})
	}
	deaEMA := NewEMA(difTempKline, e.PeriodSignal).Calculation().GetPoints()

	//将DEA并入point，同时计算MACD
	for i := 0; i < len(e.points); i++ {
		e.points[i].DEA = deaEMA[i].Value
		e.points[i].MACD = (e.points[i].DIF - e.points[i].DEA) * 2
	}
	return e
}

// GetPoints return Point
func (e *MACD) GetPoints() []macdPoint {
	return e.points
}

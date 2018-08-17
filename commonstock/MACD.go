package commonstock

import (
	"time"

	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

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

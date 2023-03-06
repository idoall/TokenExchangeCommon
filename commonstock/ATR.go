package commonstock

import (
	"math"
	"time"

	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

// ATR struct
type ATR struct {
	Period int //默认计算几天的MA,KDJ一般是9，OBV是10、20、30
	data   []ATRPoint
	kline  []*commonmodels.Kline
}

type ATRPoint struct {
	Time time.Time
	TR   float64
	ATR  float64
}

// NewATR new Func
func NewATR(list []*commonmodels.Kline, period int) *ATR {
	m := &ATR{kline: list, Period: period}
	return m
}

// Calculation Func
func (e *ATR) Calculation() *ATR {

	for i := 0; i < len(e.kline); i++ {
		klineItem := e.kline[i]
		var ATRPointStruct ATRPoint
		ATRPointStruct.TR = math.Max(klineItem.High-klineItem.Low, math.Max(klineItem.High-klineItem.Close, klineItem.Close-klineItem.Low))
		ATRPointStruct.Time = e.kline[i].KlineTime
		e.data = append(e.data, ATRPointStruct)
	}

	var tempKline []*commonmodels.Kline
	for _, v := range e.data {
		tempKline = append(tempKline, &commonmodels.Kline{
			Close: v.TR,
		})
	}

	var atr = NewEMA(tempKline, e.Period).Calculation().GetPoints()
	for i := 0; i < len(atr); i++ {
		e.data[i].ATR = atr[i].Value
	}
	return e
}

// GetPoints Func
func (e *ATR) GetPoints() []ATRPoint {
	return e.data
}

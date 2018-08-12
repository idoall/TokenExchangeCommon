package commonstock

import (
	"fmt"
	"log"
	"time"

	"github.com/idoall/TokenExchangeCommon/commonmodels"
	"github.com/idoall/TokenExchangeCommon/commonutils"
)

// SMA struct
type SMA struct {
	Period int //默认计算几天的MA,KDJ一般是9，OBV是10、20、30
	Value  []float64
	Time   []time.Time
	kline  []*commonmodels.Kline
}

// NewSMA new Func
func NewSMA(list []*commonmodels.Kline, period int) *SMA {
	m := &SMA{kline: list, Period: period}
	return m
}

// Calculation Func
func (e *SMA) Calculation() {
	for i := 0; i < len(e.kline); i++ {
		if i > e.Period-1 {
			var sum float64
			for j := i; j >= (i - (e.Period - 1)); j-- {

				sum += e.kline[j].Close
			}
			tempValue, err := commonutils.FloatFromString(fmt.Sprintf("%d", e.Period))
			if err != nil {
				log.Panicf("MA commonutils.FloatFromString() Err:%s", err.Error())
			}
			e.Value = append(e.Value, +(sum / tempValue))
		} else {
			e.Value = append(e.Value, 0.0)
		}

		e.Time = append(e.Time, e.kline[i].KlineTime)
	}
}

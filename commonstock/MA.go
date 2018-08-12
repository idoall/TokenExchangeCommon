package commonstock

import (
	"fmt"
	"log"

	"github.com/idoall/TokenExchangeCommon/commonutils"
)

// MA struct
type MA struct {
	Period  int //默认计算几天的MA,KDJ一般是9，OBV是10、20、30
	Value   []float64
	inValue []float64
}

// NewMA new Func
func NewMA(value []float64, period int) *MA {
	m := &MA{inValue: value, Period: period}
	return m
}

// Calculation Func
func (e *MA) Calculation() *MA {
	for i := 0; i < len(e.inValue); i++ {
		if i < e.Period-1 {
			e.Value = append(e.Value, 0.0)
			continue
		}
		var sum float64
		for j := 0; j < e.Period; j++ {

			sum += e.inValue[i-j]
		}

		tempValue, err := commonutils.FloatFromString(fmt.Sprintf("%d", e.Period))
		if err != nil {
			log.Panicf("MA commonutils.FloatFromString() Err:%s", err.Error())
		}
		e.Value = append(e.Value, +(sum / tempValue))
	}
	return e
}

package commonstock

import (
	"math"
	"time"

	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

// Vortex struct
type Vortex struct {
	Period int //默认一般是13
	data   []VortexPoint
	kline  []*commonmodels.Kline
}

type VortexPoint struct {
	Time    time.Time
	PlusVi  float64
	MinusVi float64
}

// NewVortex new Func
func NewVortex(list []*commonmodels.Kline, period int) *Vortex {
	m := &Vortex{kline: list, Period: period}
	return m
}

// Calculation Func
func (e *Vortex) Calculation() *Vortex {

	period := e.Period
	var high, low, closing []float64
	for _, v := range e.kline {
		high = append(high, v.High)
		low = append(low, v.Low)
		closing = append(closing, v.Close)
	}

	plusVi := make([]float64, len(high))
	minusVi := make([]float64, len(high))

	plusVm := make([]float64, period)
	minusVm := make([]float64, period)
	tr := make([]float64, period)

	var plusVmSum, minusVmSum, trSum float64

	for i := 1; i < len(high); i++ {
		j := i % period

		plusVmSum -= plusVm[j]
		plusVm[j] = math.Abs(high[i] - low[i-1])
		plusVmSum += plusVm[j]

		minusVmSum -= minusVm[j]
		minusVm[j] = math.Abs(low[i] - high[i-1])
		minusVmSum += minusVm[j]

		highLow := high[i] - low[i]
		highPrevClosing := math.Abs(high[i] - closing[i-1])
		lowPrevClosing := math.Abs(low[i] - closing[i-1])

		trSum -= tr[j]
		tr[j] = math.Max(highLow, math.Max(highPrevClosing, lowPrevClosing))
		trSum += tr[j]

		plusVi[i] = plusVmSum / trSum
		minusVi[i] = minusVmSum / trSum
	}

	for i := 0; i < len(plusVi); i++ {
		e.data = append(e.data, VortexPoint{
			Time:    e.kline[i].KlineTime,
			MinusVi: minusVi[i],
			PlusVi:  plusVi[i],
		})
	}
	return e
}

// GetPoints Func
func (e *Vortex) GetPoints() []VortexPoint {
	return e.data
}

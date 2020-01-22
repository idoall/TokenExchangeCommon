package commonstock

import (
	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

// RSI is the main object
type RSI struct {
	Period int //默认计算几天的
	points []RSIPoint
	kline  []*commonmodels.Kline
}

type RSIPoint struct {
	point
}

// NewRSI new Func
// 使用方法，先添加最早日期的数据,最后一条应该是当前日期的数据，结果与 AICoin 对比完全一致
func NewRSI(list []*commonmodels.Kline, period int) *RSI {
	m := &RSI{kline: list, Period: period}
	return m
}

// Calculation Func
func (e *RSI) Calculation() *RSI {
	var closeArray []float64
	for _, v := range e.kline {
		closeArray = append(closeArray, v.Close)
	}
	rsiArray := e.rsi(closeArray, e.Period)

	rsiArrayLen := len(rsiArray)
	for i := 0; i <= (rsiArrayLen - 1); i++ {
		var p RSIPoint
		p.Time = e.kline[i].KlineTime
		p.Value = rsiArray[i]
		e.points = append(e.points, p)
	}
	return e
}

// GetPoints return Point
func (e *RSI) GetPoints() []RSIPoint {
	return e.points
}

// GetValue return Value
func (e *RSI) GetValue() []float64 {
	var result []float64
	for _, v := range e.points {
		result = append(result, v.Value)
	}
	return result
}

func (e *RSI) rsi(inReal []float64, inTimePeriod int) []float64 {

	outReal := make([]float64, len(inReal))

	if len(inReal) < inTimePeriod {
		return outReal
	}

	if inTimePeriod < 2 {
		return outReal
	}

	// variable declarations
	tempValue1 := 0.0
	tempValue2 := 0.0
	outIdx := inTimePeriod
	today := 0
	prevValue := inReal[today]
	prevGain := 0.0
	prevLoss := 0.0
	today++

	for i := inTimePeriod; i > 0; i-- {
		tempValue1 = inReal[today]
		today++
		tempValue2 = tempValue1 - prevValue
		prevValue = tempValue1
		if tempValue2 < 0 {
			prevLoss -= tempValue2
		} else {
			prevGain += tempValue2
		}
	}

	prevLoss /= float64(inTimePeriod)
	prevGain /= float64(inTimePeriod)

	if today > 0 {

		tempValue1 = prevGain + prevLoss
		if !((-0.00000000000001 < tempValue1) && (tempValue1 < 0.00000000000001)) {
			outReal[outIdx] = 100.0 * (prevGain / tempValue1)
		} else {
			outReal[outIdx] = 0.0
		}
		outIdx++

	} else {

		for today < 0 {
			tempValue1 = inReal[today]
			tempValue2 = tempValue1 - prevValue
			prevValue = tempValue1
			prevLoss *= float64(inTimePeriod - 1)
			prevGain *= float64(inTimePeriod - 1)
			if tempValue2 < 0 {
				prevLoss -= tempValue2
			} else {
				prevGain += tempValue2
			}
			prevLoss /= float64(inTimePeriod)
			prevGain /= float64(inTimePeriod)
			today++
		}
	}

	for today < len(inReal) {

		tempValue1 = inReal[today]
		today++
		tempValue2 = tempValue1 - prevValue
		prevValue = tempValue1
		prevLoss *= float64(inTimePeriod - 1)
		prevGain *= float64(inTimePeriod - 1)
		if tempValue2 < 0 {
			prevLoss -= tempValue2
		} else {
			prevGain += tempValue2
		}
		prevLoss /= float64(inTimePeriod)
		prevGain /= float64(inTimePeriod)
		tempValue1 = prevGain + prevLoss
		if !((-0.00000000000001 < tempValue1) && (tempValue1 < 0.00000000000001)) {
			outReal[outIdx] = 100.0 * (prevGain / tempValue1)
		} else {
			outReal[outIdx] = 0.0
		}
		outIdx++
	}

	return outReal
}

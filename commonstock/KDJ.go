package commonstock

import (
	"time"

	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

// KDJ is the main object
type KDJ struct {
	Period int //默认计算几天的
	points []kdjPoint
	kline  []*commonmodels.Kline
}

type kdjPoint struct {
	Time time.Time
	RSV  float64
	K    float64
	D    float64
	J    float64
}

// NewKDJ new Func
func NewKDJ(list []*commonmodels.Kline, period int) *KDJ {
	m := &KDJ{kline: list, Period: period}
	return m
}

// Calculation Func
func (e *KDJ) Calculation() *KDJ {
	//计算 rsv , k , d
	rsv, k, d := e.calculationKD(e.kline)
	arrayLen := len(rsv)
	for i := 0; i < arrayLen; i++ {
		e.points = append(e.points, kdjPoint{
			RSV:  rsv[i],
			K:    k[i],
			D:    d[i],
			Time: e.kline[i].KlineTime,
		})
	}
	j := e.calculationJ()
	for i := 0; i < arrayLen; i++ {
		e.points[i].J = j[i]
	}
	return e
}

// GetPoints return Point
func (e *KDJ) GetPoints() []kdjPoint {
	return e.points
}

// GetListK Func
func (e *KDJ) GetListK() []float64 {
	var result []float64
	for _, v := range e.points {
		result = append(result, v.K)
	}
	return result
}

// GetListD Func
func (e *KDJ) GetListD() []float64 {
	var result []float64
	for _, v := range e.points {
		result = append(result, v.D)
	}
	return result
}

// GetListJ Func
func (e *KDJ) GetListJ() []float64 {
	var result []float64
	for _, v := range e.points {
		result = append(result, v.J)
	}
	return result
}

// calculationKD 计算出kd值
func (e *KDJ) calculationKD(records []*commonmodels.Kline) (rsv, k, d []float64) {

	var periodLowArr, periodHighArr []float64
	length := len(records)
	rsv = make([]float64, length)
	k = make([]float64, length)
	d = make([]float64, length)

	// Loop through the entire array.
	for i := 0; i < length; i++ {
		// add points to the array.
		periodLowArr = append(periodLowArr, records[i].Low)
		periodHighArr = append(periodHighArr, records[i].High)

		// 1: Check if array is "filled" else create null point in line.
		// 2: Calculate average.
		// 3: Remove first value.

		if e.Period == len(periodLowArr) {
			lowest := e.arrayLowest(periodLowArr)
			highest := e.arrayHighest(periodHighArr)
			//logger.Infoln(i, records[i].Close, lowest, highest)
			if highest-lowest < 0.000001 {
				rsv[i] = 100
			} else {
				rsv[i] = (records[i].Close - lowest) / (highest - lowest) * 100
			}

			// k[i] = (rsv[i] + 2.0*k[i-1]) / 3
			// d[i] = (k[i] + 2.0*d[i-1]) / 3
			k[i] = (2.0/3)*k[i-1] + 1.0/3*rsv[i]
			d[i] = (2.0/3)*d[i-1] + 1.0/3*k[i]
			// remove first value in array.
			periodLowArr = periodLowArr[1:]
			periodHighArr = periodHighArr[1:]
		} else {
			k[i] = 50
			d[i] = 50
			rsv[i] = 0
		}
	}
	return rsv, k, d
	// _kdj.RSV = rsv
	// _kdj.K = k
	// _kdj.D = d
}

// calculationJ 计算J值
func (e *KDJ) calculationJ() []float64 {
	length := len(e.points)
	var j []float64 = make([]float64, length)

	// Loop through the entire array.
	for i := 0; i < length; i++ {
		item := e.points[i]
		j[i] = 3*item.K - 2*item.D

	}
	return j
}

func (e *KDJ) highest(priceArray []float64, periods int) []float64 {
	var periodArr []float64
	length := len(priceArray)
	var HighestLine []float64 = make([]float64, length)

	// Loop through the entire array.
	for i := 0; i < length; i++ {
		// add points to the array.
		periodArr = append(periodArr, priceArray[i])
		// 1: Check if array is "filled" else create null point in line.
		// 2: Calculate average.
		// 3: Remove first value.
		if periods == len(periodArr) {
			HighestLine[i] = e.arrayHighest(periodArr)

			// remove first value in array.
			periodArr = periodArr[1:]
		} else {
			HighestLine[i] = 0
		}
	}

	return HighestLine
}

func (e *KDJ) lowest(priceArray []float64, periods int) []float64 {
	var periodArr []float64
	length := len(priceArray)
	var LowestLine []float64 = make([]float64, length)

	// Loop through the entire array.
	for i := 0; i < length; i++ {
		// add points to the array.
		periodArr = append(periodArr, priceArray[i])
		// 1: Check if array is "filled" else create null point in line.
		// 2: Calculate average.
		// 3: Remove first value.
		if periods == len(periodArr) {
			LowestLine[i] = e.arrayLowest(periodArr)

			// remove first value in array.
			periodArr = periodArr[1:]
		} else {
			LowestLine[i] = 0
		}
	}

	return LowestLine
}

func (e *KDJ) arrayLowest(priceArray []float64) float64 {
	length := len(priceArray)
	var lowest = priceArray[0]

	// Loop through the entire array.
	for i := 1; i < length; i++ {
		if priceArray[i] < lowest {
			lowest = priceArray[i]
		}
	}

	return lowest
}

func (e *KDJ) arrayHighest(priceArray []float64) float64 {
	length := len(priceArray)
	var highest = priceArray[0]

	// Loop through the entire array.
	for i := 1; i < length; i++ {
		if priceArray[i] > highest {
			highest = priceArray[i]
		}
	}

	return highest
}

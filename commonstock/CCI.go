package commonstock

import (
	"fmt"
	"math"

	"github.com/idoall/TokenExchangeCommon/commonmodels"
	"github.com/idoall/TokenExchangeCommon/commonutils"
)

/*
TYP:=(HIGH+LOW+CLOSE)/3;
       CCI:(TYP-MA(TYP,N))/(0.015*AVEDEV(TYP,N));
TYP比较容易理解，（最高价+最低价+收盘价）÷3
MA(TYP,N) 也比较简单，就是N天的TYP的平均值
AVEDEV(TYP,N) 比较难理解，是对TYP进行绝对平均偏差的计算。
也就是说N天的TYP减去MA(TYP,N)的绝对值的和的平均值。
表达式：
MA = MA(TYP,N)
AVEDEV(TYP,N) =( | 第N天的TYP - MA |   +  | 第N-1天的TYP - MA | + ...... + | 第1天的TYP - MA | ) ÷ N
CCI = （TYP－MA）÷ AVEDEV(TYP,N)   ÷0.015

计算商品通道指数有几个步骤。
以下示例适用于典型的20周期cci：
cci =（典型价格 - tp的20周期平均值）/（.015 x平均偏差）
典型价格（tp）=（高+低+近）/ 3
常数= 0.015
出于缩放目的，该常数被设置为.015。
通过包含常数，大多数cci值将落入100到-100的范围内。
计算平均偏差有三个步骤。
1.减去最近的20个期间，简单地从该时期的每个典型价格（tp）移动。
2.严格使用绝对值对这些数字进行求和。
3.将步骤3中生成的值除以期间总数
*/

// CCI struct
type CCI struct {
	Period       int     //默认计算几天的
	factor       float64 //计算系数
	points       []cciPoint
	typicalPrice []float64
	avedevPrice  []float64
	maPrice      []float64
	list         []*commonmodels.Kline
}

type cciPoint struct {
	point
}

// NewCCI new Func
func NewCCI(list []*commonmodels.Kline, period int) *CCI {
	m := &CCI{list: list, Period: period, factor: 0.015}
	return m
}

// Calculation Func
func (e *CCI) Calculation() *CCI {

	// 计算TYP
	// TYP:=(HIGH+LOW+CLOSE)/3;
	for i := 0; i < len(e.list); i++ {
		item := e.list[i]
		typicalPrice := (item.High + item.Low + item.Close) / 3.0
		e.typicalPrice = append(e.typicalPrice, typicalPrice)
	}

	// 计算MA
	// MA = MA(TYP,N)
	// var closeArray []float64
	// for _, v := range e.kline {
	// 	closeArray = append(closeArray, v.Close)
	// }
	maStruct := NewMA(e.typicalPrice, e.Period)
	maStruct.Calculation()
	e.maPrice = maStruct.Value

	//计算平均偏差有三个步骤。
	// 1.减去最近的20个期间，简单地从该时期的每个典型价格（tp）移动。
	// 2.严格使用绝对值对这些数字进行求和。
	// 3.将步骤3中生成的值除以期间总数
	for i := 0; i < len(e.maPrice); i++ {
		if i < e.Period-1 {
			e.avedevPrice = append(e.avedevPrice, 0.0)
			continue
		}

		var avedevSum float64
		for j := 0; j < e.Period; j++ {
			avedevSum += math.Abs(e.typicalPrice[i-j] - e.maPrice[i])
		}
		tempAvedevPrice, _ := commonutils.FloatFromString(fmt.Sprintf("%d", e.Period))
		e.avedevPrice = append(e.avedevPrice, avedevSum/tempAvedevPrice)
	}

	//计算 CCI
	// cci =（典型价格 - tp的20周期平均值）/（.015 x平均偏差）
	for i := 0; i < len(e.maPrice); i++ {
		var p cciPoint
		p.Time = e.list[i].KlineTime
		if i < e.Period-1 {
			p.Value = 0
			e.points = append(e.points, p)
			continue
		}

		p.Value = (e.typicalPrice[i] - e.maPrice[i]) / (e.avedevPrice[i] * e.factor)
		e.points = append(e.points, p)
	}
	return e
}

// GetPoints return Point
func (e *CCI) GetPoints() []cciPoint {
	return e.points
}

// GetValue return Value
func (e *CCI) GetValue() []float64 {
	var result []float64
	for _, v := range e.points {
		result = append(result, v.Value)
	}
	return result
}

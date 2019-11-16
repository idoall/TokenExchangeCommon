// Copyright 2016 mshk.top, lion@mshk.top
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commonstock

import (
	"math"
	"time"

	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

/*
日BOLL指标的计算公式
中轨线=N日的移动平均线
上轨线=中轨线+两倍的标准差
下轨线=中轨线－两倍的标准差
日BOLL指标的计算过程
1）计算MA
MA=N日内的收盘价之和÷N
2）计算标准差MD
MD=平方根N日的（C－MA）的两次方之和除以N
3）计算MB、UP、DN线
MB=（N－1）日的MA
UP=MB+2×MD
DN=MB－2×MD
*/

// BOLL struct 布比线
type BOLL struct {
	PeriodN int     //计算周期
	PeriodK float64 //带宽
	points  []bollPoint
	kline   []*commonmodels.Kline
}

type bollPoint struct {
	UP   float64
	MID  float64
	Low  float64
	Time time.Time
}

// NewBOLL Func
// 使用方法，先添加最早日期的数据,最后一条应该是当前日期的数据，结果与 AICoin 对比完全一致
func NewBOLL(list []*commonmodels.Kline) *BOLL {
	return &BOLL{PeriodN: 20, PeriodK: 2.0, kline: list}
}

//sma 计算移动平均线
func (e *BOLL) sma(lines []*commonmodels.Kline) float64 {
	s := len(lines)
	var sum float64 = 0
	for i := 0; i < s; i++ {
		sum += float64(lines[i].Close)
	}
	return sum / float64(s)
}

// dma MD=平方根N日的（C－MA）的两次方之和除以N
func (e *BOLL) dma(lines []*commonmodels.Kline, ma float64) float64 {
	s := len(lines)
	//log.Println(s)
	var sum float64 = 0
	for i := 0; i < s; i++ {
		sum += (lines[i].Close - ma) * (lines[i].Close - ma)
	}
	return math.Sqrt(sum / float64(e.PeriodN))
}

// Calculation Func
func (e *BOLL) Calculation() *BOLL {
	l := len(e.kline)

	e.points = make([]bollPoint, l)
	if l < e.PeriodN {
		for i := 0; i < len(e.kline); i++ {
			e.points[i].Time = e.kline[i].KlineTime
		}
		return e
	}
	for i := l - 1; i > e.PeriodN-1; i-- {

		ps := e.kline[(i - e.PeriodN + 1) : i+1]
		e.points[i].MID = e.sma(ps)

		//MD=平方根N日的（C－MA）的两次方之和除以N
		md := e.dma(ps, e.points[i].MID)
		e.points[i].UP = e.points[i].MID + e.PeriodK*md
		e.points[i].Low = e.points[i].MID - e.PeriodK*md
		e.points[i].Time = e.kline[i].KlineTime
	}
	return e
}

// GetPoints Func
func (e *BOLL) GetPoints() []bollPoint {
	return e.points
}

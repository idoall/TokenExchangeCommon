package commonstock

import (
	"fmt"
	"testing"
)

func TestMACD(t *testing.T) {
	t.Parallel()
	list := InitTestKline()

	//计算新的MACD
	stockList := NewMACD(list).Calculation().GetPoints()

	for _, v := range stockList {
		fmt.Printf("Time:%s\t DIF:%f DEA:%f MACD %f\n", v.Time.Format("2006-01-02 15:04:05"), v.DIF, v.DEA, v.MACD)
	}
}

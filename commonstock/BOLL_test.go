package commonstock

import (
	"fmt"
	"testing"
)

func TestBOLL(t *testing.T) {
	t.Parallel()
	list := InitTestKline()
	//计算新的BOLL
	stockList := NewBOLL(list).Calculation().GetPoints()
	for _, v := range stockList {
		fmt.Printf("Time:%s\t Middle:%.5f Up:%.5f Low:%.5f\n", v.Time.Format("2006-01-02 15:04:05"), v.MID, v.UP, v.Low)
	}

}

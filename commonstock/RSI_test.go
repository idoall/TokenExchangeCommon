package commonstock

import (
	"fmt"
	"testing"
)

func TestRSI(t *testing.T) {
	t.Parallel()
	list := InitTestKline()
	//计算新的OBV
	stock := NewRSI(list, 14)
	stock.Calculation()
	rsiList := stock.GetPoints()

	for i := 0; i < len(rsiList); i++ {
		item := rsiList[i]
		fmt.Printf("[%d]Time:%s RSI:%f\n", i, item.Time.Format("2006-01-02 15:04:05"), item.Value)
	}
}

package commonstock

import (
	"fmt"
	"testing"
)

func TestMA(t *testing.T) {
	t.Parallel()
	list := InitTestKline()

	//计算新的OBV
	stock := NewMA(list, 5)
	ema5List := stock.Calculation().GetPoints()

	for i := 0; i < len(ema5List); i++ {
		e5 := ema5List[i]
		fmt.Printf("[%d][%s] MA5:%.3f \n", i, e5.Time.Format("2006-01-02 15:04:05"), e5.Value)
	}
}

package commonstock

import (
	"fmt"
	"testing"
)

func TestEMA(t *testing.T) {
	t.Parallel()
	list := InitTestKline()
	//计算新的OBV
	stock := NewEMA(list, 5)
	ema5List := stock.Calculation().GetPoints()

	for i := 0; i < len(ema5List); i++ {
		e5 := ema5List[i]
		fmt.Printf("[%d][%s] EMA5:%.3f \n", i, e5.Time.Format("2006-01-02 15:04:05"), e5.Value)
	}
}

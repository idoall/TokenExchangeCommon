package commonstock

import (
	"fmt"
	"testing"
)

func TestMA(t *testing.T) {
	t.Parallel()
	list := InitTestKline()

	var valueArray []float64
	for _, v := range list {
		valueArray = append(valueArray, v.Close)
	}
	//计算新的OBV
	stock := NewMA(valueArray, 5)
	stock.Calculation()

	for i := 0; i < len(stock.Value); i++ {
		// for i := 0; i < 50; i++ {
		fmt.Printf("[%d] Value:%f\n", i, stock.Value[i])
	}
}

package commonstock

import (
	"fmt"
	"testing"
)

func TestCCI(t *testing.T) {
	t.Parallel()
	list := InitTestKline()
	//计算新的OBV
	stock := NewCCI(list, 20)
	stock.Calculation()
	cciList := stock.GetPoints()
	for _, v := range cciList {
		fmt.Printf("Time:%s\t Value:%f\n", v.Time.Format("2006-01-02 15:04:05"), v.Value)
	}

}

package commonmodels

import "time"

// Kline struct
type Kline struct {
	Amount    float64   `orm:"column(amount)"`    // 成交量
	Count     int64     `orm:"column(count)"`     // 成交笔数
	Open      float64   `orm:"column(open)"`      // 开盘价
	Close     float64   `orm:"column(close)"`     // 收盘价, 当K线为最晚的一根时, 时最新成交价
	Low       float64   `orm:"column(low)"`       // 最低价
	High      float64   `orm:"column(high)"`      // 最高价
	Vol       float64   `orm:"column(vol)"`       // 成交额, 即SUM(每一笔成交价 * 该笔的成交数量)
	KlineTime time.Time `orm:"column(klinetime)"` // k线时间
}

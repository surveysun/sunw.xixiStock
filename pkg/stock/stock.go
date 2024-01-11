package stock

import (
	"sunw.xixiStock/api"
	"sunw.xixiStock/pkg/logger"
)

type Stock struct {
}

func NewStock() *Stock {
	return &Stock{}
}

func (s *Stock) GetStockInfo(stockNo string) {
	logger.Log.Debugf("Get StockNo:%v", stockNo)

}

//检查股票行情数据是否满足模型
func (s *Stock) Check(stock api.QTStockInfo) {

}

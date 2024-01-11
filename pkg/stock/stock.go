package stock

import "zhuanqian-xixi-tool/pkg/logger"

type Stock struct {
	log *logger.Logger
}

func NewStock(log *logger.Logger) *Stock {
	return &Stock{
		log: log,
	}
}

func (s *Stock) GetStockInfo(stockNo string) {

}

package api

/*
v_sz002750="51~龙津药业~002750~10.45~10.15~10.20~146272~72304~73968~10.45~80~10.44~634~10.43~249~10.42~572~10.41~127~10.46~423~10.47~201~10.48~339~10.49~165~10.50~477~~20240111161500~0.30~2.96~10.79~10.16~10.45/146272/151549078~146272~15155~3.67~-52.68~~10.79~10.16~6.21~41.67~41.85~7.52~11.17~9.14~0.59~57~10.36~-117.82~-74.58~~~0.44~15154.9078~0.0000~0~
~GP-A~4.50~-3.51~0.00~-14.28~-14.27~13.50~8.80~6.85~-0.29~2.65~398736201~400500000~1.74~4.50~398736201~~~-6.45~0.00~~CNY~0~~10.55~-1072";
*/
//腾讯行情API.
type QTStockInfo struct {
	SC                     string //市场代码.
	Name                   string
	ID                     string  //股票代码
	LastPrice              float32 //当前价格
	PreDayPrice            float32 //昨天收盘价
	StartPrice             float32 //开盘价
	TotalVolume            int32   //成交量.单位手
	OutVolume              int32   //外盘量.单位手: 外盘是主动买
	InVolume               int32
	Buy1Price              float32 //买1价格
	Buy1Count              int32   //买1量
	Buy2Price              float32
	Buy2Count              int32
	Buy3Price              float32
	Buy3Count              int32
	Buy4Price              float32
	Buy4Count              int32
	Buy5Price              float32
	Buy5Count              int32
	Sell1Price             float32 //卖1价格
	Sell1Count             int32   //卖1量
	Sell2Price             float32
	Sell2Count             int32
	Sell3Price             float32
	Sell3Count             int32
	Sell4Price             float32
	Sell4Count             int32
	Sell5Price             float32
	Sell5Count             int32
	DateTime               string  //返回数据的时间点
	StepValue              float32 //涨跌价格
	StepRate               float32 //涨跌比例
	HighPrice              float32 //最高价格
	LowPrice               float32 //最低价格
	LastData               string  //数据格式. 价格/成交量/成交额
	NVolume                int32   //成交量
	NTurnover              int32   //成交额.单位万
	NTurnoverRatio         float32 //换手率
	PERatioTTM             float32 //市盈率.TTM
	HPrice                 float32 //最高价格
	LPrice                 float32 //最低价格
	Amplitude              float32 //振幅
	CirculationMarketValue float32 //流通市值.单位亿
	MarketValue            float32 //市值
	PBRatio                float32 //市净率
	LimitUp                float32 //涨停价格
	LimitDown              float32 //跌停价格
	VolumeRatio            float32 //量比
	Extra1                 float32 //未知
	AveragePrice           float32 //均价
	PERatioDev             float32 //动态市盈率
	PERatioStatic          float32 //静态市盈率
	Extra2                 float32 //未知
	NTurnoverF             float32 //成交额.单位万 精确4位
	//之后的字段不知道是啥.不做解析
}

//K线图
type QTFlashData struct {
}

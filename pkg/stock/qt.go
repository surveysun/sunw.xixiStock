package stock

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sunw.xixiStock/api"
	"sunw.xixiStock/pkg/logger"
)

func GetQTStockInfo(stocks []string) ([]*api.QTStockInfo, error) {
	cnt := len(stocks)
	if cnt == 0 {
		return nil, fmt.Errorf("输出参数stocks为空")
	}
	//一次返回100条最多
	if len(stocks) >= 100 {
		cnt = 100
	}

	url := "http://qt.gtimg.cn/q="
	for i := 0; i < cnt; i++ {
		if i == (cnt - 1) {
			url = url + stocks[i]
		} else {
			url = url + stocks[i] + ","
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		logger.Log.Errorf("http请求GET:%v 错误:%v", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Errorf("获取返回的行情数据错误:%v", err)
		return nil, err
	}

	return Analysis(string(body))
}

/*
原始数据格式
v_sz002449="51~国星光电~002449~8.54~8.24~8.24~130198~76768~53430~8.54~130~8.53~1070~8.52~170~8.51~735~8.50~463~8.55~627~8.56~548~8.57~622~8.58~239~8.59~375~~20240111161439~0.30~3.64~8.55~8.22~8.54/130198/109787961~130198~10979~2.12~66.98~~8.55~8.22~4.00~52.37~52.82~1.39~9.06~7.42~0.70~157~8.43~48.89~43.53~~~1.30~10978.7961~0.0000~0~
~GP-A~-9.82~-5.95~0.70~2.08~1.22~9.97~7.57~1.07~-1.50~-2.18~613182784~618477169~3.15~-4.47~613182784~~~-1.16~-0.12~~CNY~0~~8.49~657";
v_sz002660="51~茂硕电源~002660~10.33~10.07~10.03~230935~131323~99612~10.33~1690~10.32~45~10.31~720~10.30~828~10.29~133~10.34~1523~10.35~887~10.36~605~10.37~270~10.38~796~~20240111161454~0.26~2.58~10.34~10.03~10.33/230935/236232735~230935~23623~8.98~47.26~~10.34~10.03~3.08~26.56~36.84~2.93~11.08~9.06~0.54~-665~10.23~46.30~43.36~~~1.47~23623.2735~0.0000~0~
~GP-A~-16.56~-11.48~0.97~6.19~4.07~14.58~7.79~0.98~6.06~13.02~257084939~356626019~-8.87~7.72~257084939~~~28.80~0.19~~CNY~0~~10.39~-950";
*/

//解析数据
func Analysis(data string) ([]*api.QTStockInfo, error) {
	logger.Log.Debugf("行情数据:%v", data)
	entry := strings.Split(data, ";")
	stocks := make([]*api.QTStockInfo, 0)
	for _, item := range entry {
		str := item[strings.Index(item, "=")+1:]
		str = strings.Replace(str, "\"", "", -1)
		tmp, err := analysis(str)
		if err != nil {
			continue
		}
		stocks = append(stocks, tmp)
	}
	if len(stocks) == 0 {
		logger.Log.Errorf("未解析成功任何一条行情信息.原始信息为:%v", data)
		return nil, fmt.Errorf("未解析成功任何一条行情信息")
	}
	return stocks, nil
}

func analysis(data string) (*api.QTStockInfo, error) {
	return nil, nil
}

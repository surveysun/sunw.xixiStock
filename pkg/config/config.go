package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"sunw.xixiStock/pkg/logger"
)

type ServiceConfig struct {
	DetectionInterval int      `yaml:"detectionInterval"` //间隔时间.单位秒. 默认60秒
	Stocks            []string `yaml:"stocks"`            //股票池.
}

func LoadConfig(path string) (*ServiceConfig, error) {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		logger.Log.Errorf("读取配置文件失败:%v", err)
		return nil, err
	}

	cfg := ServiceConfig{}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		logger.Log.Errorf("解析配置文件失败:%v", err)
		return nil, err
	}
	logger.Log.Infof("配置文件:%+v", cfg)
	return &cfg, nil
}

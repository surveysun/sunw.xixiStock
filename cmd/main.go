package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"sunw.xixiStock/pkg/config"
	"sunw.xixiStock/pkg/logger"
)

func initLogger(path string, debug, isPrint bool) {
	logger.NewLogger(path, isPrint)
	if debug {
		logger.Log.SetLogLevel(log.DebugLevel)
	} else {
		logger.Log.SetLogLevel(log.InfoLevel)
	}
}

var (
	debug      = flag.Bool("debug", false, "是否debug输出,默认false")
	isPrint    = flag.Bool("isPrint", true, "是否控制台输出,默认true")
	logfile    = flag.String("logfile", "./sunw.xixi-stock-service.log", "日志文件.默认./zhuanqian-xixi-service.log")
	configPath = flag.String("configPath", "./config.yaml", "配置文件")
)

func main() {
	flag.Parse()
	initLogger(*logfile, *debug, *isPrint)
	defer logger.Log.Close()
	logger.Log.Info("自动化股票赚钱交易系统 START !!")
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		return
	}
	logger.Log.Infof("End:%v", cfg)
}

package main

import (
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

//zap-go logger日志记录器

var logger *zap.Logger

// var sugarlogger *zap.SugaredLogger

// func initLogger() {
// 	logger, _ = zap.NewProduction()
// 	sugalogger = logger.Sugar()
// }

//Encoder：编码器，如何写入日志，日志格式
func getEncoder() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(cfg)
}

//WriteSyncer：指定日志将写到哪里去
func getLogWriter() zapcore.WriteSyncer {
	// file, _ := os.Create("./test.log")
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,     //单位M
		MaxBackups: 3,     //最大切割数量，不含自身
		MaxAge:     30,    //最大备份天数
		Compress:   false, //是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

//自定义logger配置
func initLogger() {
	encoder := getEncoder()
	writerSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller())
	// sugarlogger = logger.Sugar()
}

func simpleHTTPGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info(
			"Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func main() {
	initLogger()
	defer logger.Sync()

	for i := 0; i < 50000; i++ {
		simpleHTTPGet("www.google.com")
		simpleHTTPGet("http://www.baidu.com")
	}
}

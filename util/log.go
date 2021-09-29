package util

import (
	"github.com/QWERKael/utility-go/log"
	"go.uber.org/zap"
)

var SugarLogger *zap.SugaredLogger

func InitLog() {
	//fmt.Println("初始化模块: log.go")
	var err error
	//SugarLogger, err = log.NewLogger(log.ConsoleEncoder, "", zapcore.DebugLevel)
	SugarLogger, err = log.NewLogger(log.ConsoleEncoder, "", log.ConvertLogLevel("debug"))
	if err != nil {
		panic(err.Error())
	}
	SugarLogger.Debug("日志记录开始...")
}

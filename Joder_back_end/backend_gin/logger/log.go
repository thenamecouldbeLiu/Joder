package logger

import (
	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

func getLevel(lvl string) log.Level {
	switch lvl {

	case "info":
		return log.InfoLevel

	case "Error":
		{
			return log.ErrorLevel
		}
	default:
		return log.DebugLevel
	}

}

func NewLogger() *log.Logger {
	var level = getLevel("")
	logger := &log.Logger{
		//Out:   os.Stdout,
		Level: level,
	}
	Logger = logger
	return Logger
}

func Init() {
	//log輸出為json格式
	//log.SetFormatter(&log.JSONFormatter{})
	//輸出設定為標準輸出(預設為stderr)
	//log.SetOutput(os.Stdout)
	//設定要輸出的log等級
	log.SetLevel(log.DebugLevel)
}

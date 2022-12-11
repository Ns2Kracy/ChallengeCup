package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	// 将日志输出到文件
	file, err := os.OpenFile("../log/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	defer file.Close()

	// 设置日志级别
	log.SetLevel(log.DebugLevel)

	// 设置日志格式
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	// 设置日志输出位置
	log.SetReportCaller(true)
}

func Debug(message string, args ...interface{}) {
	log.Debugf(message, args...)
}

func Info(message string, args ...interface{}) {
	log.Infof(message, args...)
}

func Warn(message string, args ...interface{}) {
	log.Warnf(message, args...)
}

func Error(message string, args ...interface{}) {
	log.Errorf(message, args...)
}

func Fatal(message string, args ...interface{}) {
	log.Fatalf(message, args...)
}

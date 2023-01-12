package main

import (
	"net"
	"os"
	"time"

	"ChallengeCup/config"
	"ChallengeCup/router"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	router.InitRoute(app)
	conf, err := config.NewConfig("config.yaml")
	if err != nil {
		return
	}
	// logFile := NewLogFile()
	// app.Logger().SetOutput(logFile)
	listener, err := net.Listen("tcp", conf.System.Host+":"+conf.System.Port)
	if err != nil {
		return
	}

	defer listener.Close()

	if err := app.Run(
		iris.Listener(listener),
		iris.WithOptimizations,
		iris.WithoutBanner,
		iris.WithConfiguration(iris.Configuration{
			Charset:  "UTF-8",
			LogLevel: "DEBUG",
		}),
		iris.WithTimeFormat("2006-01-02 15:04:05"),
	); err != nil {
		return
	}
}

func NewLogFile() *os.File {
	logFile := time.Now().Format("2006-01-02") + ".log"
	file, err := os.OpenFile("./log/"+logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return nil
	}
	return file
}

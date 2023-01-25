package main

import (
	"net"
	"os"
	"time"

	"ChallengeCup/config"
	"ChallengeCup/router"
	"ChallengeCup/utils/file"

	"github.com/kataras/iris/v12"
)

func main() {
	conf := config.InitConfig("config.yaml")
	app := iris.New()
	router.InitRoute(app)

	logFile := NewLogFile()
	app.Logger().AddOutput(logFile)
	defer logFile.Close()

	listener, err := net.Listen("tcp", conf.System.Host+":"+conf.System.Port)
	if err != nil {
		return
	}

	defer listener.Close()

	if err := app.Run(
		iris.Listener(listener),
		iris.WithOptimizations,
		iris.WithoutInterruptHandler,
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
	logdir := "./log"
	if !file.IsExist(logdir) {
		err := file.NewDir(logdir)
		if err != nil {
			panic(err)
		}
	}
	file, err := os.OpenFile("./log/"+logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		panic(err)
	}
	return file
}

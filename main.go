package main

import (
	"net"
	"time"

	"ChallengeCup/config"
	"ChallengeCup/dao"
	"ChallengeCup/router"

	"ChallengeCup/utils/mqtt"

	"github.com/kataras/iris/v12"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func main() {
	conf := config.InitConfig("config.yaml")
	app := iris.New()
	router.InitRoute(app)

	logPath := "./log/"
	logWriter, _ := rotatelogs.New(
		logPath+"%Y-%m-%d"+".log",
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)
	app.Logger().AddOutput(logWriter)

	mqtt.InitMqtt()
	dao.NewRedis()

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

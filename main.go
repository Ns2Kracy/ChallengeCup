package main

import (
	"net"

	"ChallengeCup/dao"
	"ChallengeCup/service"

	"ChallengeCup/config"
	"ChallengeCup/router"

	"github.com/kataras/iris/v12"
)

var Configuration = iris.Configuration{
	Charset: "UTF-8",
}

func main() {
	app := iris.New()
	router.InitRoute(app)
	_ = service.NewService(dao.InitMysql())
	conf, err := config.NewConfig("config.yaml")
	if err != nil {
		return
	}

	listener, err := net.Listen("tcp", conf.System.Host+":"+conf.System.Port)
	if err != nil {
		return
	}

	defer listener.Close()

	if err := app.Run(
		iris.Listener(listener),
		iris.WithLogLevel("DEBUG"),
		iris.WithOptimizations,
		iris.WithConfiguration(Configuration),
		iris.WithTimeFormat("2006-01-02 15:04:05"),
	); err != nil {
		return
	}
}

package main

import (
	"net"

	"ChallengeCup/config"
	router "ChallengeCup/router"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	router.InitRoute(app)

	config := config.InitConfig()

	listener, err := net.Listen("tcp", config.System.Host+":"+config.System.Port)
	if err != nil {
		return
	}

	defer listener.Close()

	if err := app.Run(
		iris.Listener(listener),
		iris.WithLogLevel("DEBUG"),
		iris.WithOptimizations,
		iris.WithConfiguration(iris.Configuration{
			// DisableStartupLog: true,
			Charset: "UTF-8",
		}),
		iris.WithTimeFormat("2006-01-02 15:04:05"),
	); err != nil {
		return
	}
}

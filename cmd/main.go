package main

import (
	"fmt"
	"os"
	"os/signal"

	server "github.com/DarkSoul94/simple-websocket/cmd/httpserver"
	"github.com/DarkSoul94/simple-websocket/config"
	"github.com/DarkSoul94/simple-websocket/pkg/logger"
)

func main() {
	conf := config.InitConfig()
	logger.InitLogger(conf)

	apphttp := server.NewApp()
	go apphttp.Run(conf)

	fmt.Println(
		fmt.Sprintf(
			"Service %s is running",
			conf.AppName,
		),
	)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	apphttp.Stop()

	fmt.Println(
		fmt.Sprintf(
			"Service %s is stopped",
			conf.AppName,
		),
	)
}

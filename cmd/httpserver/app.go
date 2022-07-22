package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/DarkSoul94/simple-websocket/app"
	apphttp "github.com/DarkSoul94/simple-websocket/app/delivery/http"
	appusecase "github.com/DarkSoul94/simple-websocket/app/usecase"
	"github.com/DarkSoul94/simple-websocket/config"
)

// App ...
type App struct {
	uc         app.IUsecase
	httpServer *http.Server
}

// NewApp ...
func NewApp() *App {
	uc := appusecase.NewUsecase()

	return &App{
		uc: uc,
	}
}

// Run run application
func (a *App) Run(conf config.Config) {
	mux := http.NewServeMux()

	apphttp.RegisterHTTPEndpoints(mux, a.uc)

	a.httpServer = &http.Server{
		Addr:           ":" + conf.HTTPport,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	var l net.Listener
	var err error
	l, err = net.Listen("tcp", a.httpServer.Addr)
	if err != nil {
		panic(err)
	}

	if err := a.httpServer.Serve(l); err != nil {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}

func (a *App) Stop() error {
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

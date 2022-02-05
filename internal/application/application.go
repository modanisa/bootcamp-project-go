package application

import (
	"github.com/modanisa/bootcamp-project-go/internal/hello"
	"github.com/modanisa/bootcamp-project-go/pkg/server"
	"os"
	"os/signal"
	"syscall"
)

type Application struct {
	server server.Server
}

func New() (*Application, error) {
	helloHandler := hello.NewHandler()

	s := server.New([]server.RegisterRoutesFunc{
		helloHandler.RegisterRoutes,
	})

	return &Application{server: s}, nil
}

func (a *Application) Run() error {
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdownChan
		_ = a.server.Close()
	}()

	return a.server.Run()
}

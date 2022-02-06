package application

import (
	"github.com/modanisa/bootcamp-project-go/internal/hello"
	"github.com/modanisa/bootcamp-project-go/internal/product"
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

	productRepository := product.NewRepository()
	productService := product.NewService(productRepository)
	productHandler := product.NewHandler(productService)

	s := server.New([]server.RegisterRoutesFunc{
		helloHandler.RegisterRoutes,
		productHandler.RegisterRoutes,
	})

	return &Application{server: s}, nil
}

func (a *Application) Run() error {
	shutdownChan := make(chan os.Signal, 1)
	// "Ctrl+C" sends `SIGINT`
	// "kill" commands SEND `SIGTERM`
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdownChan
		_ = a.server.Close()
	}()

	return a.server.Run()
}

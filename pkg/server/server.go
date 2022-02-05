package server

import "github.com/gofiber/fiber/v2"

type Server struct {
	app *fiber.App
}

type RegisterRoutesFunc func(app *fiber.App)

func New(registerHandlerFuncs []RegisterRoutesFunc) Server {
	app := fiber.New(
		fiber.Config{
			DisableStartupMessage: true,
		},
	)

	server := Server{app: app}

	for _, registerHandlersFunc := range registerHandlerFuncs {
		registerHandlersFunc(server.app)
	}

	return server
}

func (s Server) Run() error {
	return s.app.Listen(":8080")
}

func (s Server) Close() error {
	return s.app.Shutdown()
}

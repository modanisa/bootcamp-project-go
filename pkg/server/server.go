package server

import "github.com/gofiber/fiber/v2"

type Server struct {
	app *fiber.App
}

type RegisterRoutesFunc func(app *fiber.App)

func New(registerHandlerFuncs []RegisterRoutesFunc) {
	/*app := fiber.New(
		fiber.Config{
			DisableStartupMessage: true
		},
	)*/
}

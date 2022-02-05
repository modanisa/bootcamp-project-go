package main

import (
	"github.com/modanisa/bootcamp-project-go/internal/application"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	app, err := application.New()
	if err != nil {
		return err
	}

	return app.Run()
}

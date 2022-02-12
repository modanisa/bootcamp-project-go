package main

import (
	"bootcamp/server"
	"log"
)

func main() {
	svr := server.NewServer()
	err := svr.StartServer(3000)
	if err != nil {
		log.Fatalln(err)
	}
}

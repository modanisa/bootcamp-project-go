package main

import (
	"fmt"
	"net/http"

	"bootcamp/client"
	"bootcamp/config"
	"bootcamp/handler"
	"bootcamp/service"
)

func main() {
	config := config.Get()

	client := client.NewClient(config.ServiceURL)
	service := service.NewService(client)
	h := handler.NewHander(service)

	http.HandleFunc("/", h.Quotes)
	err := http.ListenAndServe(config.ServerAddr, nil)
	if err != nil {
		fmt.Println(err)
	}
}

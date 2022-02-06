package main

import (
	"bootcamp/client"
	"bootcamp/config"
	"bootcamp/handler"
	"bootcamp/service"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(config.C.ServiceURL)

	client := client.NewClient(config.C.ServiceURL)

	service := service.NewService(client)
	fmt.Println(service.Quotes())

	h := handler.NewHander(service)

	http.HandleFunc("/", h.Quotes)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

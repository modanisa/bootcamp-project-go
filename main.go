package main

import (
	"bootcamp/handler"
	"bootcamp/repository"
	"bootcamp/service"
	"fmt"
	"net/http"
)

func main() {
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	http.HandleFunc("/api/v1/products", productHandler.GetProducts)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

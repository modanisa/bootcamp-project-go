package server

import (
	"bootcamp/handler"
	"bootcamp/repository"
	"bootcamp/service"
	"fmt"
	"net/http"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer(port int) error {
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	http.HandleFunc("/api/v1/products", productHandler.GetProducts)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return err
}

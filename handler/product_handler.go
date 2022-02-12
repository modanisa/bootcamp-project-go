package handler

import (
	"bootcamp/service"
	"encoding/json"
	"net/http"
)

type IProductHandler interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
}

type ProductHandler struct {
	psvc service.IProductService
}

func NewProductHandler(psvc service.IProductService) IProductHandler {
	return &ProductHandler{
		psvc: psvc,
	}
}

func (p *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	products, err := p.psvc.Products()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	productsBytes, _ := json.Marshal(products)

	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(productsBytes)
}
package handler

import (
	"bootcamp/service"
	"net/http"
)

type IHandler interface {
	Quotes(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.IQuotesService
}

func (c *Handler) Quotes(w http.ResponseWriter, r *http.Request) {
	c.service.Quotes()

	w.Write([]byte(""))
}

func NewHander(service service.IQuotesService) IHandler {
	return &Handler{service: service}
}

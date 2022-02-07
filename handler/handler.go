package handler

import (
	"encoding/json"
	"net/http"

	"bootcamp/service"
)

type IHandler interface {
	Quotes(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.IQuotesService
}

func (c *Handler) Quotes(w http.ResponseWriter, r *http.Request) {
	response, err := c.service.Quotes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(json)
}

func NewHander(service service.IQuotesService) IHandler {
	return &Handler{service: service}
}

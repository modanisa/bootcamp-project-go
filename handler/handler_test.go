package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"bootcamp/mock"
	"bootcamp/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandlerServiceReturnQuotes(t *testing.T) {
	service := mock.NewMockIQuotesService(gomock.NewController(t))
	serviceReturn := &model.QuotesResponse{
		"": []string{"test"},
	}
	service.EXPECT().
		Quotes().
		Return(serviceReturn, nil).
		Times(1)

	handler := NewHander(service)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler.Quotes(w, r)

	actual := &model.QuotesResponse{}
	json.Unmarshal(w.Body.Bytes(), actual)

	assert.Equal(t, serviceReturn, actual)
	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

func TestHandlerServiceReturnError(t *testing.T) {
	service := mock.NewMockIQuotesService(gomock.NewController(t))
	serviceError := errors.New("test error")
	service.EXPECT().
		Quotes().
		Return(nil, serviceError).
		Times(1)

	handler := NewHander(service)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler.Quotes(w, r)

	assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
	assert.Equal(t, w.Body.String(), serviceError.Error())
}

func TestHandlerServiceReturnNilQuotes(t *testing.T) {
	service := mock.NewMockIQuotesService(gomock.NewController(t))
	service.EXPECT().
		Quotes().
		Return(nil, nil).
		Times(1)

	handler := NewHander(service)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler.Quotes(w, r)

	assert.Equal(t, w.Body.String(), "null")
	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

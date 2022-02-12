package handler_test

import (
	"bootcamp/handler"
	"bootcamp/mock"
	"bootcamp/model"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProducts(t *testing.T) {
	t.Run("only GET method allowed", func(t *testing.T) {
		productHandler := handler.NewProductHandler(nil)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/products", http.NoBody)
		res := httptest.NewRecorder()

		productHandler.GetProducts(res, req)

		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
	t.Run("POST not allowed", func(t *testing.T) {
		productHandler := handler.NewProductHandler(nil)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", http.NoBody)
		res := httptest.NewRecorder()

		productHandler.GetProducts(res, req)
		assert.Equal(t, http.StatusNotImplemented, res.Result().StatusCode)
	})
	t.Run("Should return all products correctly", func(t *testing.T) {
		mockService := mock.NewMockIProductService(gomock.NewController(t))

		mockService.
			EXPECT().
			Products().
			Return(model.ProductsResponse{
				model.ProductResponse{ID: 100},
			}, nil).
			Times(1)

		productHandler := handler.NewProductHandler(mockService)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/products", http.NoBody)
		res := httptest.NewRecorder()

		productHandler.GetProducts(res, req)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
		assert.Equal(t, "application/json", res.Result().Header.Get("content-type"))

		expectedResBody := model.ProductsResponse{}
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)
		assert.Nil(t, err, "json unmarshal'da err oldu")

		assert.Equal(t, expectedResBody[0].ID, 100)
	})
	t.Run("should return server internal error when product service failed", func(t *testing.T) {
		mockService := mock.NewMockIProductService(gomock.NewController(t))

		mockService.
			EXPECT().
			Products().
			Return(model.ProductsResponse{}, errors.New("error occured")).
			Times(1)

		productHandler := handler.NewProductHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/products", http.NoBody)
		res := httptest.NewRecorder()
		productHandler.GetProducts(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Result().StatusCode)
	})
}

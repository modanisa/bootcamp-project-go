package service_test

import (
	"bootcamp/mock"
	"bootcamp/model"
	"bootcamp/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldDelegateRepositoryCorrectly(t *testing.T) {
	mockController := gomock.NewController(t)
	mockProductRepository := mock.NewMockIProductRepository(mockController)

	mockProductRepository.
		EXPECT().
		FindAllProducts().
		Return(model.ProductsResponse{}, nil).
		Times(1)

	productService := service.NewProductService(mockProductRepository)
	products, err := productService.Products()

	assert.Equal(t, model.ProductsResponse{}, products)
	assert.Nil(t, err)
}

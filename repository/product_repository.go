package repository

import "bootcamp/model"

type IProductRepository interface {
	FindAllProducts() (model.ProductsResponse, error)
}

type ProductRepository struct {
}

func NewProductRepository() IProductRepository {
	return &ProductRepository{}
}

func (p *ProductRepository) FindAllProducts() (model.ProductsResponse, error) {
	products := model.ProductsResponse{}
	products = append(products, model.ProductResponse{
		ID:          1,
		Name:        "tavin",
		Slug:        "tavin",
		Description: "tavin marka kıyafet",
		Image:       "https://fns.modanisa.com/r/pro2/2021/11/01/n-desenli-mevlana-elbise-tsd4414-turuncu-8067476-7.jpg",
	})

	return products, nil
}

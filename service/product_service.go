package service

import (
	"bootcamp/model"
	"bootcamp/repository"
)

type IProductService interface {
	Products() (model.ProductsResponse, error)
}

type ProductService struct {
	repo repository.IProductRepository
}

func NewProductService(repo repository.IProductRepository) IProductService {
	return &ProductService{
		repo: repo,
	}
}

func (p *ProductService) Products() (model.ProductsResponse, error) {
	return p.repo.FindAllProducts()
}

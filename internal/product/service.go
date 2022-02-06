package product

import "context"

type Repository interface {
	GetProducts(ctx context.Context) ([]Product, error)
}

type DefaultService struct {
	repo Repository
}

func NewService(repository Repository) *DefaultService {
	return &DefaultService{repo: repository}
}

func (s *DefaultService) GetProducts(ctx context.Context) ([]Product, error) {
	return s.repo.GetProducts(ctx)
}

package product

import "context"

type Repository interface {
	GetProducts(ctx context.Context) ([]Product, error)
}

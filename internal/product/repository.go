package product

import "context"

type Repository interface {
	GetProducts(ctx context.Context) ([]Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetProducts(ctx context.Context) ([]Product, error) {
	return []Product{
		{
			ID:          1,
			Name:        "Tesettür Dünyası",
			Description: "Desenli Mevlana Elbise TSD4414 Turuncu",
			Image:       "https://fns.modanisa.com/r/pro2/2021/11/01/n-desenli-mevlana-elbise-tsd4414-turuncu-8067476-7.jpg",
		},
	}, nil
}

package product

import "context"

type inMemoryRepository struct{}

func NewInMemoryRepository() Repository {
	return &inMemoryRepository{}
}

func (r *inMemoryRepository) GetProducts(ctx context.Context) ([]Product, error) {
	return []Product{
		{
			ID:          1,
			Name:        "Tesettür Dünyası",
			Slug:        "tesettur-dunyasi",
			Description: "Desenli Mevlana Elbise TSD4414 Turuncu",
			Image:       "https://fns.modanisa.com/r/pro2/2021/11/01/n-desenli-mevlana-elbise-tsd4414-turuncu-8067476-7.jpg",
		},
	}, nil
}

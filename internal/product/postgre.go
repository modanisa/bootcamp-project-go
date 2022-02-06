package product

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

type postgreRepository struct {
	db *gorm.DB
}

func NewPostgreRepository() Repository {
	const dsn = "host=localhost user=modanisa password=mdns-pw dbname=mdns-bootcamp port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.Fatalf("Error connecting postgresql db %v", err)
	}

	db.AutoMigrate(&Product{})

	db.Model(&Product{}).Create(&Product{
		ID:          1,
		Name:        "Tesettür Dünyası",
		Description: "Desenli Mevlana Elbise TSD4414 Turuncu",
		Image:       "https://fns.modanisa.com/r/pro2/2021/11/01/n-desenli-mevlana-elbise-tsd4414-turuncu-8067476-7.jpg",
	})

	return &postgreRepository{db: db}
}

func (p *postgreRepository) GetProducts(ctx context.Context) ([]Product, error) {
	var products []Product
	err := p.db.Model(&Product{}).Find(&products)

	return products, err.Error
}

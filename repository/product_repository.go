package repository

import (
	"store-dashboard-service/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (p *productRepository) Create(product *model.Product) error {
	return p.db.Create(product).Error
}

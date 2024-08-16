package model

import (
	"mime/multipart"
	"time"

	"github.com/lib/pq"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       int64
	CategoryId  int
	Stock       int
	Images      pq.StringArray `gorm:"type:text[]"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateProductRequest struct {
	Name        string                  `json:"name" validate:"required"`
	Description string                  `json:"description" validate:"required"`
	Price       int64                   `json:"price" validate:"required"`
	CategoryId  int                     `json:"categoryId" validate:"required"`
	Stock       int                     `json:"stock" validate:"required"`
	Images      []*multipart.FileHeader `json:"images" validate:"required"`
}

type ProductResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	Category    Category  `json:"category"`
	Stock       int       `json:"stock"`
	Images      []string  `json:"images"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

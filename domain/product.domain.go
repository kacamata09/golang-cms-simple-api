package domain

import (
	// "database/sql"
	// "time"
	// "github.com/labstack/echo"
)

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Price int32 `json:"price"`
	Photo string `json:"photo"`
	Description string `json:"description"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(id string) (Product, error)
	Create(product *Product) error
	// Update(ar *Article) error
	// Delete(id string) error
}

type ProductUsecase interface {
	GetAllData() ([]Product, error)
	GetByID(id string) (Product, error)
	Create(product *Product) error
}
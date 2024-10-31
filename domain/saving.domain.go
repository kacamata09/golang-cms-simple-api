package domain

import (
	// "database/sql"
	// "time"
	// "github.com/labstack/echo"
)

type Saving struct {
	ID string `json:"id"`
	UserID string `json:"user_id"`
	Saldo int32 `json:"saldo"`
	SaldoIn int32 `json:"saldo_in"`
	SaldoOut int32 `json:"saldo_out"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type SavingRepository interface {
	GetAll() ([]Saving, error)
	GetByID(id string) (Saving, error)
	Create(saving *Saving) error
	// Update(ar *Article) error
	// Delete(id string) error
}

type SavingUsecase interface {
	GetAllData() ([]Saving, error)
	GetByID(id string) (Saving, error)
	Create(saving *Saving) error
}
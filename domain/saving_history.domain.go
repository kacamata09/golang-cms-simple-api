package domain

import (
	// "database/sql"
	// "time"
	// "github.com/labstack/echo"
)



type SavingHistory struct {
	ID string `json:"id"`
	UserID string `json:"user_id"`
	Amount int32 `json:"amount"`
	Type string `json:"type"`
	Detail string `json:"detail"`
	InputedBy string `json:"inputed_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type SavingHistoryRepository interface {
	GetAll() ([]SavingHistory, error)
	GetByID(id string) (SavingHistory, error)
	Create(saving_history *SavingHistory) error
	// Update(ar *Article) error
	// Delete(id string) error
}

type SavingHistoryUsecase interface {
	GetAllData() ([]SavingHistory, error)
	GetByID(id string) (SavingHistory, error)
	Create(saving_history *SavingHistory) error
}
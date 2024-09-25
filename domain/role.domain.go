package domain

import (
	// "database/sql"
	"time"
	// "github.com/labstack/echo"
)



type Role struct {
	// ID string `json:"id"`
	ID string `json:"id"`
	RoleName string `json:"role_name"`
	RoleDesc string `json:"role_desc"`
	AccessPermission string `json:"access_perm"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoleRepository interface {
	GetAll() ([]Role, error)
	GetByID(id string) (Role, error)
	Create(role *Role) error
	Update(id string, role *Role) (affect int64, err error)
	Delete(id string) (affect int64, err error)
}

type RoleUsecase interface {
	GetAllData() ([]Role, error)
	GetByID(id string) (Role, error)
	Create(role *Role) error
	Update(id string, role *Role) (affect int64, err error)
	Delete(id string) (affect int64, err error)
}
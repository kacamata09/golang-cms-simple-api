package usecase

import (
	"database/sql"
	"go-clean-architecture-by-ahr/domain"
	// "time"
	// "github.com/labstack/echo"
)

type RoleUsecase struct {
	RoleRepo domain.RoleRepository
	DB *sql.DB

}

func CreateRoleUseCase(repo domain.RoleRepository) domain.RoleUsecase {
	usecase := RoleUsecase {
		RoleRepo: repo,
	}

	return &usecase
}

func (uc RoleUsecase) GetAllData() ([]domain.Role, error) {
	data, err := uc.RoleRepo.GetAll()
	return data, err
}

func (uc RoleUsecase) GetByID(id string) (domain.Role, error) {
	data, err := uc.RoleRepo.GetByID(id)
	return data, err
}

func (uc RoleUsecase) Create(input *domain.Role)  error {
	// roleExisted := uc.RoleRepo.GetByName(input.RoleName)
	// if roleExisted {
	// 	return "sudah ada coy"
	// }
	
	err := uc.RoleRepo.Create(input)
	return err
}

func (uc RoleUsecase) Update(id string, input *domain.Role)  (affect int64, err error) {
	// roleExisted := uc.RoleRepo.GetByName(input.RoleName)
	// if roleExisted {
	// 	return "sudah ada coy"
	// }
	
	affect, err = uc.RoleRepo.Update(id, input)
	return 
}

func (uc RoleUsecase) Delete(id string)  (affect int64, err error) {
	affect, err = uc.RoleRepo.Delete(id)
	return 
}


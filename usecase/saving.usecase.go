package usecase

import (
	"database/sql"
	"go-clean-architecture-by-ahr/domain"
	// "time"
	// "github.com/labstack/echo"
)

type SavingUsecase struct {
	SavingRepo domain.SavingRepository
	DB *sql.DB
}

func CreateSavingUseCase(repo domain.SavingRepository) domain.SavingUsecase {
	usecase := SavingUsecase {
		SavingRepo: repo,
	}

	return &usecase
}

func (uc SavingUsecase) GetAllData() ([]domain.Saving, error) {
	data, err := uc.SavingRepo.GetAll()
	return data, err
}

func (uc SavingUsecase) GetByID(id string) (domain.Saving, error) {
	data, err := uc.SavingRepo.GetByID(id)
	return data, err
}

func (uc SavingUsecase) Create(input *domain.Saving)  error {
	// usernameExisted, _ := uc.SavingRepo.GetByUsername(input.Username)
	// if usernameExisted {
	// 	return "sudah ada coy"
	// }

	// emailExisted, _ := uc.SavingRepo.GetByEmail(input.Email)
	// if emailExisted {
	// 	return "sudah ada coy"
	// }
	
	err := uc.SavingRepo.Create(input)
	return err
}


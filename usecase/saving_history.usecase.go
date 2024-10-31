package usecase

import (
	"database/sql"
	"go-clean-architecture-by-ahr/domain"
	// "time"
	// "github.com/labstack/echo"
)

type SavingHistoryUsecase struct {
	SavingHistoryRepo domain.SavingHistoryRepository
	DB *sql.DB
}

func CreateSavingHistoryUseCase(repo domain.SavingHistoryRepository) domain.SavingHistoryUsecase {
	usecase := SavingHistoryUsecase {
		SavingHistoryRepo: repo,
	}

	return &usecase
}

func (uc SavingHistoryUsecase) GetAllData() ([]domain.SavingHistory, error) {
	data, err := uc.SavingHistoryRepo.GetAll()
	return data, err
}

func (uc SavingHistoryUsecase) GetByID(id string) (domain.SavingHistory, error) {
	data, err := uc.SavingHistoryRepo.GetByID(id)
	return data, err
}

func (uc SavingHistoryUsecase) Create(input *domain.SavingHistory)  error {
	// usernameExisted, _ := uc.SavingHistoryRepo.GetByUsername(input.Username)
	// if usernameExisted {
	// 	return "sudah ada coy"
	// }

	// emailExisted, _ := uc.SavingHistoryRepo.GetByEmail(input.Email)
	// if emailExisted {
	// 	return "sudah ada coy"
	// }
	
	err := uc.SavingHistoryRepo.Create(input)
	return err
}


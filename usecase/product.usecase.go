package usecase

import (
	"database/sql"
	"go-clean-architecture-by-ahr/domain"
	// "time"
	// "github.com/labstack/echo"
)

type ProductUsecase struct {
	ProductRepo domain.ProductRepository
	DB *sql.DB
}

func CreateProductUseCase(repo domain.ProductRepository) domain.ProductUsecase {
	usecase := ProductUsecase {
		ProductRepo: repo,
	}

	return &usecase
}

func (uc ProductUsecase) GetAllData() ([]domain.Product, error) {
	data, err := uc.ProductRepo.GetAll()
	return data, err
}

func (uc ProductUsecase) GetByID(id string) (domain.Product, error) {
	data, err := uc.ProductRepo.GetByID(id)
	return data, err
}

func (uc ProductUsecase) Create(input *domain.Product)  error {
	// usernameExisted, _ := uc.ProductRepo.GetByUsername(input.Username)
	// if usernameExisted {
	// 	return "sudah ada coy"
	// }

	// emailExisted, _ := uc.ProductRepo.GetByEmail(input.Email)
	// if emailExisted {
	// 	return "sudah ada coy"
	// }
	
	err := uc.ProductRepo.Create(input)
	return err
}


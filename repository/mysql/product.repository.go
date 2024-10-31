package repositoryMySql

import (
	"database/sql"
	"fmt"

	// "time"
	"go-clean-architecture-by-ahr/domain"
)

type repoProduct struct{
    DB *sql.DB
}

func CreateRepoProduct(db *sql.DB) domain.ProductRepository {
    return &repoProduct{DB: db}
}

func (repo *repoProduct) GetAll() ([]domain.Product, error) {

	rows, err := repo.DB.Query("SELECT * FROM products")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var data []domain.Product
    
    for rows.Next() {
        var product domain.Product
        err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Photo, &product.Description, 
             &product.CreatedAt, &product.UpdatedAt)
        fmt.Println(err)
        if err != nil {
            return data, err
        }
        data = append(data, product)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }
    return data, err
}

func (repo *repoProduct) GetByID(id string) (domain.Product, error) {
    row := repo.DB.QueryRow("SELECT * FROM products where id=?", id)
    fmt.Println(id)
    
    var data domain.Product
    
    err := row.Scan(&data.ID, &data.Name, &data.Price, &data.Photo, &data.Description,
         &data.CreatedAt, &data.UpdatedAt)
    if err != nil {
        return data, err
    }
    
    if err := row.Err(); err != nil {
        return domain.Product{}, err
    }
    // fmt.Println(data)
    return data, err
}

func (repo *repoProduct) Create(product *domain.Product) error {
    _, err := repo.DB.Exec("INSERT INTO products (name, price, photo, description) values (?, ?, ?, ?)", 
     product.Name, product.Price, product.Photo, product.Description)
    return err
}
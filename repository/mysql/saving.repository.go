package repositoryMySql

import (
	"database/sql"
	"fmt"

	// "time"
	"go-clean-architecture-by-ahr/domain"
)

type repoSaving struct{
    DB *sql.DB
}

func CreateRepoSaving(db *sql.DB) domain.SavingRepository {
    return &repoSaving{DB: db}
}

func (repo *repoSaving) GetAll() ([]domain.Saving, error) {

	rows, err := repo.DB.Query("SELECT * FROM savings")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var data []domain.Saving
    
    for rows.Next() {
        var saving domain.Saving
        err := rows.Scan(&saving.ID, &saving.UserID, &saving.Saldo, &saving.SaldoIn, &saving.SaldoOut, 
             &saving.CreatedAt, &saving.UpdatedAt)
        fmt.Println(err)
        if err != nil {
            return data, err
        }
        data = append(data, saving)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }
    return data, err
}

func (repo *repoSaving) GetByID(id string) (domain.Saving, error) {
    row := repo.DB.QueryRow("SELECT * FROM savings where id=?", id)
    fmt.Println(id)
    
    var data domain.Saving
    
    err := row.Scan(&data.ID, &data.UserID, &data.Saldo, &data.SaldoIn, &data.SaldoOut, 
         &data.CreatedAt, &data.UpdatedAt)
    if err != nil {
        return data, err
    }
    
    if err := row.Err(); err != nil {
        return domain.Saving{}, err
    }
    // fmt.Println(data)
    return data, err
}

func (repo *repoSaving) Create(saving *domain.Saving) error {
    _, err := repo.DB.Exec("INSERT INTO savings (user_id, saldo, saldo_in, saldo_out) values (?, ?, ?, ?)", 
    saving.UserID, saving.Saldo, saving.SaldoIn, saving.SaldoOut)
    return err
}
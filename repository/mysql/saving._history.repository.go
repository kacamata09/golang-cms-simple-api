package repositoryMySql

import (
	"database/sql"
	"fmt"

	// "time"
	"go-clean-architecture-by-ahr/domain"
)

type repoSavingHistory struct{
    DB *sql.DB
}

func CreateRepoHistorySaving(db *sql.DB) domain.SavingHistoryRepository {
    return &repoSavingHistory{DB: db}
}

func (repo *repoSavingHistory) GetAll() ([]domain.SavingHistory, error) {

	rows, err := repo.DB.Query("SELECT * FROM saving_histories")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var data []domain.SavingHistory
    
    for rows.Next() {
        var saving_history domain.SavingHistory
        err := rows.Scan(&saving_history.ID, &saving_history.UserID, &saving_history.Amount, &saving_history.Type, &saving_history.Detail, &saving_history.InputedBy, 
             &saving_history.CreatedAt, &saving_history.UpdatedAt)
        fmt.Println(err)
        if err != nil {
            return data, err
        }
        data = append(data, saving_history)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }
    return data, err
}

func (repo *repoSavingHistory) GetByID(id string) (domain.SavingHistory, error) {
    row := repo.DB.QueryRow("SELECT * FROM saving_histories where id=?", id)
    fmt.Println(id)
    
    var data domain.SavingHistory
    
    err := row.Scan(&data.ID, &data.UserID, &data.Amount, &data.Type, &data.Detail, &data.InputedBy, 
         &data.CreatedAt, &data.UpdatedAt)
    if err != nil {
        return data, err
    }
    
    if err := row.Err(); err != nil {
        return domain.SavingHistory{}, err
    }
    // fmt.Println(data)
    return data, err
}

func (repo *repoSavingHistory) Create(saving_history *domain.SavingHistory) error {
    _, err := repo.DB.Exec("INSERT INTO saving_histories (user_id, amount, type, detail, inputed_by) values (?, ?, ?, ?, ?)", 
    saving_history.UserID, saving_history.Amount, saving_history.Type, saving_history.Detail, saving_history.InputedBy, )
    return err
}
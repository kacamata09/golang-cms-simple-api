package repositoryPgSQL

import (
	"database/sql"
	"fmt"

	// "time"
	"go-clean-architecture-by-ahr/domain"
)

type repoRole struct{
    DB *sql.DB
}

func CreateRepoRole(db *sql.DB) domain.RoleRepository {
    return &repoRole{DB: db}
}

func (repo *repoRole) GetAll() ([]domain.Role, error) {

	rows, err := repo.DB.Query("SELECT * FROM roles")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var data []domain.Role
    
    for rows.Next() {
        var role domain.Role
        err := rows.Scan(&role.ID, &role.RoleName, &role.RoleDesc, &role.AccessPermission, &role.CreatedAt, &role.UpdatedAt)
        if err != nil {
            panic("Failed to scan row: " + err.Error())
        }
        // role.CreatedAt = time.Now().Add(24 * time.Hour)
        // role.UpdatedAt = time.Now().Add(24 * time.Hour)
        data = append(data, role)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }
    return data, err
}

func (repo *repoRole) GetByID(id string) (domain.Role, error) {
    row := repo.DB.QueryRow("SELECT * FROM roles where id=$1", id)
    fmt.Println(id)
    
    var data domain.Role
    
    err := row.Scan(&data.ID, &data.RoleName, &data.RoleDesc, &data.AccessPermission, &data.CreatedAt, &data.UpdatedAt)
    if err != nil {
        panic("Failed to scan row: " + err.Error())
    }
    // data.CreatedAt = time.Now().Add(24 * time.Hour)
    // data.UpdatedAt = time.Now().Add(24 * time.Hour)
    
    if err := row.Err(); err != nil {
        return domain.Role{}, err
    }
    // fmt.Println(data)
    return data, err
}

func (repo *repoRole) Create(role *domain.Role) error {
    row, err := repo.DB.Exec("INSERT INTO roles (role_name, role_desc, access_permission) values ($1, $2, $3)", role.RoleName, role.RoleDesc, role.AccessPermission)
    fmt.Println(row)
    return err
}

func (repo *repoRole) Update(id string, role *domain.Role) (affect int64, err error) {
    result, err := repo.DB.Exec("UPDATE roles SET role_name = $1, role_desc = $2, access_permission = $3 where id = $4", 
    role.RoleName, role.RoleDesc, role.AccessPermission, id)
    
    if err != nil {
        return 0, err
    } 

    affect, err = result.RowsAffected()


    if affect == 0 {
        err = fmt.Errorf("resource not found or id that input is false")
        return 
    } else if affect != 1 {
        err = fmt.Errorf("error function, too many deleted resource. Affected : %d", affect)
        return
    }

    return

}
func (repo *repoRole) Delete(id string) (affect int64, err error) {
    result, err := repo.DB.Exec("DELETE FROM roles where id = $1", id)
    if  err != nil {
        return -1, err
    }
    affect, err = result.RowsAffected()

    if  err != nil {
        return
    }

    if affect == 0 {
        err = fmt.Errorf("resource not found or id that input is false")
        return 
    } else if affect != 1 {
        err = fmt.Errorf("error function, too many deleted resource. Affected : %d", affect)
        return 

    }
    return
}
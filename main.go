package main

import (
	"database/sql"
	"fmt"
	httpRoutes "go-clean-architecture-by-ahr/transport/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("env.yaml")
	if err := viper.ReadInConfig(); err != nil {
        panic(err)
    }

	// postgresql
	dbHost := viper.GetString("database.host")
	dbUser := viper.GetString("database.username")
	dbPort := viper.GetInt("database.port")
	dbPass := viper.GetString("database.password")
	dbName := viper.GetString("database.name")
	connection := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		fmt.Println(err)
	}
	
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	e := echo.New()
	httpRoutes.StartHttp(e, db)

	port := viper.GetInt("server.port")
    e.Start(fmt.Sprintf(":%d", port))
	
}
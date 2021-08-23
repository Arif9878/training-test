package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Arif9878/stockbit-test/question-2/cmd/server"
	"github.com/Arif9878/stockbit-test/question-2/helper"
	"github.com/Arif9878/stockbit-test/question-2/omdbapi"
	"github.com/DATA-DOG/go-sqlmock"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Run() error {
	ApiKey := helper.GetEnv("API_KEY_OMDB")
	omdbapi := omdbapi.Init(ApiKey)

	// dbHost := helper.GetEnv("DB_HOST")
	// dbPort := helper.GetEnv("DB_PORT")
	// dbUser := helper.GetEnv("DB_USER")
	// dbPass := helper.GetEnv("DB_PASSWORD")
	// dbName := helper.GetEnv("DB_NAME")
	// connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	// val := url.Values{}
	// val.Add("parseTime", "1")
	// val.Add("loc", "Asia/Jakarta")
	// dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	// db, err := sql.Open(`mysql`, dsn)
	db, _, err := sqlmock.New()
	if err != nil {
		fmt.Printf("failed to connect to database: %s", err)
		return err
	}
	defer db.Close()

	r := server.Routing(db, omdbapi)
	return server.Start(r)
}

func main() {
	fmt.Println("Server Http and Grpc Run at 8080")
	if err := Run(); err != nil {
		os.Exit(1)
	}
}

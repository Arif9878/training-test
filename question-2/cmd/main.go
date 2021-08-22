package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Arif9878/stockbit-test/question-2/cmd/server"
	"github.com/Arif9878/stockbit-test/question-2/helper"
	"github.com/Arif9878/stockbit-test/question-2/imdbapi"
	"github.com/DATA-DOG/go-sqlmock"
	// imdb_grpc "github.com/Arif9878/stockbit-test/question-2/transport/grpc"
)

var db *sql.DB

func Run() error {
	ApiKey := helper.GetEnv("API_KEY_IMDB")
	imdbapi := imdbapi.Init(ApiKey)
	db, _, err := sqlmock.New()
	if err != nil {
		fmt.Printf("failed to connect to database: %s", err)
		return err
	}
	r := server.Routing(db, imdbapi)
	return server.Start(r)
}

func main() {
	fmt.Println("Server Http and Grpc Run at 8080")
	if err := Run(); err != nil {
		os.Exit(1)
	}
}

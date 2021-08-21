package main

import (
	"os"

	"github.com/Arif9878/stockbit-test/question-2/cmd/server"
	"github.com/Arif9878/stockbit-test/question-2/helper"
	"github.com/Arif9878/stockbit-test/question-2/imdbapi"
)

func run() error {
	ApiKey := helper.GetEnv("API_KEY_IMDB")
	imdbapi := imdbapi.Init(ApiKey)

	r := server.Routing(imdbapi)
	return server.Start(r)
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

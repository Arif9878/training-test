package repository_test

import (
	"testing"

	"github.com/Arif9878/stockbit-test/question-2/models"
	"github.com/Arif9878/stockbit-test/question-2/omdbapi"
	"github.com/Arif9878/stockbit-test/question-2/repository"
	"github.com/stretchr/testify/assert"
)

func Test_ErrorFetchList(t *testing.T) {
	omdbapi := omdbapi.Init("2354534")
	repo := repository.NewAPIRepository(omdbapi)
	var mockQuery models.QueryParams
	list, err := repo.FetchList(&mockQuery)

	assert.Nil(t, list)
	assert.Error(t, err, "Status Code 401 received from IMDB")

}

func Test_ErrorGetByID(t *testing.T) {
	omdbapi := omdbapi.Init("2354534")
	repo := repository.NewAPIRepository(omdbapi)
	detail, err := repo.GetByID("tt0372784")

	assert.Nil(t, detail)
	assert.Error(t, err, "Status Code 401 received from IMDB")
}

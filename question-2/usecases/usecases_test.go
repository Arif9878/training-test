package usecases_test

import (
	"testing"

	mock "github.com/Arif9878/stockbit-test/question-2/mocks/repository"
	"github.com/Arif9878/stockbit-test/question-2/models"
	"github.com/Arif9878/stockbit-test/question-2/usecases"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

func Test_FetchList(t *testing.T) {
	mockRepo := new(mock.ImdbRepository)
	var mockSearchResponse models.SearchResponse
	var mockQuery models.QueryParams

	mockRepo.On("FetchList", mockQuery).Return(&mockSearchResponse, nil)
	u := usecases.NewImdbUsecase(mockRepo, nil)

	list, err := u.FetchList(&mockQuery)
	mockQuery.Search = "Batman"
	mockQuery.Page = "1"
	assert.NotEmpty(t, mockQuery)
	assert.NoError(t, err)
	assert.Equal(t, list, &mockSearchResponse)
}

func Test_GetByID(t *testing.T) {
	mockRepo := new(mock.ImdbRepository)
	var mockMovie models.MovieDetail
	err := faker.FakeData(&mockMovie)
	assert.NoError(t, err)

	mockRepo.On("GetByID", mockMovie.ImdbID).Return(&mockMovie, nil)
	defer mockRepo.AssertCalled(t, "GetByID", mockMovie.ImdbID)
	u := usecases.NewImdbUsecase(mockRepo, nil)

	a, err := u.GetByID(mockMovie.ImdbID)

	assert.NoError(t, err)
	assert.NotNil(t, a)
}

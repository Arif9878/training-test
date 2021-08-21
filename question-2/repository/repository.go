package repository

import (
	"github.com/Arif9878/stockbit-test/question-2/models"
)

type ImdbRepository interface {
	Fetch(query *models.QueryParams) (*models.SearchResponse, error)
	GetByID(ID string) (*models.MovieDetail, error)
}

package repository

import (
	"github.com/Arif9878/stockbit-test/question-2/imdbapi"
	"github.com/Arif9878/stockbit-test/question-2/models"
)

type ImdbRepository interface {
	FetchList(query *models.QueryParams) (*models.SearchResponse, error)
	GetByID(ID string) (*models.MovieDetail, error)
}

// A ImdbRepository belong to the inteface layer
type imdbRepository struct {
	imdbapi *imdbapi.OmdbApi
}

func NewAPIRepository(imdbapi *imdbapi.OmdbApi) imdbRepository {
	return imdbRepository{imdbapi}
}

func (a *imdbRepository) FetchList(query *models.QueryParams) (*models.SearchResponse, error) {
	results, err := a.imdbapi.Search(query)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (a *imdbRepository) GetByID(ID string) (*models.MovieDetail, error) {
	result, err := a.imdbapi.GetImdbDetail(ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

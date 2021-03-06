package repository

import (
	"github.com/Arif9878/stockbit-test/question-2/models"
	"github.com/Arif9878/stockbit-test/question-2/omdbapi"
)

type ImdbRepository interface {
	FetchList(query *models.QueryParams) (*models.SearchResponse, error)
	GetByID(ID string) (*models.MovieDetail, error)
}

// A ImdbRepository belong to the inteface layer
type imdbRepository struct {
	omdbapi *omdbapi.OmdbApi
	log     LogsRepository
}

func NewAPIRepository(omdbapi *omdbapi.OmdbApi, log LogsRepository) imdbRepository {
	return imdbRepository{omdbapi: omdbapi, log: log}
}

func (a *imdbRepository) FetchList(query *models.QueryParams) (*models.SearchResponse, error) {
	results, err := a.omdbapi.Search(query)
	if err != nil {
		return nil, err
	}

	_ = a.log.StoreLogSearch(query)

	return results, nil
}

func (a *imdbRepository) GetByID(ID string) (*models.MovieDetail, error) {
	result, err := a.omdbapi.GetImdbDetail(ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

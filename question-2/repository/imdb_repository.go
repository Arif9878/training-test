package repository

import (
	"fmt"

	"github.com/Arif9878/stockbit-test/question-2/imdbapi"
	"github.com/Arif9878/stockbit-test/question-2/models"
)

// A ImdbRepository belong to the inteface layer
type ImdbMovieRepository struct {
	imdbapi *imdbapi.OmdbApi
}

func NewAPIRepository(imdbapi *imdbapi.OmdbApi) ImdbMovieRepository {
	return ImdbMovieRepository{imdbapi}
}

func (a *ImdbMovieRepository) Fetch(query *models.QueryParams) (*models.SearchResponse, error) {
	results, err := a.imdbapi.Search(query)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return nil, err
	}

	return results, nil
}

func (a *ImdbMovieRepository) GetByID(ID string) (*models.MovieDetail, error) {
	result, err := a.imdbapi.GetImdbDetail(ID)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return nil, err
	}
	return result, nil
}

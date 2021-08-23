package usecases

import (
	"github.com/Arif9878/stockbit-test/question-2/models"
	"github.com/Arif9878/stockbit-test/question-2/repository"
)

type ImdbUsecase struct {
	ImdbRepo repository.ImdbRepository
}

func NewImdbUsecase(imdbRepo repository.ImdbRepository) ImdbUsecase {
	return ImdbUsecase{ImdbRepo: imdbRepo}
}

func (a *ImdbUsecase) FetchList(query *models.QueryParams) (*models.SearchResponse, error) {
	if query.Search == "" {
		return &models.SearchResponse{}, nil
	}

	results, err := a.ImdbRepo.FetchList(query)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (i *ImdbUsecase) GetByID(ID string) (*models.MovieDetail, error) {

	return i.ImdbRepo.GetByID(ID)
}

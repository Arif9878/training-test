package usecases

import (
	"fmt"

	"github.com/Arif9878/stockbit-test/question-2/models"
	"github.com/Arif9878/stockbit-test/question-2/repository"
)

type ImdbUsecase struct {
	ImdbRepo repository.ImdbRepository
	Log      repository.LogsRepository
}

func NewImdbUsecase(imdbRepo repository.ImdbRepository, Log repository.LogsRepository) ImdbUsecase {
	return ImdbUsecase{ImdbRepo: imdbRepo, Log: Log}
}

func (a *ImdbUsecase) FetchList(query *models.QueryParams) (*models.SearchResponse, error) {
	if query.Search == "" {
		return &models.SearchResponse{}, nil
	}

	err := a.Log.StoreLogSearch(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
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

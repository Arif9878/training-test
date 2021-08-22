package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc/reflection"

	"github.com/Arif9878/stockbit-test/question-2/models"
	imdb_grpc "github.com/Arif9878/stockbit-test/question-2/transport/grpc/imdb_grpc"
	"github.com/Arif9878/stockbit-test/question-2/usecases"
	"google.golang.org/grpc"
)

func NewArticleServerGrpc(gserver *grpc.Server, ImdbUsecase usecases.ImdbUsecase) {

	imdbServer := &server{
		usecase: ImdbUsecase,
	}

	imdb_grpc.RegisterImdbGrpcHandlerServer(gserver, imdbServer)
	reflection.Register(gserver)
}

type server struct {
	usecase usecases.ImdbUsecase
}

func (s *server) transformListImbdRPC(ar *models.SearchResponse) *imdb_grpc.SearchResponse {
	if ar == nil {
		return nil
	}
	var releases []*imdb_grpc.MovieDetail

	// build slice with all the available movie detail
	for _, v := range ar.Search {
		ri := &imdb_grpc.MovieDetail{
			Title: v.Title,
		}

		releases = append(releases, ri)
	}

	return &imdb_grpc.SearchResponse{
		ListResults: releases,
	}
}

func (s *server) transformDetailImbdRPC(ar *models.MovieDetail) *imdb_grpc.MovieDetail {

	if ar == nil {
		return nil
	}

	res := &imdb_grpc.MovieDetail{
		Title:  ar.Title,
		Year:   ar.Year,
		ImdbID: ar.ImdbID,
		Type:   ar.Type,
		Poster: ar.Poster,
	}
	return res
}

func (s *server) FetchList(ctx context.Context, in *imdb_grpc.FetchRequest) (*imdb_grpc.SearchResponse, error) {
	var query = &models.QueryParams{}
	if in != nil {
		query.Search = in.Search
		query.Page = in.Page
	}

	list, err := s.usecase.FetchList(query)
	if err != nil {
		return nil, err
	}
	result := s.transformListImbdRPC(list)
	return result, nil
}

func (s *server) GetByID(ctx context.Context, in *imdb_grpc.SingleRequest) (*imdb_grpc.MovieDetail, error) {
	ID := ""
	if in != nil {
		ID = in.Id
	}
	ar, err := s.usecase.GetByID(ID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	res := s.transformDetailImbdRPC(ar)
	return res, nil
}

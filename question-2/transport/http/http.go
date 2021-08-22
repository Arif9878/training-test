package http

import (
	"database/sql"
	"net/http"

	"github.com/Arif9878/stockbit-test/question-2/models"
	"github.com/Arif9878/stockbit-test/question-2/omdbapi"
	"github.com/Arif9878/stockbit-test/question-2/repository"
	imdb_grpc "github.com/Arif9878/stockbit-test/question-2/transport/grpc"
	middleware "github.com/Arif9878/stockbit-test/question-2/transport/http/middleware"
	"github.com/Arif9878/stockbit-test/question-2/usecases"
	"github.com/go-chi/render"
	"google.golang.org/grpc"
)

// HTTP defines work unit http handlers interface
type HTTP interface {
	List(w http.ResponseWriter, r *http.Request)
	Detail(w http.ResponseWriter, r *http.Request)
}

type apiHTTP struct {
	svc usecases.ImdbUsecase
}

// NewWorkUnitHTTP returns ne WorkUnitHTTP struct instance
func NewImdbHTTP(grpc *grpc.Server, db *sql.DB, omdbapi *omdbapi.OmdbApi) HTTP {
	repo := repository.NewAPIRepository(omdbapi)
	log := repository.NewLogsRepository(db)
	svc := usecases.NewImdbUsecase(&repo, log)
	imdb_grpc.NewArticleServerGrpc(grpc, svc)

	return apiHTTP{svc: svc}
}

// paramsDecoder decode http request to construct ListWorkUnitRequest
func paramsDecoder(r *http.Request) *models.QueryParams {
	search := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	listRequest := &models.QueryParams{Search: search, Page: page}
	return listRequest
}

// List is handler for list imdb
func (h apiHTTP) List(w http.ResponseWriter, r *http.Request) {
	listRequest := paramsDecoder(r)
	data, err := h.svc.FetchList(listRequest)
	if err != nil {
		render.Render(w, r, middleware.ErrBadRequest)
		return
	}
	render.Respond(w, r, data)
	return
}

// Detail is handler for list imdb
func (h apiHTTP) Detail(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value(contextKeyRequestID).(string)
	data, err := h.svc.GetByID(ID)
	if err != nil {
		render.Render(w, r, middleware.ErrBadRequest)
		return
	}
	render.Respond(w, r, data)
	return
}

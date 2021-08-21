package server

import (
	"github.com/Arif9878/stockbit-test/question-2/imdbapi"
	httpApi "github.com/Arif9878/stockbit-test/question-2/transport/http"
	"github.com/go-chi/chi"
)

// Routing setup api routing
func Routing(imdbapi *imdbapi.OmdbApi) chi.Router {

	r := chi.NewRouter()

	// register v1 api path group
	r.Route("/", func(r chi.Router) {
		httpApi.RegisterHandlers(r, imdbapi)
	})

	return r
}

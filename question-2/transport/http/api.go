package http

import (
	"context"
	"net/http"

	"github.com/Arif9878/stockbit-test/question-2/imdbapi"
	middleware "github.com/Arif9878/stockbit-test/question-2/transport/http/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type contextID string

const (
	contextKeyRequestID contextID = "ID"
)

// RegisterHandlers registers handlers for specified path
func RegisterHandlers(r chi.Router, imdbapi *imdbapi.OmdbApi) {
	r.Mount("/imdb", RegisterHTTPHandlers(NewImdbHTTP(imdbapi)))
}

// RegisterHTTPHandlers registers http handlers for work unit endpoint
func RegisterHTTPHandlers(http HTTP) http.Handler {
	r := chi.NewRouter()

	r.Get("/", http.List)
	r.Route("/{ID}", func(r chi.Router) {
		r.Use(RequestContextID)
		r.Get("/", http.Detail)
	})
	return r
}

// RequestContextID register get id
func RequestContextID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ID := chi.URLParam(r, "ID")
		if ID == "" {
			render.Render(w, r, middleware.BadRequest(nil, "ID is required"))
			return
		}
		ctx := context.WithValue(r.Context(), contextKeyRequestID, ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

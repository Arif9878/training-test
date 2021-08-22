package server

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/Arif9878/stockbit-test/question-2/omdbapi"
	httpApi "github.com/Arif9878/stockbit-test/question-2/transport/http"
	"github.com/go-chi/chi"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

// Routing setup api routing
func Routing(db *sql.DB, omdbapi *omdbapi.OmdbApi) chi.Router {
	grpcServ := grpc.NewServer()
	r := chi.NewRouter()

	// register api path group
	r.Route("/", func(r chi.Router) {
		httpApi.RegisterHandlers(r, grpcServ, db, omdbapi)
	})

	// Handler Http and Grpc
	mixedHandler := NewHTTPandGRPC(r, grpcServ)
	http2Server := &http2.Server{}
	http1Server := &http.Server{Handler: h2c.NewHandler(mixedHandler, http2Server)}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	grpcErrors := make(chan error, 1)

	go func() {
		grpcErrors <- http1Server.Serve(lis)
	}()

	select {
	case err := <-grpcErrors:
		fmt.Printf("server closed %s", err)
	}

	return r
}

func NewHTTPandGRPC(httpHand http.Handler, grpcHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("content-type"), "application/grpc") {
			grpcHandler.ServeHTTP(w, r)
			return
		}
		httpHand.ServeHTTP(w, r)
	})
}

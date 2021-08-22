package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
)

const (
	port = ":8080"
)

// Start to start the server
func Start(r http.Handler) error {
	serverErrors := make(chan error, 1)

	server := http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:

		ctx, cancel := context.WithTimeout(context.Background(), 10)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			err = server.Close()
			return err
		}

		switch {
		case sig == syscall.SIGSTOP:
			return errors.New("integrity issue caused shutdown")
		case err != nil:
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil
}

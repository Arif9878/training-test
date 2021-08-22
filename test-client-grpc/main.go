package main

import (
	"context"
	"log"

	"github.com/Arif9878/stockbit-test/question-2/transport/grpc/imdb_grpc"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	// Set up a connection to the server.
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := imdb_grpc.NewImdbGrpcHandlerClient(conn)

	detail, err := c.GetByID(context.Background(), &imdb_grpc.SingleRequest{Id: "tt0372784"})
	list, err := c.FetchList(context.Background(), &imdb_grpc.FetchRequest{Search: "Batman", Page: "1"})
	if err != nil {
		log.Fatalf("Error when calling: %s", err)
	}
	log.Printf("Response detail from server: %s", detail)
	log.Printf("Response list from server: %s", list.GetListResults())
}

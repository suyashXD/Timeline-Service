package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/suyashXD/Timeline-Service/internal/graphql/generated"
	"github.com/suyashXD/Timeline-Service/internal/graphql/resolvers"
	"github.com/suyashXD/Timeline-Service/internal/grpc/client"
	"github.com/suyashXD/Timeline-Service/internal/service"
	"github.com/suyashXD/Timeline-Service/internal/store"
)

const (
	defaultPort      = "8080"
	defaultGRPCAddr  = "localhost:50051"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	grpcAddr := os.Getenv("GRPC_ADDR")
	if grpcAddr == "" {
		grpcAddr = defaultGRPCAddr
	}

	// Initialize the data store
	dataStore := store.NewDataStore()

	// Initialize the gRPC client
	postClient, err := client.NewPostServiceClient(grpcAddr)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer postClient.Close()

	// Initialize the timeline service
	timelineService := service.NewTimelineService(dataStore, postClient)

	// Initialize the resolver
	resolver := resolvers.NewRoot(timelineService)

	// Create GraphQL server handler
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	// Playground for testing
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("🚀 GraphQL server ready at http://localhost:%s/", port)
	log.Printf("🎮 GraphQL playground available at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

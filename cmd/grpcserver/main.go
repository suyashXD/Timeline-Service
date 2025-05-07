package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/suyashXD/Timeline-Service/internal/grpc/proto/notification"
	"github.com/suyashXD/Timeline-Service/internal/grpc/server"
)

const defaultPort = "50051"

func main() {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = defaultPort
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register service
	mockService := server.NewMockPostService()
	notification.RegisterPostServiceServer(grpcServer, mockService)

	log.Printf("🚀 gRPC server listening on :%s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

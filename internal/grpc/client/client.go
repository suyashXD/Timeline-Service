package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/suyashXD/Timeline-Service/internal/grpc/proto/notification"
)

// Wraps the gRPC client for post service
type PostServiceClient struct {
	client pb.PostServiceClient
	conn   *grpc.ClientConn
}

// Creates a new gRPC client connected to the post service
func NewPostServiceClient(serverAddr string) (*PostServiceClient, error) {
	// Sets up connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	conn, err := grpc.DialContext(
		ctx,
		serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Printf("Failed to connect to gRPC server: %v", err)
		return nil, err
	}
	
	client := pb.NewPostServiceClient(conn)
	
	return &PostServiceClient{
		client: client,
		conn:   conn,
	}, nil
}

// Close closes the gRPC connection
func (c *PostServiceClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// Fetches posts for a specific user
func (c *PostServiceClient) ListPostsByUser(ctx context.Context, userID string) ([]*pb.Post, error) {
	resp, err := c.client.ListPostsByUser(ctx, &pb.ListPostsRequest{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	
	return resp.Posts, nil
}

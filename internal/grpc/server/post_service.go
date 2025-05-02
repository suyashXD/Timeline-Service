package server

import (
    "context"
    "time"

    pb "github.com/suyashXD/Timeline-Service/internal/grpc/proto/notification"
)

// MockPostService implements the PostServiceServer interface
type MockPostService struct {
    pb.UnimplementedPostServiceServer
    userPosts map[string][]*pb.Post
}

// NewMockPostService initializes the service with mock data
func NewMockPostService() *MockPostService {
    now := time.Now().Unix()
    return &MockPostService{
        userPosts: map[string][]*pb.Post{
            "user1": {
                {Id: "p1", AuthorId: "user1", Content: "User1 Post 1", Timestamp: now - 60},
                {Id: "p2", AuthorId: "user1", Content: "User1 Post 2", Timestamp: now - 30},
                {Id: "p3", AuthorId: "user1", Content: "User1 Post 3", Timestamp: now - 20},
                {Id: "p4", AuthorId: "user1", Content: "User1 Post 4", Timestamp: now - 5},
            },
            "user2": {
                {Id: "p5", AuthorId: "user2", Content: "User2 Post 1", Timestamp: now - 120},
                {Id: "p6", AuthorId: "user2", Content: "User2 Post 2", Timestamp: now - 30},
            },
            "user3": {
                {Id: "p7", AuthorId: "user3", Content: "User3 Post 1", Timestamp: now - 10},
                {Id: "p8", AuthorId: "user3", Content: "User3 Post 2", Timestamp: now - 8},
                {Id: "p9", AuthorId: "user3", Content: "User3 Post 3", Timestamp: now - 4},
            },
            "user4": {
                {Id: "p10", AuthorId: "user4", Content: "User4 Post 1", Timestamp: now - 5},
                {Id: "p11", AuthorId: "user4", Content: "User4 Post 2", Timestamp: now - 2},
                {Id: "p12", AuthorId: "user4", Content: "User4 Post 3", Timestamp: now - 1},
            },
            "user5": {
                {Id: "p13", AuthorId: "user5", Content: "User5 Post 1", Timestamp: now - 15},
                {Id: "p14", AuthorId: "user5", Content: "User5 Post 2", Timestamp: now - 10},
            },
        },
    }
}

// Implements: ListPostsByUser
func (s *MockPostService) ListPostsByUser(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
    posts := s.userPosts[req.UserId]
    return &pb.ListPostsResponse{Posts: posts}, nil
}

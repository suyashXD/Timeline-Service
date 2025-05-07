package service

import (
	"context"
	"sort"
	"sync"

	"github.com/suyashXD/Timeline-Service/internal/grpc/client"
	"github.com/suyashXD/Timeline-Service/internal/models"
	"github.com/suyashXD/Timeline-Service/internal/store"
	pb "github.com/suyashXD/Timeline-Service/internal/grpc/proto/notification"
)

// Handles the business logic for generating user timelines
type TimelineService struct {
	store       *store.DataStore
	postClient  *client.PostServiceClient
}

// Creates a new timeline service
func NewTimelineService(store *store.DataStore, postClient *client.PostServiceClient) *TimelineService {
	return &TimelineService{
		store:      store,
		postClient: postClient,
	}
}

// Retrieves the timeline for a specific user
func (s *TimelineService) GetUserTimeline(ctx context.Context, userID string) ([]*models.GraphQLPost, error) {
	// Gets the list of users that the specified user follows
	followedUsers := s.store.GetFollowing(userID)
	
	// Waitgroup for concurrent fetching
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	allPosts := make([]*pb.Post, 0)
	errors := make([]error, 0)
	
	// Fetches posts from each followed user concurrently
	for _, followedUserID := range followedUsers {
		wg.Add(1)
		go func(uid string) {
			defer wg.Done()
			
			// Fetches posts from the gRPC service
			posts, err := s.postClient.ListPostsByUser(ctx, uid)
			
			mu.Lock()
			defer mu.Unlock()
			
			if err != nil {
				errors = append(errors, err)
				return
			}
			
			allPosts = append(allPosts, posts...)
		}(followedUserID)
	}
	
	wg.Wait()
	
	sort.Slice(allPosts, func(i, j int) bool {
		return allPosts[i].Timestamp > allPosts[j].Timestamp
	})
	
	// Limit to 20 posts
	if len(allPosts) > 20 {
		allPosts = allPosts[:20]
	}
	
	// Convert to GraphQL post format
	result := make([]*models.GraphQLPost, len(allPosts))
	for i, post := range allPosts {
		result[i] = &models.GraphQLPost{
			ID:        post.Id,
			AuthorID:  post.AuthorId,
			Content:   post.Content,
			Timestamp: int(post.Timestamp),
		}
	}
	
	return result, nil
}

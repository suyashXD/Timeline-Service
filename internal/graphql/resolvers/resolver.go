package resolvers

import (
    "context"
    // "sort"
    // "sync"

    // pb "github.com/suyashXD/Timeline-Service/internal/grpc/proto/notification"
    "github.com/suyashXD/Timeline-Service/internal/service"
    // "github.com/suyashXD/Timeline-Service/internal/models"

)

// PostResolver maps gRPC Post to GraphQL Post
type PostResolver struct {
    ID        string
    Content   string
    Timestamp int
    AuthorID  string
}

type Resolver struct {
    timelineService *service.TimelineService
}

func New(timelineService *service.TimelineService) *Resolver {
    return &Resolver{
        timelineService: timelineService,
    }
}

func (r *Resolver) GetTimeline(ctx context.Context, userID string) ([]*PostResolver, error) {
    // Use the TimelineService to get the timeline
    posts, err := r.timelineService.GetUserTimeline(ctx, userID)
    if err != nil {
        return nil, err
    }
    
    // Convert to GraphQL format
    var timeline []*PostResolver
    for _, p := range posts {
        timeline = append(timeline, &PostResolver{
            ID:        p.ID,
            Content:   p.Content,
            Timestamp: p.Timestamp,
            AuthorID:  p.AuthorID,
        })
    }

    return timeline, nil
}
package resolvers

import (
    "context"
    "sort"
    "time"
    "sync"

    pb "github.com/suyashXD/Timeline-Service/internal/grpc/proto/notification"
    "google.golang.org/grpc"
)

// PostResolver maps gRPC Post to GraphQL Post
type PostResolver struct {
    ID        string
    Content   string
    Timestamp int
    AuthorID  string
}

type Resolver struct {
    GRPCClient pb.PostServiceClient
    FollowMap  map[string][]string
}

func (r *Resolver) GetTimeline(ctx context.Context, userID string) ([]*PostResolver, error) {
    followedUsers := r.FollowMap[userID]
    var wg sync.WaitGroup
    mu := sync.Mutex{}
    var posts []*pb.Post

    for _, uid := range followedUsers {
        wg.Add(1)
        go func(uid string) {
            defer wg.Done()
            resp, err := r.GRPCClient.ListPostsByUser(ctx, &pb.ListPostsRequest{UserId: uid})
            if err != nil {
                return
            }
            mu.Lock()
            posts = append(posts, resp.Posts...)
            mu.Unlock()
        }(uid)
    }

    wg.Wait()

    // Sort posts by timestamp (newest first)
    sort.Slice(posts, func(i, j int) bool {
        return posts[i].Timestamp > posts[j].Timestamp
    })

    // Limit to 20
    if len(posts) > 20 {
        posts = posts[:20]
    }

    // Convert to GraphQL format
    var timeline []*PostResolver
    for _, p := range posts {
        timeline = append(timeline, &PostResolver{
            ID:        p.Id,
            Content:   p.Content,
            Timestamp: int(p.Timestamp),
            AuthorID:  p.AuthorId,
        })
    }

    return timeline, nil
}

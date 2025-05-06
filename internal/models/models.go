package models

// User represents a user in the system
type User struct {
	ID       string
	Username string
	// Add other user fields as needed
}

// Post represents a social media post
type Post struct {
	ID        string
	AuthorID  string
	Content   string
	Timestamp int64
}

// GraphQLPost represents a post in GraphQL response format
type GraphQLPost struct {
	ID        string `json:"id"`
	AuthorID  string `json:"authorId"`
	Content   string `json:"content"`
	Timestamp int    `json:"timestamp"`
}

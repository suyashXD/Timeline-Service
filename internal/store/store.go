package store

import (
	"sync"

	"github.com/suyashXD/Timeline-Service/internal/models"
)

// DataStore holds all in-memory data for the application
type DataStore struct {
	Users    map[string]*models.User
	Follows  map[string][]string // key: userID, value: list of followed userIDs
	mu       sync.RWMutex
}

// NewDataStore creates a new in-memory data store with mock data
func NewDataStore() *DataStore {
	ds := &DataStore{
		Users:    make(map[string]*models.User),
		Follows:  make(map[string][]string),
	}
	
	// Initialize with mock data
	ds.initMockData()
	
	return ds
}

// initMockData populates the store with mock data
func (ds *DataStore) initMockData() {
	// Create mock users
	users := []models.User{
		{ID: "user1", Username: "alice"},
		{ID: "user2", Username: "bob"},
		{ID: "user3", Username: "charlie"},
		{ID: "user4", Username: "dave"},
		{ID: "user5", Username: "eve"},
	}
	
	for _, u := range users {
		user := u // Create a new variable to avoid pointer issues
		ds.Users[u.ID] = &user
	}
	
	// Create mock follower relationships
	ds.Follows = map[string][]string{
		"user1": {"user2", "user3", "user4"},
		"user2": {"user1", "user5"},
		"user3": {"user1", "user2", "user4"},
		"user4": {"user5"},
		"user5": {"user1", "user3"},
	}
}

// GetFollowing returns the list of users that the given user follows
func (ds *DataStore) GetFollowing(userID string) []string {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	
	following, exists := ds.Follows[userID]
	if !exists {
		return []string{}
	}
	
	// Return a copy to avoid concurrency issues
	result := make([]string, len(following))
	copy(result, following)
	return result
}

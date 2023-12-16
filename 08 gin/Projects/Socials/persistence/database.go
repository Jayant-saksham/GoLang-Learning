package persistence

import (
	"Socials/model"
	"sync"
)

type Database struct {
	users    map[string]*model.User
	posts    map[string]*model.Post
	comments map[string]*model.Comment
	mu       sync.RWMutex
}

func NewDatabase() *Database {
	return &Database{
		users:    make(map[string]*model.User),
		posts:    make(map[string]*model.Post),
		comments: make(map[string]*model.Comment),
	}
}

func (db *Database) GetUserMap() map[string]*model.User {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.users
}

func (db *Database) GetPostMap() map[string]*model.Post {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.posts
}

func (db *Database) GetCommentMap() map[string]*model.Comment {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.comments
}

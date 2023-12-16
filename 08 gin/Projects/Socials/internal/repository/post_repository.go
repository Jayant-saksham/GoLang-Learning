
package persistence

import (
    "errors"
    "Socials/model"
    "sync"
)

type PostRepository struct {
    db *Database
}

func NewPostRepository(db *Database) *PostRepository {
    return &PostRepository{db: db}
}


func (pr *PostRepository) Create(post *model.Post) error {
    pr.db.mu.Lock()
    defer pr.db.mu.Unlock()

    if _, exists := pr.db.posts[post.ID]; exists {
        return errors.New("post ID already taken")
    }
    pr.db.posts[post.ID] = post
    return nil
}


func (pr *PostRepository) GetByID(postID string) (*model.Post, error) {
    pr.db.mu.RLock()
    defer pr.db.mu.RUnlock()

    post, exists := pr.db.posts[postID]
    if !exists {
        return nil, errors.New("post not found")
    }
    return post, nil
}


func (pr *PostRepository) Update(post *model.Post) error {
    pr.db.mu.Lock()
    defer pr.db.mu.Unlock()

    _, exists := pr.db.posts[post.ID]
    if !exists {
        return errors.New("post not found")
    }

    pr.db.posts[post.ID] = post
    return nil
}

func (pr *PostRepository) Delete(postID string) error {
    pr.db.mu.Lock()
    defer pr.db.mu.Unlock()

    _, exists := pr.db.posts[postID]
    if !exists {
        return errors.New("post not found")
    }

    delete(pr.db.posts, postID)
    return nil
}

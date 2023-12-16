package persistence

import (
    "errors"
    "Socials/model"
    "sync"
)

type CommentRepository struct {
    db *Database
}

func NewCommentRepository(db *Database) *CommentRepository {
    return &CommentRepository{db: db}
}

func (cr *CommentRepository) Create(comment *model.Comment) error {
    cr.db.mu.Lock()
    defer cr.db.mu.Unlock()

    if _, exists := cr.db.comments[comment.ID]; exists {
        return errors.New("comment ID already taken")
    }

    cr.db.comments[comment.ID] = comment
    return nil
}

func (cr *CommentRepository) GetByID(commentID string) (*model.Comment, error) {
    cr.db.mu.RLock()
    defer cr.db.mu.RUnlock()

    comment, exists := cr.db.comments[commentID]
    if !exists {
        return nil, errors.New("comment not found")
    }
    return comment, nil
}

func (cr *CommentRepository) Update(comment *model.Comment) error {
    cr.db.mu.Lock()
    defer cr.db.mu.Unlock()

    _, exists := cr.db.comments[comment.ID]
    if !exists {
        return errors.New("comment not found")
    }

    cr.db.comments[comment.ID] = comment
    return nil
}

func (cr *CommentRepository) Delete(commentID string) error {
    cr.db.mu.Lock()
    defer cr.db.mu.Unlock()

    _, exists := cr.db.comments[commentID]
    if !exists {
        return errors.New("comment not found")
    }

    delete(cr.db.comments, commentID)
    return nil
}

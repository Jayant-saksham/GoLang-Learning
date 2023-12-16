package persistence

import (
    "errors"
    "Socials/model"
    "sync"
)

type UserRepository struct {
    db *Database
}

func NewUserRepository(db *Database) *UserRepository {
    return &UserRepository{db: db}
}

func (ur *UserRepository) Create(user *model.User) error {
    ur.db.mu.Lock()
    defer ur.db.mu.Unlock()

    if _, exists := ur.db.users[user.Username]; exists {
        return errors.New("username already taken")
    }

    user.ID = generateID()
    ur.db.users[user.ID] = user
    return nil
}

func (ur *UserRepository) GetByID(userID string) (*model.User, error) {
    ur.db.mu.RLock()
    defer ur.db.mu.RUnlock()

    user, exists := ur.db.users[userID]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}

func (ur *UserRepository) Update(user *model.User) error {
    ur.db.mu.Lock()
    defer ur.db.mu.Unlock()

    _, exists := ur.db.users[user.ID]
    if !exists {
        return errors.New("user not found")
    }

    ur.db.users[user.ID] = user
    return nil
}


func (ur *UserRepository) Delete(userID string) error {
    ur.db.mu.Lock()
    defer ur.db.mu.Unlock()

    _, exists := ur.db.users[userID]
    if !exists {
        return errors.New("user not found")
    }

    delete(ur.db.users, userID)
    return nil
}

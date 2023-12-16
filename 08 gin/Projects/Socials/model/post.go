package model

import "time"

type Post struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	Caption      string    `json:"caption"`
	ImageURL     string    `json:"image_url"`
	Location     string    `json:"location"`
	LikeCount    int       `json:"like_count"`
	CommentCount int       `json:"comment_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

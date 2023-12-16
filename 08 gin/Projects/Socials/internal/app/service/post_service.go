package service

import (
	"Socials/model"
	"time"
)

func NewPost() *model.Post {
	return &model.Post{
		LikeCount:    0,
		CommentCount: 0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

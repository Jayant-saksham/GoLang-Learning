package service

import (
	"time"
	"Socials/model"
)

func NewComment() *model.Comment {
	return &model.Comment{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
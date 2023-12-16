
package service

import (
	"Socials/model"
	"time"
)

func NewUser() *model.User {
	return &model.User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

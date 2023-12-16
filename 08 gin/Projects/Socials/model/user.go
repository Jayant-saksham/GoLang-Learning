package model

import "time"

type User struct {
	ID             	string    `json:"id"`
	Username       	string    `json:"username"`
	PasswordHash   	string    `json:"-"`
	Bio            	string    `json:"bio"`
	ProfilePicture 	string    `json:"profile_picture"`
	PostsCount     	int       `json:"posts_count"`
	FollowersCount 	int       `json:"followers_count"`
	FollowingCount 	int       `json:"following_count"`
	CreatedAt		time.Time `json:"created_at"`
	UpdatedAt		time.Time `json:"updated_at"`
 }
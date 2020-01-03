package types

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Body string `json:"body"`
	User string `json:"user"`
}

func NewPost(title, body, user string) *Post {
	return &Post{
		Title: title,
		Body:  body,
		User:  user,
	}
}

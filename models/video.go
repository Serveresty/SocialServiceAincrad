package models

import "time"

type Video struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AuthorID    string    `json:"author_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	CreatedAt   time.Time `json:"created_at"`
	Views       int       `json:"views"`
	Likes       int       `json:"likes"`
	Comments    []Comment `json:"comments"`
	Filename    string    `json:"filename"`
	PreviewName string    `json:"preview_name"`
	Preview     string    `json:"preview"`
}

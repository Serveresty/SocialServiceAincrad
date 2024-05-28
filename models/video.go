package models

import "time"

type Video struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	CreatedAt   time.Time `json:"created_at"`
	Views       int       `json:"views"`
	Likes       int       `json:"likes"`
	Comments    []string  `json:"comments"`
	Filename    string    `json:"filename"`
}

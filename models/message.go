package models

import "time"

type Message struct {
	Sender    string    `json:"sender"`
	Receiver  string    `json:"receiver"`
	Messages  string    `json:"messages"`
	CreatedAt time.Time `json:"created_at"`
}

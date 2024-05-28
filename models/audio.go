package models

type Audio struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Filename string `json:"filename"`
	Hash     string `json:"hash"`
}

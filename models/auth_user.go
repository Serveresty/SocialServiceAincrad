package models

type AuthUser struct {
	Email        string   `json:"email"`
	Password     string   `json:"password"`
	Roles        []string `json:"roles"`
	StayLoggedIn bool     `json:"stay_logged"`
}

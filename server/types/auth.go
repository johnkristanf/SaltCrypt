package types

import "time"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDTO struct {
	Email    string
	Password string
}

type FetchUsers struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	Pepper    string    `json:"pepper"`
	CreatedAt time.Time `json:"created_at"`
}
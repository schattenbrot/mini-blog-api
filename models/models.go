package models

import "time"

type Post struct {
	ID        string    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty" validate:"omitempty,min=3,max=40"`
	Text      string    `json:"text,omitempty" validate:"omitempty,min=5,max=700"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// User      string    `json:"user,omitempty" validate:"omitempty"`

type User struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name" validate:"omitempty,min=3,max=20"`
	Email     string    `json:"email" validate:"omitempty,email"`
	Password  string    `json:"password" validate:"omitempty,min=8,max=24"`
	Roles     []string  `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
}

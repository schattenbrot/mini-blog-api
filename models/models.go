package models

import "time"

// Post describes the globally used Post type.
type Post struct {
	ID        string    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty" validate:"omitempty,min=3,max=40"`
	Text      string    `json:"text,omitempty" validate:"omitempty,min=5,max=700"`
	User      string    `json:"user,omitempty" validate:"omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// User describes the globally used User type.
type User struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name" validate:"omitempty,min=3,max=20"`
	Email     string    `json:"email" validate:"omitempty,email"`
	Password  string    `json:"password,omitempty" validate:"omitempty,min=8,max=24"`
	Roles     []string  `json:"roles" validate:"omitempty,dive,eq=user|admin"`
	CreatedAt time.Time `json:"created_at"`
}

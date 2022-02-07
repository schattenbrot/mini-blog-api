package models

import "time"

type Post struct {
	ID        string    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty" validate:"omitempty,min=3,max=40"`
	Text      string    `json:"text,omitempty" validate:"omitempty,min=5,max=700"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

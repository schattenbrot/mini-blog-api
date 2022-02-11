package database

import (
	"github.com/schattenbrot/mini-blog-api/models"
)

type DatabaseRepo interface {
	// InsertPost inserts a post and returns the id or an error if any
	InsertPost(p models.Post) (*string, error)
	GetPosts() ([]*models.Post, error)
	GetPostById(id string) (*models.Post, error)
	GetPostsByPage(page, limit int) ([]*models.Post, error)
	UpdatePost(p models.Post) error
	DeleteOnePost(id string) error

	InsertUser(u models.User) (*string, error)
	GetUserById(id string) (*models.User, error)
	UpdateUser(u models.User) error
	DeleteUser(id string) error
}

package database

import (
	"github.com/schattenbrot/mini-blog-api/models"
)

// DatabaseRepo represents the database repository.
type DatabaseRepo interface {
	InsertPost(p models.Post) (*string, error)
	GetPosts() ([]*models.Post, error)
	GetPostCreator(id string) (string, error)
	GetPostById(id string) (*models.Post, error)
	GetPostsByPage(page, limit int) ([]*models.Post, error)
	UpdatePost(p models.Post) error
	DeleteOnePost(id string) error

	InsertUser(u models.User) (*string, error)
	GetUserRoles(id string) ([]string, error)
	GetUserById(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(u models.User) error
	DeleteUser(id string) error
}

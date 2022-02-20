package middlewares

import (
	"github.com/schattenbrot/mini-blog-api/config"
	"github.com/schattenbrot/mini-blog-api/database"
	"github.com/schattenbrot/mini-blog-api/database/dbrepo"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository is the router repository type.
type Repository struct {
	App *config.AppConfig
	DB  database.DatabaseRepo
}

// Repo is the router repository.
var Repo *Repository

// NewTestDBRepo returns a new repository for testing purposes.
func NewTestRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewTestingRepo(a),
	}
}

// NewMongoDBRepo returns a new instance of a repository for the mongo driver.
func NewMongoDBRepo(a *config.AppConfig, db *mongo.Database) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewMongoDBRepo(a, db),
	}
}

// NewHandlers sets the Repo.
func NewRouter(r *Repository) {
	Repo = r
}

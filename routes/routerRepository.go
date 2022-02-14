package routes

import (
	"github.com/schattenbrot/mini-blog-api/config"
	"github.com/schattenbrot/mini-blog-api/database"
	"github.com/schattenbrot/mini-blog-api/database/dbrepo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	App *config.AppConfig
	DB  database.DatabaseRepo
}

var Repo *Repository

func NewTestRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewTestingRepo(a),
	}
}

func NewMongoDBRepo(a *config.AppConfig, db *mongo.Database) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewMongoDBRepo(a, db),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

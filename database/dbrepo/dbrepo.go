package dbrepo

import (
	"database/sql"

	"github.com/schattenbrot/mini-blog-api/config"
	"github.com/schattenbrot/mini-blog-api/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type mongoDBRepo struct {
	App *config.AppConfig
	DB  *mongo.Database
}

func NewTestingRepo(app *config.AppConfig) database.DatabaseRepo {
	return &testDBRepo{
		App: app,
	}
}

func NewMongoDBRepo(app *config.AppConfig, conn *mongo.Database) database.DatabaseRepo {
	return &mongoDBRepo{
		App: app,
		DB:  conn,
	}
}

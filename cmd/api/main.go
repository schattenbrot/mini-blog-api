package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/schattenbrot/mini-blog-api/config"
	"github.com/schattenbrot/mini-blog-api/controllers"
	"github.com/schattenbrot/mini-blog-api/middlewares"
	"github.com/schattenbrot/mini-blog-api/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Application struct {
	App *config.AppConfig
}

var App *Application

func main() {
	var cfg config.Config
	config.LoadConfig(&cfg)

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	validator := validator.New()

	app := &config.AppConfig{
		Version:         "1.0.0",
		ServerStartTime: time.Now(),
		Config:          cfg,
		Logger:          logger,
		Validator:       validator,
	}
	App = &Application{
		App: app,
	}

	db := openDB()

	repo := controllers.NewMongoDBRepo(app, db)
	controllers.NewHandlers(repo)
	middlewareRepo := middlewares.NewMongoDBRepo(app, db)
	middlewares.NewRouter(middlewareRepo)

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      routes.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println(fmt.Sprintf("Starting server on port %d", cfg.Port))

	err := serve.ListenAndServe()
	if err != nil {
		logger.Fatal("Welp ... uwuff")
	}
}

// openDB creates a new database connection and returns the Database
func openDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(App.App.Config.DB.DSN))
	if err != nil {
		App.App.Logger.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		App.App.Logger.Fatal(err)
	}
	db := client.Database("mini-blog")

	return db
}

package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/schattenbrot/mini-blog-api/controllers"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		MaxAge:         300,
	}))

	r.Get("/", controllers.Repo.StatusHandler)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/posts", postRouter)
		r.Route("/users", userRouter)
	})

	return r
}

func postRouter(r chi.Router) {
	r.Post("/", controllers.Repo.InsertPost)
	r.Get("/", controllers.Repo.GetAllPosts)
	r.Get("/paging", controllers.Repo.GetAllPostsPaginated)
	r.Get("/{id}", controllers.Repo.GetPostById)
	r.Patch("/{id}", controllers.Repo.UpdatePostById)
	r.Delete("/{id}", controllers.Repo.DeletePost)
}

func userRouter(r chi.Router) {
	r.Post("/", controllers.Repo.InsertUser)
	r.Get("/{id}", controllers.Repo.GetUserById)
	r.Patch("/{id}", controllers.Repo.UpdateUserById)
	r.Delete("/{id}", controllers.Repo.DeleteUser)
}

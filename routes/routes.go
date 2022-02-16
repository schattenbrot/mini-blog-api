package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/schattenbrot/mini-blog-api/controllers"
)

// Routes returns a fully configured Mux of the chi-router.
func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", controllers.Repo.StatusHandler)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/posts", postRouter)
		r.Route("/users", userRouter)
	})

	return r
}

func postRouter(r chi.Router) {
	r.With(Repo.Auth).Post("/", controllers.Repo.InsertPost)
	r.With(Repo.Auth).Patch("/{id}", controllers.Repo.UpdatePostById)
	r.With(Repo.Auth).Delete("/{id}", controllers.Repo.DeletePost)

	r.Get("/", controllers.Repo.GetAllPosts)
	r.Get("/paging", controllers.Repo.GetPaginatedPosts)
	r.Get("/{id}", controllers.Repo.GetPostById)
}

func userRouter(r chi.Router) {
	r.Post("/", controllers.Repo.InsertUser)
	r.Post("/login", controllers.Repo.Login)

	r.With(Repo.Auth).Get("/{id}", controllers.Repo.GetUserById)
	r.With(Repo.Auth).Patch("/{id}", controllers.Repo.UpdateUserById)
	r.With(Repo.Auth).Delete("/{id}", controllers.Repo.DeleteUser)
	r.With(Repo.Auth).Get("/logout", controllers.Repo.Logout)
}

package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/schattenbrot/mini-blog-api/controllers"
	"github.com/schattenbrot/mini-blog-api/middlewares"
)

// Routes returns a fully configured Mux of the chi-router.
func Routes(corsAllowedOrigins []string) *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   corsAllowedOrigins,
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
	r.Get("/", controllers.Repo.GetAllPosts)
	r.Get("/paging", controllers.Repo.GetPaginatedPosts)
	r.Get("/{id}", controllers.Repo.GetPostById)

	r.With(middlewares.Repo.Auth).Post("/", controllers.Repo.InsertPost)
	r.With(middlewares.Repo.Auth).With(middlewares.Repo.IsPostCreatorOrAdmin).Patch("/{id}", controllers.Repo.UpdatePostById)
	r.With(middlewares.Repo.Auth).With(middlewares.Repo.IsPostCreatorOrAdmin).Delete("/{id}", controllers.Repo.DeletePost)
}

func userRouter(r chi.Router) {
	r.Post("/", controllers.Repo.InsertUser)
	r.Post("/login", controllers.Repo.Login)

	r.With(middlewares.Repo.Auth).Get("/{id}", controllers.Repo.GetUserById)
	r.With(middlewares.Repo.Auth).With(middlewares.Repo.IsUserOrAdmin).Patch("/{id}", controllers.Repo.UpdateUserById)
	r.With(middlewares.Repo.Auth).With(middlewares.Repo.IsUserOrAdmin).Delete("/{id}", controllers.Repo.DeleteUser)
	r.With(middlewares.Repo.Auth).Get("/logout", controllers.Repo.Logout)
}

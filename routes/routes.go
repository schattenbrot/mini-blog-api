package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/schattenbrot/mini-blog-api/controllers"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/v1", func(r chi.Router) {
		r.Route("/posts", postRouter)
		r.Route("/users", userRouter)
	})

	return r
}
func postRouter(r chi.Router) {
	r.With(Repo.IsAuth).Post("/", controllers.Repo.InsertPost)
	r.With(Repo.IsAuth).Patch("/{id}", controllers.Repo.UpdatePostById)
	r.With(Repo.IsAuth).Delete("/{id}", controllers.Repo.DeletePost)

	r.Get("/", controllers.Repo.GetAllPosts)
	r.Get("/paging", controllers.Repo.GetAllPostsPaginated)
	r.Get("/{id}", controllers.Repo.GetPostById)
}

func userRouter(r chi.Router) {
	r.Post("/", controllers.Repo.InsertUser)
	r.Post("/login", controllers.Repo.LoginUser)

	r.With(Repo.IsAuth).Get("/{id}", controllers.Repo.GetUserById)
	r.With(Repo.IsAuth).Patch("/{id}", controllers.Repo.UpdateUserById)
	r.With(Repo.IsAuth).Delete("/{id}", controllers.Repo.DeleteUser)
	r.With(Repo.IsAuth).Get("/logout", controllers.Repo.LogoutUser)
}

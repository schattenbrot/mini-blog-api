package middlewares

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/schattenbrot/mini-blog-api/utils"
)

// Auth checks if the requests is authorized to access the endpoint.
func (m *Repository) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		issuer, err := utils.GetIssuerFromCookie(r, m.App.Config.CookieName, m.App.Config.JWT)
		if err != nil {
			notAuthenticated(w, err)
			return
		}

		_, err = m.DB.GetUserById(issuer)
		if err != nil {
			notAuthenticated(w, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// notAuthenticated handles the response if the request is not authorized.
func notAuthenticated(w http.ResponseWriter, err error) {
	statusCode := http.StatusUnauthorized
	w.Header().Add("WWW-Authenticate", err.Error())
	w.WriteHeader(statusCode)
}

// IsPostCreatorOrAdmin is a middleware to check if the user got the permission
// to modify or delete the target post.
func (m *Repository) IsPostCreatorOrAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		issuer, err := utils.GetIssuerFromCookie(r, m.App.Config.CookieName, m.App.Config.JWT)
		if err != nil {
			setStatusForbidden(w)
			return
		}

		// check for creator
		postID := chi.URLParam(r, "id")
		creator, err := m.DB.GetPostCreator(postID)
		if err == nil {
			if issuer == creator {
				next.ServeHTTP(w, r)
				return
			}
		}

		// check for admin rights
		userRoles, err := m.DB.GetUserRoles(issuer)
		if err == nil {
			for _, role := range userRoles {
				if role == issuer {
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		setStatusForbidden(w)
	})
}

// IsUserOrAdmin is a middleware to check if the user got the permission to
// change or delete the target user.
func (m *Repository) IsUserOrAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		issuer, err := utils.GetIssuerFromCookie(r, m.App.Config.CookieName, m.App.Config.JWT)
		if err != nil {
			setStatusForbidden(w)
			return
		}

		// check if correct user
		userID := chi.URLParam(r, "id")
		if issuer == userID {
			next.ServeHTTP(w, r)
			return
		}

		// check for admin rights
		userRoles, err := m.DB.GetUserRoles(issuer)
		if err == nil {
			for _, role := range userRoles {
				if role == "admin" {
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		setStatusForbidden(w)
	})
}

// setStatusForbidden sets the status to StatusForbidden
func setStatusForbidden(w http.ResponseWriter) {
	statusCode := http.StatusForbidden
	w.WriteHeader(statusCode)
}

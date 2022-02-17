package routes

import (
	"net/http"

	"github.com/schattenbrot/mini-blog-api/utils"
)

// Auth checks if the requests is authorized to access the endpoint.
func (m *Repository) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		issuer, err := utils.GetIssuerFromCookie(r, m.App.Config.JWT)
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

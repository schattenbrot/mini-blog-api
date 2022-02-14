package routes

import (
	"net/http"

	"github.com/golang-jwt/jwt"
)

func (m *Repository) IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			notAuthenticated(w, err)
			return
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.App.Config.JWT), nil
		})

		if err != nil {
			notAuthenticated(w, err)
			return
		}

		claims := token.Claims.(*jwt.StandardClaims)
		issuer := claims.Issuer

		_, err = m.DB.GetUserById(issuer)
		if err != nil {
			notAuthenticated(w, err)
		}

		next.ServeHTTP(w, r)
	})
}

func notAuthenticated(w http.ResponseWriter, err error) {
	statusCode := http.StatusUnauthorized
	w.Header().Add("WWW-Authenticate", err.Error())
	w.WriteHeader(statusCode)
}
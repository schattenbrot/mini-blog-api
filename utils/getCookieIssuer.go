package utils

import (
	"net/http"

	"github.com/golang-jwt/jwt"
)

// GetIssuerFromCookie is a helper function that takes a request and tht
// JWT_SECRET_TOKEN to retrieve the issuer and an error if any occured.
func GetIssuerFromCookie(r *http.Request, cookieName string, jwtSecret []byte) (string, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return "", err
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	issuer := claims.Issuer
	return issuer, nil
}

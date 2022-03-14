package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// LoginUser is the type for authentication-request bodies
type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=24"`
}

// Login is the handler for logging a user in with the given email and password.
// Sets a cookie if successful or an error message.
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	var loginUser LoginUser
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		errorJSON(w, err)
		return
	}

	err = m.App.Validator.Struct(loginUser)
	if err != nil {
		errorJSON(w, err)
		return
	}

	user, err := m.DB.GetUserByEmail(loginUser.Email)
	if err != nil {
		if err == mongo.ErrNilDocument {
			errorJSON(w, err, http.StatusNotFound)
			return
		}
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		errorJSON(w, err)
		return
	}

	currTime := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.ID,
		ExpiresAt: currTime.Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(m.App.Config.JWT)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     m.App.Config.CookieName,
		Path:     "/",
		Value:    tokenString,
		Expires:  currTime.Add(time.Hour * 24),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		// Secure:   true,
	}

	http.SetCookie(w, cookie)

	type jsonResp struct {
		OK bool `json:"ok"`
	}

	writeJSON(w, http.StatusOK, jsonResp{
		OK: true,
	})
}

// Logout logs the user out
func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     m.App.Config.CookieName,
		Path:     "/",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	type jsonResp struct {
		OK bool `json:"ok"`
	}

	writeJSON(w, http.StatusOK, jsonResp{
		OK: true,
	})
}

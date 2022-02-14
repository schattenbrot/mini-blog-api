package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	ID       string `json:"id" validate:"required,len=24"`
	Password string `json:"password" validate:"required,min=8,max=24"`
}

func (m *Repository) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser LoginUser
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		errorJSON(w, err)
		return
	}
	m.App.Logger.Println(loginUser.ID)

	// validate inputs
	err = m.App.Validator.Struct(loginUser)
	if err != nil {
		errorJSON(w, err)
		return
	}

	user, err := m.DB.GetUserById(loginUser.ID)
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
		Name:     "jwt",
		Path:     "/",
		Value:    tokenString,
		Expires:  currTime.Add(time.Hour * 24),
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

func (m *Repository) LogoutUser(w http.ResponseWriter, r *http.Request) {

}

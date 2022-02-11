package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/schattenbrot/mini-blog-api/helpers"
	"github.com/schattenbrot/mini-blog-api/models"
	"golang.org/x/crypto/bcrypt"
)

func (m *Repository) InsertUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorJSON(w, err)
		return
	}

	user.CreatedAt = time.Now()
	user.Roles = []string{"user"}

	v := validator.New()
	err = v.Struct(user)
	if err != nil {
		errorJSON(w, err)
		return
	}

	passwordValid := helpers.PasswordIsValid(user.Password)
	if !passwordValid {
		err = errors.New("password is not valid")
		errorJSON(w, err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	_, err = m.DB.InsertUser(user)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	type jsonResp struct {
		OK bool `json:"ok"`
	}

	response := jsonResp{
		OK: true,
	}

	err = writeJSON(w, http.StatusCreated, response)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
	}
}

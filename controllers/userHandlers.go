package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/schattenbrot/mini-blog-api/database/dbrepo"
	"github.com/schattenbrot/mini-blog-api/models"
	"github.com/schattenbrot/mini-blog-api/utils"
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

	passwordValid := utils.PasswordIsValid(user.Password)
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

	user.Roles = []string{"user"}

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

func (m *Repository) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := m.DB.GetUserById(id)
	if err != nil {
		if err.Error() == "the provided hex string is not a valid ObjectID" {
			errorJSON(w, err)
			return
		}
		if err.Error() == "mongo: no documents in result" {
			errorJSON(w, err, http.StatusNotFound)
			return
		}
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = writeJSON(w, http.StatusOK, user)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
	}
}

func (m *Repository) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorJSON(w, err)
		return
	}
	user.ID = id

	v := validator.New()
	err = v.Struct(user)
	if err != nil {
		errorJSON(w, err)
		return
	}
	if user.Password != "" {
		passwordValid := utils.PasswordIsValid(user.Password)
		if !passwordValid {
			err = errors.New("password is not valid")
			errorJSON(w, err)
			return
		}
	}

	err = m.DB.UpdateUser(user)
	if err != nil {
		if err.Error() == dbrepo.ErrorDocumentNotFound {
			errorJSON(w, err, http.StatusNotFound)
			return
		}
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = writeJSON(w, http.StatusNoContent)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
	}
}

func (m *Repository) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := m.DB.DeleteUser(id)
	if err != nil {
		if err.Error() == dbrepo.ErrorDocumentNotFound {
			errorJSON(w, err, http.StatusNotFound)
		}
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = writeJSON(w, http.StatusNoContent)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
	}
}

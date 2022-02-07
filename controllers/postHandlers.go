package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/schattenbrot/mini-blog-api/database/dbrepo"
	"github.com/schattenbrot/mini-blog-api/models"
)

func (m *Repository) InsertPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		errorJSON(w, err)
		return
	}

	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	// TODO: validation
	v := validator.New()
	err = v.Struct(post)
	if err != nil {
		errorJSON(w, err)
		return
	}
	if post.Text == "" || post.Title == "" {
		err := errors.New("text and title are required")
		errorJSON(w, err)
		return
	}

	id, err := m.DB.InsertPost(post)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	type jsonResp struct {
		OK bool    `json:"ok"`
		ID *string `json:"id"`
	}

	response := jsonResp{
		OK: true,
		ID: id,
	}

	err = writeJSON(w, http.StatusCreated, response)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
	}
}

func (m *Repository) GetPostById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	post, err := m.DB.GetPostById(id)
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

	err = writeJSON(w, http.StatusOK, post)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
	}
}

func (m *Repository) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := m.DB.GetPosts()
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = writeJSON(w, http.StatusOK, posts)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
	}
}

func (m *Repository) GetAllPostsPaginated(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		errorJSON(w, err)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		errorJSON(w, err)
		return
	}

	posts, err := m.DB.GetPostsByPage(page, limit)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = writeJSON(w, http.StatusOK, posts)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}
}

func (m *Repository) UpdatePostById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		errorJSON(w, err)
		return
	}
	post.ID = id

	v := validator.New()
	err = v.Struct(post)
	if err != nil {
		errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = m.DB.UpdatePost(post)
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

func (m *Repository) DeletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := m.DB.DeleteOnePost(id)
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

package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
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

func (m *Repository) UpdatePostById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		errorJSON(w, err)
		return
	}
	post.ID = id

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

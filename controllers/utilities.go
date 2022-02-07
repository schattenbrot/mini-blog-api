package controllers

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data ...interface{}) error {
	if status == http.StatusNoContent {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		return nil
	}

	js, err := json.Marshal(data[0])
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	writeJSON(w, statusCode, theError)
}

package controllers

import (
	"encoding/json"
	"net/http"
)

// writeJSON is the helperfunction for sending back an HTTP response.
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

// errorJSON is the helper function for creating an error message.
// This internally then runs the writeJSON function to send the HTTP response.
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

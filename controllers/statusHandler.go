package controllers

import (
	"encoding/json"
	"net/http"
	"time"
)

type appStatus struct {
	Status        string        `json:"status"`
	UpSince       time.Time     `json:"up_since"`
	CurrentUptime time.Duration `json:"current_uptime"`
	Environment   string        `json:"environment"`
	Version       string        `json:"version"`
}

// StatusHandler is the handler for getting the apps status information.
func (m *Repository) StatusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := appStatus{
		Status:        "Available",
		UpSince:       m.App.ServerStartTime,
		CurrentUptime: time.Since(m.App.ServerStartTime),
		Environment:   m.App.Config.Env,
		Version:       m.App.Version,
	}

	js, err := json.Marshal(currentStatus)
	if err != nil {
		m.App.Logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

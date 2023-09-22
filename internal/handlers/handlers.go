package handlers

import (
	"github.com/anras5/go-with-docker/internal/config"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Todos API is up and running",
		Version: "1.0.1",
	}

	_ = config.WriteJSON(w, http.StatusOK, payload)

}

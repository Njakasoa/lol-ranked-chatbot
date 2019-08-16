package controllers

import (
	"net/http"
)

// ServerHealthHandler test server health
func ServerHealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status: OK"))
}

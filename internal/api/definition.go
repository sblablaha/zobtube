package api

import (
	"github.com/gorilla/mux"

	httpHandler "github.com/sblablaha/zobtube/internal/handler/http"
)

// Setup defines all API routes in the router
func Setup(router *mux.Router) {
	// health
	router.HandleFunc("/api/health", httpHandler.HealthCheck)

	// videos
	router.HandleFunc("/api/video/{id}", httpHandler.VideoGet)
}

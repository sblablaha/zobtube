package http

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// HealthCheck is just a basic health check returning ok when alive
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Trace().Msg("health endpoint called")
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

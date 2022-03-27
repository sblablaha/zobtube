package http

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// VideoGet output a video details
func VideoGet(w http.ResponseWriter, r *http.Request) {
	log.Trace().Msg("video requested")
	json.NewEncoder(w).Encode(map[string]bool{"not available yet": true})
}

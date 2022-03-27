package http

import (
	"encoding/json"
	"net/http"

	"github.com/sblablaha/zobtube/internal/controller"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

// VideoGet output a video details
func VideoGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Info().Str("id", id).Msg("video requested")

	v, err := controller.GetVideo(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	err = json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Fatal().Err(err).Msg("encoding issue")
		json.NewEncoder(w).Encode(map[string]string{"error": "uanble to encode response"})
		return
	}
}

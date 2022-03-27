package controller

import (
	"github.com/sblablaha/zobtube/internal/model"
)

// GetVideo returns a video from its ID
func GetVideo(id string) (*model.Video, error) {
	return &model.Video{
		Title: "test",
	}, nil
}

package model

import (
	"time"
)

// Video main type
type Video struct {
	Title      string     `json:"title"`
	CreatedAt  time.Time  `json:"created_at"`
	Production Production `json:"production"`
}

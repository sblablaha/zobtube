package webserver

import (
	"net/http"
)

// WebServer is a web server
type WebServer struct {
	server *http.Server
}

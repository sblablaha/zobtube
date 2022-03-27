package webserver

import (
	"fmt"
	"net/http"
	"time"
)

// New creates a new web server from its listening port
func New(port int) (*WebServer) {
	listeingAddress := fmt.Sprintf("0.0.0.0:%d", port)
	srv := &http.Server{
		//Handler:      router,
		Addr:         listeingAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &WebServer{
		server: srv,
	}
}

// AssignHandler assign handler to the web server
func (ws* WebServer) AssignHandler(router http.Handler) {
	ws.server.Handler = router
}

// ListenAndServe listen and serve
func (ws *WebServer) ListenAndServe() error {
	return ws.server.ListenAndServe()
}

package server

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Server holds the dependencies for a HTTP server
type Server struct {
	logger *log.Logger
	router *mux.Router
}

// New returns a new HTTP server
func New(logger *log.Logger) *Server {
	s := &Server{
		logger: logger,
	}

	r := mux.NewRouter()
	s.router = r

	return s
}

// ServeHTTP dispatches the handler registered in the matched route
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

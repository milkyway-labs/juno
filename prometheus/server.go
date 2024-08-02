package prometheus

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	port   int16
	server *http.Server
}

// NewServer returns a new prometheus server instance
func NewServer(port int16) *Server {
	return &Server{
		port: port,
	}
}

// Start starts the prometheus server
func (s *Server) Start() {
	// Server already started
	if s.server != nil {
		return
	}

	http.Handle("/metrics", promhttp.Handler())
	s.server = &http.Server{
		Addr:              fmt.Sprintf(":%d", s.port),
		ReadHeaderTimeout: 3 * time.Second,
	}
	go s.server.ListenAndServe()
}

// Stop stops the prometheus server
func (s *Server) Stop() {
	if s.server != nil {
		s.server.Close()
		s.server = nil
	}
}

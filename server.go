package main

import (
	"github.com/hectane/go-asyncserver"
)

// HTTP server providing the web interface and API.
type Server struct {
	server   *server.AsyncServer
	channels []*Channel
}

// Create a new server instance.
func NewServer(addr string, channels []*Channel) *Server {
	return &Server{
		server:   server.New(addr),
		channels: channels,
	}
}

// Start the server.
func (s *Server) Start() error {
	return s.server.Start()
}

// Stop the server.
func (s *Server) Stop() {
	s.server.Stop()
}

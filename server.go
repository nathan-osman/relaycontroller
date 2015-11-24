package main

import (
	"github.com/hectane/go-asyncserver"

	"net/http"
)

// Configuration for the server.
type ServerConfig struct {
	Addr string `json:"addr"`
	Root string `json:"root"`
}

// HTTP server providing the web interface and API.
type Server struct {
	server   *server.AsyncServer
	channels []*Channel
}

// Create a new server instance.
func NewServer(config *ServerConfig, channels []*Channel) *Server {
	s := &Server{
		server:   server.New(config.Addr),
		channels: channels,
	}
	m := http.NewServeMux()
	m.Handle("/", http.FileServer(http.Dir(config.Root)))
	s.server.Handler = m
	return s
}

// Start the server.
func (s *Server) Start() error {
	return s.server.Start()
}

// Stop the server.
func (s *Server) Stop() {
	s.server.Stop()
}

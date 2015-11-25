package main

import (
	"github.com/gorilla/mux"
	"github.com/hectane/go-asyncserver"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

// Configuration for the server.
type ServerConfig struct {
	Addr string `json:"addr"`
	Root string `json:"root"`
}

// HTTP server providing the web interface and API.
type Server struct {
	mutex    sync.Mutex
	server   *server.AsyncServer
	channels []*Channel
}

// Write a JSON response.
func (s *Server) writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// Return a list of all channels.
func (s *Server) channelsHandler(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.writeJSON(w, s.channels)
}

// Update a channel.
func (s *Server) channelHandler(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, channel := range s.channels {
		if channel.Config.Name == mux.Vars(r)["name"] {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			if err := channel.UnmarshalJSON(b); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			s.writeJSON(w, channel)
			return
		}
	}
	http.Error(w, "channel not found", http.StatusNotFound)
}

// Create a new server instance.
func NewServer(config *ServerConfig, channels []*Channel) *Server {
	s := &Server{
		server:   server.New(config.Addr),
		channels: channels,
	}
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir(config.Root)))
	r.HandleFunc("/api/channels", s.channelsHandler).Methods("GET")
	r.HandleFunc("/api/channels/{name}", s.channelHandler).Methods("PUT")
	s.server.Handler = r
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

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Ensure a single parameter was supplied - the filename of the config file
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s [CONFFILE]\n", os.Args[0])
		os.Exit(1)
	}

	// Attempt to load the configuration from the config file
	log.Print("Reading configuration...")
	config, err := LoadConfig(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Configuration read")

	// Set up all of the channels
	log.Print("Initializing channels...")
	channels := make([]*Channel, len(config.Channels))
	for i, config := range config.Channels {
		channels[i], err = NewChannel(&config)
		if err != nil {
			log.Fatal(err)
		}
		defer channels[i].Close()
	}
	log.Print("Channels initialized")

	// Create the server and bind to the specified port
	log.Print("Starting server...")
	server := NewServer(config.Addr, channels)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
	log.Print("Server started")

	// Wait for SIGINT
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c

	// Stop the server
	log.Print("Stopping server...")
	server.Stop()
	log.Print("Server stopped")
}

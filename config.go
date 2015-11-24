package main

import (
	"encoding/json"
	"os"
)

// Application configuration.
type Config struct {
	Addr     string          `json:"addr"`
	Channels []ChannelConfig `json:"channels"`
}

// Load configuration from the specified file.
func LoadConfig(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	config := &Config{}
	if err := json.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

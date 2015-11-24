package main

import (
	"github.com/nathan-osman/go-rpigpio"
)

// Configuration for a channel.
type ChannelConfig struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Number int    `json:"number"`
}

// Controller for a specific channel.
type Channel struct {
	config *ChannelConfig
	pin    *rpi.Pin
	state  bool
}

// Create a new controller for a specific channel. The pin is automatically set
// to LOW during initialization so that it has a known state.
func NewChannel(config *ChannelConfig) (*Channel, error) {
	p, err := rpi.OpenPin(config.Number, rpi.OUT)
	if err != nil {
		return nil, err
	}
	if err := p.Write(rpi.LOW); err != nil {
		return nil, err
	}
	return &Channel{
		config: config,
		pin:    p,
	}, nil
}

//...

func (c *Channel) Close() {
	c.pin.Close()
}

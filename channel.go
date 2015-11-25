package main

import (
	"github.com/nathan-osman/go-rpigpio"

	"encoding/json"
)

// Configuration for a channel.
type ChannelConfig struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Number int    `json:"number"`
}

// Controller for a specific channel.
type Channel struct {
	Config ChannelConfig
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
		Config: *config,
		pin:    p,
	}, nil
}

// Generate a JSON representation of the channel.
func (c *Channel) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"name":  c.Config.Name,
		"title": c.Config.Title,
		"state": c.state,
	}
	return json.Marshal(v)
}

// Apply the provided JSON to the channel. Only state may be set.
func (c *Channel) UnmarshalJSON(data []byte) error {
	var v struct {
		State bool `json:"state"`
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	var value rpi.Value
	if v.State {
		value = rpi.HIGH
	} else {
		value = rpi.LOW
	}
	if err := c.pin.Write(value); err != nil {
		return err
	}
	c.state = v.State
	return nil
}

// Close the channel.
func (c *Channel) Close() {
	c.pin.Close()
}

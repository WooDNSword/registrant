// Package config contains utilities for dealing with configuration data.
package config

import (
	"encoding/json"
	"io"
)

type Domain struct {
	name string
}

type Endpoint struct {
	address, port string
}

type Config struct {
	debug     bool
	domains   []Domain
	registrar Endpoint
}

func Eval(reader io.Reader) (Config, error) {
	var cfg Config
	
	dec := json.NewDecoder(reader)
	err := dec.Decode(&cfg)
	
	return cfg, err
}

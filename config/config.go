// Package config contains utilities for dealing with configuration data.
package config

import (
	"encoding/json"
	"io"
)

type Domain struct {
	Name string
}

type Endpoint struct {
	Address, Port string
}

type Config struct {
	Debug     bool
	Domains   []Domain
	Registrar Endpoint
}

// TODO: Move Eval and associated tests into their own repository to conform to
// DRY principle (it is identical in both registrar and registrant).

/*
Eval accepts a value of type io.Reader, tries to read all data from it, then
decodes the data from JSON to a struct `cfg` of type Config.

A tuple containing that struct value `cfg` as well as a value of type `err`,
denoting whether any issues occurred during decoding, is then returned.
*/
func Eval(reader io.Reader) (Config, error) {
	var cfg Config

	dec := json.NewDecoder(reader)
	err := dec.Decode(&cfg)
	// TODO: Panic on err != nil.

	return cfg, err
}

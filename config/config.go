<<<<<<< HEAD
package config

/* TODO: Uncomment this.
import (
	"encoding/json"
)
*/

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
=======
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
>>>>>>> config-dev
}

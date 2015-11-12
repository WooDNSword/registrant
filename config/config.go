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
}

package config

/* TODO: Uncomment this.
import (
	"encoding/json"
)
*/

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

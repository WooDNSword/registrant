// Package main implements the WooDNSword Registrant program.
package main

import (
	"github.com/WooDNSword/registrant/config"
	"github.com/WooDNSword/registrant/connection"
	"github.com/WooDNSword/registrant/session"
	"os"
)

func main() {
	cfgFile, err := os.Open("res/json/cfg.json")
	if err != nil {
		panic(err)
	}

	cfg, err := config.Eval(cfgFile)
	if err != nil {
		panic(err)
	}

	registrar := cfg.Registrar

	connection.Initiate(registrar, session.Handler)
}

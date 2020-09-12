package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Represent
type Config struct {
	Server   string
	Database string
}

// leer achico config
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}

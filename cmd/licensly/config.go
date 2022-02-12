package main

import (
	"go.sancus.dev/config"
	"go.sancus.dev/config/yaml"
)

type Config struct{}

func (cfg *Config) ReadInFile(filename string) error {
	return yaml.LoadFile(filename, cfg)
}

func (cfg *Config) Prepare() error {
	return config.Prepare(cfg)
}

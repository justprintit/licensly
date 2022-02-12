package main

import (
	"io"

	"go.sancus.dev/config/yaml"

	"github.com/justprintit/licensly/config"
)

type Config struct{}

func (cfg *Config) ReadInFile(filename string) error {
	return yaml.LoadFile(filename, cfg)
}

func (cfg *Config) Prepare() error {
	return config.Prepare(cfg)
}

func (cfg *Config) WriteTo(w io.Writer) error {
	_, err := yaml.WriteTo(w, cfg)
	return err
}

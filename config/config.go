package config

import (
	"go.sancus.dev/config"
)

func Prepare(v interface{}) error {
	return config.Prepare(v)
}

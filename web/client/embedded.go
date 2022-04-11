package client

//go:generate go run go.sancus.dev/web/cmd/embed -o out.go -p client -n embedded out

import (
	"io/fs"
)

func Embedded() (fs.FS, error) {
	return fs.Sub(&embedded, "out")
}

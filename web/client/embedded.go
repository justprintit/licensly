package client

//go:generate go run go.sancus.dev/web/cmd/embed -o out.go -p client -n embedded out

import (
	"io/fs"
	"log"

	"go.sancus.dev/web"
	"go.sancus.dev/web/embed"
)

func Embedded() (fs.FS, error) {
	return fs.Sub(&embedded, "out")
}

func Middleware() web.MiddlewareHandlerFunc {
	fsys, err := Embedded()
	if err != nil {
		log.Fatal(err)
	}

	m := embed.EmbedFS{
		FS:    fsys,
		Index: "index.html",
	}

	return m.Middleware
}

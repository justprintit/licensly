package client

import (
	"embed"
	"io/fs"
)

//go:embed out/*
var embedded embed.FS

func GetFilesystem() fs.FS {
	fsys, err := fs.Sub(embedded, "out")
	if err != nil {
		panic(err)
	}
	return fsys
}

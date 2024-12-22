package static

import (
	"embed"
	"io/fs"
)

//go:embed ui
var ui embed.FS

func UI() fs.FS {
	subFS, err := fs.Sub(ui, "ui")
	if err != nil {
		panic(err)
	}

	return subFS
}

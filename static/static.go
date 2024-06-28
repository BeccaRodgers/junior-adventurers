package static

import (
	"embed"
	"net/http"
)

//go:embed static
var staticFiles embed.FS

func Handler() http.Handler {
	return http.FileServerFS(staticFiles)
}

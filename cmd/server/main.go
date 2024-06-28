package main

import (
	"embed"
	"errors"
	"github.com/a-h/templ"
	"junior-adventurers/views"
	"log"
	"net/http"
)

//go:embed static
var staticFiles embed.FS

func main() {
	handler := http.NewServeMux()
	handler.Handle("GET /static/*", http.FileServerFS(staticFiles))
	handler.Handle("GET /", templ.Handler(views.Homepage()))
	handler.Handle("GET /guilds/1", templ.Handler(views.Guild()))
	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

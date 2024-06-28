package main

import (
	"embed"
	"errors"
	"log"
	"net/http"
)

//go:embed *.html *.css
var publicFiles embed.FS

func main() {
	handler := http.FileServerFS(publicFiles)
	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

package main

import (
	"errors"
	"junior-adventurers/controllers"
	"log"
	"net/http"
)

func main() {

	server := http.Server{
		Addr:    ":8080",
		Handler: controllers.Handler(),
	}

	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"errors"
	"junior-adventurers/controller"
	"junior-adventurers/fixtures"
	"junior-adventurers/memdb"
	"log"
	"net/http"
)

func main() {
	guilds := memdb.NewGuildRepository()
	if err := fixtures.InsertGuilds(context.Background(), guilds); err != nil {
		log.Print("Inserting fake data:", err)
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: controller.Handler(guilds),
	}

	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

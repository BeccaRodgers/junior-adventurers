package controllers

import (
	"github.com/a-h/templ"
	"junior-adventurers/static"
	"junior-adventurers/views"
	"net/http"
)

func Handler() http.Handler {
	handler := http.NewServeMux()
	handler.Handle("GET /static/*", static.Handler())
	handler.Handle("GET /", templ.Handler(views.Homepage()))
	handler.Handle("GET /guilds/1", templ.Handler(views.Guild()))
	return handler
}

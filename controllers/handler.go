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
	handler.Handle("GET /guilds/1", templ.Handler(views.Guild(guild)))
	return handler
}

var guild = views.GuildData{
	Name:        "Founders Guild",
	GuildMaster: "Beyoncé",
	Leaders: []views.Leader{
		{
			Member: views.Member{
				Name:    "Alex",
				Age:     40,
				Species: "Human",
			},
		},
		{
			Member: views.Member{
				Name:    "Beyoncé",
				Age:     42,
				Species: "Dwarf",
			},
		},
	},
	Members: []views.Member{
		{
			Name:    "Angela",
			Age:     10,
			Species: "Human",
		},
		{
			Name:    "Bob",
			Age:     12,
			Species: "Human",
		},
	},
}

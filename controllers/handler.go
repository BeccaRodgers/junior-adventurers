package controllers

import (
	"github.com/a-h/templ"
	"junior-adventurers/static"
	"junior-adventurers/views"
	"net/http"
	"time"
)

func Handler() http.Handler {
	handler := http.NewServeMux()
	handler.Handle("GET /static/*", static.Handler())
	handler.Handle("GET /", templ.Handler(views.Homepage()))
	handler.Handle("GET /guilds/1", templ.Handler(views.Guild(guild)))
	return handler
}

var guild = views.GuildData{
	Name:         "Founders Guild",
	Founded:      time.Now(),
	MeetingPlace: "Hall of the Mountain King",
	MeetingTime:  "Thursday, 7pm",
	Email:        "foundersguild@example.com",
	GuildMaster:  "Beyoncé",
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
		{
			Name:    "Charlotte",
			Age:     11,
			Species: "Werewolf",
		},
		{
			Name:    "David",
			Age:     11,
			Species: "Dwarf",
		},
	},
}

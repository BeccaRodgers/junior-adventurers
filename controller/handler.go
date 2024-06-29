package controller

import (
	"github.com/a-h/templ"
	"junior-adventurers/static"
	"junior-adventurers/view"
	"net/http"
	"time"
)

func Handler() http.Handler {
	handler := http.NewServeMux()
	handler.Handle("GET /static/*", static.Handler())
	handler.Handle("GET /", templ.Handler(view.Homepage()))
	handler.Handle("GET /guilds/1", templ.Handler(view.Guild(guild)))
	return handler
}

var guild = view.GuildData{
	Name:         "Founders Guild",
	Founded:      time.Now(),
	MeetingPlace: "Hall of the Mountain King",
	MeetingTime:  "Thursday, 7pm",
	Email:        "foundersguild@example.com",
	GuildMaster:  "Beyoncé",
	Leaders: []view.Leader{
		{
			Member: view.Member{
				Name:    "Alex",
				Age:     40,
				Species: "Human",
			},
		},
		{
			Member: view.Member{
				Name:    "Beyoncé",
				Age:     42,
				Species: "Dwarf",
			},
		},
	},
	Members: []view.Member{
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

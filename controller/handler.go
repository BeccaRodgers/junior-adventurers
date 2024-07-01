package controller

import (
	"artk.dev/apperror"
	"artk.dev/httperror"
	"github.com/a-h/templ"
	"junior-adventurers/model"
	"junior-adventurers/static"
	"junior-adventurers/view"
	"net/http"
	"strconv"
	"time"
)

func Handler(guilds model.GuildRepository) http.Handler {
	controller := controller{
		ServeMux: http.NewServeMux(),
		guilds:   guilds,
	}
	controller.Handle("GET /static/*", static.Handler())
	controller.Handle("GET /", templ.Handler(view.Homepage()))
	controller.HandleFunc("GET /guilds/{guildID}", controller.getGuild)
	return controller
}

type controller struct {
	*http.ServeMux
	guilds model.GuildRepository
}

func (c controller) getGuild(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idString := r.PathValue("guildID")
	id, err := strconv.Atoi(idString)
	if err != nil {
		httperror.EncodeToText(w, apperror.Validationf("invalid guild ID: %v", idString))
		return
	}

	guild, err := c.guilds.Get(ctx, model.GuildID(id))
	if err != nil {
		httperror.EncodeToText(w, err)
		return
	}

	viewModel := exampleGuildViewModel()
	viewModel.Name = string(guild.Name())

	_ = view.Guild(viewModel).Render(r.Context(), w)
}

// Deprecated: use model instead.
func exampleGuildViewModel() view.GuildData {
	return view.GuildData{
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
}

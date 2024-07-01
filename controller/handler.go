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

	viewModel := c.assembleGuildData(guild)
	_ = view.Guild(viewModel).Render(r.Context(), w)
}

func (c controller) assembleGuildData(guild *model.Guild) view.GuildData {
	return view.GuildData{
		Name:         string(guild.Name()),
		Founded:      guild.FoundingDate(),
		MeetingPlace: string(guild.MeetingPlace()),
		MeetingTime:  string(guild.MeetingTime()),
		Email:        string(guild.Email()),
		GuildMaster:  "Beyoncé",
		Leaders:      c.assembleLeaders(),
		Members:      c.assembleMembers(),
	}
}

func (c controller) assembleMembers() []view.Member {
	return []view.Member{
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
	}
}

func (c controller) assembleLeaders() []view.Leader {
	return []view.Leader{
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
	}
}

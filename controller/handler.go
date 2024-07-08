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

func Handler(guilds model.GuildRepository, members model.MemberRepository) http.Handler {
	controller := controller{
		ServeMux: http.NewServeMux(),
		guilds:   guilds,
		members:  members,
	}
	controller.Handle("GET /static/*", static.Handler())
	controller.Handle("GET /", templ.Handler(view.Homepage()))
	controller.Handle("GET /members", templ.Handler(view.NewMember(model.MemberSpeciesValues())))
	controller.HandleFunc("GET /guilds/{guildID}", controller.getGuild)
	return controller
}

type controller struct {
	*http.ServeMux
	guilds  model.GuildRepository
	members model.MemberRepository
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

	var members []*model.Member
	for _, memberID := range guild.Members() {
		member, err := c.members.Get(ctx, memberID)
		if err != nil {
			httperror.EncodeToText(w, err)
			return
		}
		members = append(members, member)
	}

	var leaders []*model.Member
	var guildMaster *model.Member
	for _, memberID := range guild.Leaders() {
		member, err := c.members.Get(ctx, memberID)
		if err != nil {
			httperror.EncodeToText(w, err)
			return
		}
		leaders = append(leaders, member)
		if memberID == guild.GuildMaster() {
			guildMaster = member
		}
	}

	if guildMaster == nil {
		httperror.EncodeToText(w, apperror.NotFound("guildMaster not found in members"))
		return
	}

	var enquiries []*model.Member
	for _, memberID := range guild.Enquiries() {
		member, err := c.members.Get(ctx, memberID)
		if err != nil {
			httperror.EncodeToText(w, err)
			return
		}
		enquiries = append(enquiries, member)
	}

	var waitingList []*model.Member
	for _, memberID := range guild.WaitingList() {
		member, err := c.members.Get(ctx, memberID)
		if err != nil {
			httperror.EncodeToText(w, err)
			return
		}
		waitingList = append(waitingList, member)
	}

	viewModel := c.assembleGuildData(guild, members, leaders, enquiries, waitingList, guildMaster)
	_ = view.Guild(viewModel).Render(r.Context(), w)
}

func (c controller) assembleGuildData(guild *model.Guild, members, leaders, enquiries, waitingList []*model.Member, guildMaster *model.Member) view.GuildData {
	return view.GuildData{
		Name:         string(guild.Name()),
		Capacity:     int(guild.Capacity()),
		Founded:      guild.FoundingDate(),
		MeetingPlace: string(guild.MeetingPlace()),
		MeetingTime:  string(guild.MeetingTime()),
		Email:        string(guild.Email()),
		GuildMaster: view.GuildMaster{
			Name:  string(guildMaster.Name()),
			Image: string(guildMaster.Image()),
		},
		Leaders:     c.assembleLeaders(leaders),
		Members:     c.assembleMembers(members),
		Enquiries:   c.assembleMembers(enquiries),
		WaitingList: c.assembleMembers(waitingList),
	}
}

func (c controller) assembleMembers(members []*model.Member) []view.Member {
	var viewMembers []view.Member
	for _, member := range members {
		viewMembers = append(viewMembers, view.Member{
			Name:    string(member.Name()),
			Age:     member.Age(),
			Species: member.Species().String(),
		})
	}
	return viewMembers
}

func (c controller) assembleLeaders(leaders []*model.Member) []view.Leader {
	var viewLeaders []view.Leader
	for _, leader := range leaders {
		viewLeaders = append(viewLeaders, view.Leader{
			Member: view.Member{
				Name:    string(leader.Name()),
				Age:     leader.Age(),
				Species: leader.Species().String(),
			},
		})
	}
	return viewLeaders
}

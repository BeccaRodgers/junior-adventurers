package controller

import (
	"artk.dev/apperror"
	"artk.dev/httperror"
	"cmp"
	"github.com/a-h/templ"
	"junior-adventurers/model"
	"junior-adventurers/static"
	"junior-adventurers/view"
	"net/http"
	"slices"
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
	controller.HandleFunc("GET /guilds/{guildID}/enquiries", controller.getGuildEnquiries)
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

	var waitingList int
	for _, status := range guild.Enquiries() {
		if status == model.WaitingList {
			waitingList++
		}
	}

	viewModel := c.assembleGuildData(guild, members, leaders, waitingList, guildMaster)
	_ = view.Guild(viewModel).Render(r.Context(), w)
}

func (c controller) getGuildEnquiries(w http.ResponseWriter, r *http.Request) {
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

	var enquiries []*model.Member
	var waitingList []*model.Member
	for memberID, status := range guild.Enquiries() {
		member, err := c.members.Get(ctx, memberID)
		if err != nil {
			httperror.EncodeToText(w, err)
			return
		}
		switch status {
		case model.Enquired:
			enquiries = append(enquiries, member)
		case model.WaitingList:
			waitingList = append(waitingList, member)
		}
	}

	viewModel := c.assembleGuildEnquiriesData(guild, enquiries, waitingList)
	_ = view.GuildEnquiries(viewModel).Render(r.Context(), w)
}

func (c controller) assembleGuildData(guild *model.Guild, members, leaders []*model.Member, waitingList int, guildMaster *model.Member) view.GuildData {
	return view.GuildData{
		ID:           int(guild.ID()),
		Name:         string(guild.Name()),
		Type:         guild.Type().String(),
		Capacity:     int(guild.Capacity()),
		Founded:      guild.FoundingDate(),
		MeetingPlace: string(guild.MeetingPlace()),
		MeetingTime:  string(guild.MeetingTime()),
		Email:        string(guild.Email()),
		GuildMaster: view.GuildMaster{
			Name:  string(guildMaster.Name()),
			Image: string(guildMaster.Image()),
		},
		Leaders:        c.assembleMembers(leaders),
		Members:        c.assembleMembers(members),
		WaitingListLen: waitingList,
	}
}

func (c controller) assembleGuildEnquiriesData(guild *model.Guild, enquiries, waitingList []*model.Member) view.GuildEnquiriesData {
	return view.GuildEnquiriesData{
		ID:          int(guild.ID()),
		Name:        string(guild.Name()),
		Type:        guild.Type().String(),
		Capacity:    int(guild.Capacity()),
		NumMembers:  len(guild.Members()),
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
	sortFunc := func(a, b view.Member) int {
		return cmp.Compare(b.Age, a.Age)
	}
	slices.SortStableFunc(viewMembers, sortFunc)
	return viewMembers
}

package controller

import (
	"artk.dev/apperror"
	"artk.dev/httperror"
	"cmp"
	"fmt"
	"github.com/a-h/templ"
	"junior-adventurers/model"
	"junior-adventurers/static"
	"junior-adventurers/view"
	"net/http"
	"slices"
	"strconv"
	"time"
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
	controller.HandleFunc("POST /members", controller.postMember)
	controller.HandleFunc("GET /members/{memberID}", controller.getMember)
	controller.HandleFunc("GET /guilds/{guildID}", controller.getGuild)
	controller.HandleFunc("GET /guilds/{guildID}/enquiries", controller.getGuildEnquiries)
	return controller
}

type controller struct {
	*http.ServeMux
	guilds  model.GuildRepository
	members model.MemberRepository
}

func (c controller) getMember(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idString := r.PathValue("memberID")
	id, err := strconv.Atoi(idString)
	if err != nil {
		httperror.EncodeToText(w, apperror.Validationf("invalid member ID: %v", idString))
		return
	}

	member, err := c.members.Get(ctx, model.MemberID(id))
	if err != nil {
		httperror.EncodeToText(w, err)
		return
	}

	viewModel := c.assembleMember(member)
	_ = view.Member(viewModel).Render(r.Context(), w)
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

func (c controller) postMember(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	//guildId, err := strconv.Atoi(r.Form.Get("guild"))
	//if err != nil {
	//	httperror.EncodeToText(w, err)
	//	return
	//}

	err := r.ParseForm()
	if err != nil {
		httperror.EncodeToText(w, err)
		return
	}

	dob, err := time.Parse(time.DateOnly, r.Form.Get("dob"))
	if err != nil {
		httperror.EncodeToText(w, err)
		return
	}
	memberForm := view.NewMemberForm{
		Name:    r.Form.Get("name"),
		DOB:     dob,
		Species: r.Form.Get("species"),
		//Guild:   guildId,
	}

	member, err := c.assembleNewMember(memberForm)
	if err != nil {
		httperror.EncodeToText(w, err)
		return
	}

	err = c.members.Insert(ctx, member)
	if err != nil {
		httperror.EncodeToText(w, err)
		return
	}

	// TODO add member to guild enquiries.

	w.Header().Add("HX-Push-Url", fmt.Sprintf("/members/%v", member.ID()))
	w.WriteHeader(http.StatusCreated)
	viewModel := c.assembleMember(member)
	_ = view.Member(viewModel).Render(r.Context(), w)
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

func (c controller) assembleMembers(members []*model.Member) []view.MemberData {
	var viewMembers []view.MemberData
	for _, member := range members {
		viewMembers = append(viewMembers, c.assembleMember(member))
	}
	sortFunc := func(a, b view.MemberData) int {
		return cmp.Compare(b.Age, a.Age)
	}
	slices.SortStableFunc(viewMembers, sortFunc)
	return viewMembers
}

func (c controller) assembleMember(member *model.Member) view.MemberData {
	return view.MemberData{
		ID:      int(member.ID()),
		Name:    string(member.Name()),
		Age:     member.Age(),
		Species: member.Species().String(),
	}
}

func (c controller) assembleNewMember(member view.NewMemberForm) (*model.Member, error) {
	return model.MemberSerialization{
		ID:      model.MemberID(11), // TODO use next available id
		Name:    model.MemberName(member.Name),
		DOB:     member.DOB,
		Species: model.UnmarshalMemberSpecies(member.Species),
	}.Deserialize(), nil
}

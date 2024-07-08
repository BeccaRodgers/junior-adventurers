package model

import "time"

type Guild struct {
	id           GuildID
	name         GuildName
	capacity     GuildCapacity
	foundingDate time.Time
	meetingPlace GuildMeetingPlace
	meetingTime  GuildMeetingTime
	email        GuildEmail
	guildMaster  MemberID
	members      []MemberID
	leaders      []MemberID
	enquiries    []MemberID
	waitingList  []MemberID
}

func (g *Guild) ID() GuildID {
	return g.id
}

func (g *Guild) Name() GuildName {
	return g.name
}

func (g *Guild) Capacity() GuildCapacity {
	return g.capacity
}

func (g *Guild) FoundingDate() time.Time {
	return g.foundingDate
}

func (g *Guild) MeetingPlace() GuildMeetingPlace {
	return g.meetingPlace
}

func (g *Guild) MeetingTime() GuildMeetingTime {
	return g.meetingTime
}

func (g *Guild) Email() GuildEmail {
	return g.email
}

func (g *Guild) GuildMaster() MemberID {
	return g.guildMaster
}

func (g *Guild) Members() []MemberID {
	return g.members
}

func (g *Guild) Leaders() []MemberID {
	return g.leaders
}

func (g *Guild) Enquiries() []MemberID {
	return g.enquiries
}

func (g *Guild) WaitingList() []MemberID {
	return g.waitingList
}

func (g *Guild) Serialize() GuildSerialization {
	return GuildSerialization{
		ID:           g.id,
		Name:         g.name,
		Capacity:     g.capacity,
		FoundingDate: g.foundingDate,
		MeetingPlace: g.meetingPlace,
		MeetingTime:  g.meetingTime,
		Email:        g.email,
		GuildMaster:  g.guildMaster,
		Members:      g.members,
		Leaders:      g.leaders,
		Enquiries:    g.enquiries,
		WaitingList:  g.waitingList,
	}
}

type GuildSerialization struct {
	ID           GuildID
	Name         GuildName
	Capacity     GuildCapacity
	FoundingDate time.Time
	MeetingPlace GuildMeetingPlace
	MeetingTime  GuildMeetingTime
	Email        GuildEmail
	GuildMaster  MemberID
	Members      []MemberID
	Leaders      []MemberID
	Enquiries    []MemberID
	WaitingList  []MemberID
}

func (s GuildSerialization) Deserialize() *Guild {
	return &Guild{
		id:           s.ID,
		name:         s.Name,
		capacity:     s.Capacity,
		foundingDate: s.FoundingDate,
		meetingPlace: s.MeetingPlace,
		meetingTime:  s.MeetingTime,
		email:        s.Email,
		guildMaster:  s.GuildMaster,
		members:      s.Members,
		leaders:      s.Leaders,
		enquiries:    s.Enquiries,
		waitingList:  s.WaitingList,
	}
}

type GuildID int64
type GuildName string
type GuildCapacity int
type GuildMeetingPlace string
type GuildMeetingTime string
type GuildEmail string

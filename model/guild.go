package model

import "time"

type Guild struct {
	id           GuildID
	name         GuildName
	foundingDate time.Time
	meetingPlace GuildMeetingPlace
	meetingTime  GuildMeetingTime
	email        GuildEmail
	members      []MemberID
	leaders      []MemberID
}

func (g *Guild) Name() GuildName {
	return g.name
}

func (g *Guild) ID() GuildID {
	return g.id
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

func (g *Guild) Members() []MemberID {
	return g.members
}

func (g *Guild) Leaders() []MemberID {
	return g.leaders
}

func (g *Guild) Serialize() GuildSerialization {
	return GuildSerialization{
		ID:           g.id,
		Name:         g.name,
		FoundingDate: g.foundingDate,
		MeetingPlace: g.meetingPlace,
		MeetingTime:  g.meetingTime,
		Email:        g.email,
		Members:      g.members,
		Leaders:      g.leaders,
	}
}

type GuildSerialization struct {
	ID           GuildID
	Name         GuildName
	FoundingDate time.Time
	MeetingPlace GuildMeetingPlace
	MeetingTime  GuildMeetingTime
	Email        GuildEmail
	Members      []MemberID
	Leaders      []MemberID
}

func (s GuildSerialization) Deserialize() *Guild {
	return &Guild{
		id:           s.ID,
		name:         s.Name,
		foundingDate: s.FoundingDate,
		meetingPlace: s.MeetingPlace,
		meetingTime:  s.MeetingTime,
		email:        s.Email,
		members:      s.Members,
		leaders:      s.Leaders,
	}
}

type GuildID int64
type GuildName string
type GuildMeetingPlace string
type GuildMeetingTime string
type GuildEmail string

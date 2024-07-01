package model

import "time"

type Guild struct {
	id           GuildID
	name         GuildName
	foundingDate time.Time
	meetingPlace GuildMeetingPlace
	meetingTime  GuildMeetingTime
	email        GuildEmail
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

func (g *Guild) Serialize() GuildSerialization {
	return GuildSerialization{
		ID:           g.id,
		Name:         g.name,
		FoundingDate: g.foundingDate,
		MeetingPlace: g.meetingPlace,
		MeetingTime:  g.meetingTime,
		Email:        g.email,
	}
}

type GuildSerialization struct {
	ID           GuildID
	Name         GuildName
	FoundingDate time.Time
	MeetingPlace GuildMeetingPlace
	MeetingTime  GuildMeetingTime
	Email        GuildEmail
}

func (s GuildSerialization) Deserialize() *Guild {
	return &Guild{
		id:           s.ID,
		name:         s.Name,
		foundingDate: s.FoundingDate,
		meetingPlace: s.MeetingPlace,
		meetingTime:  s.MeetingTime,
		email:        s.Email,
	}
}

type GuildID int64
type GuildName string
type GuildMeetingPlace string
type GuildMeetingTime string
type GuildEmail string

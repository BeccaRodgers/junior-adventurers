package views

import "time"

type GuildData struct {
	Name         string
	Founded      time.Time
	MeetingPlace string
	MeetingTime  string
	Email        string
	GuildMaster  string
	Leaders      []Leader
	Members      []Member
}

type Leader struct {
	Member
}

type Member struct {
	Name    string
	Age     int
	Species string
}

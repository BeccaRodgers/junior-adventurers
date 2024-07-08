package view

import "time"

type GuildData struct {
	Name         string
	Type         string
	Capacity     int
	Founded      time.Time
	MeetingPlace string
	MeetingTime  string
	Email        string
	GuildMaster  GuildMaster
	Leaders      []Leader
	Members      []Member
	Enquiries    []Member
	WaitingList  []Member
}

type Leader struct {
	Member
}

type Member struct {
	Name    string
	Age     int
	Species string
}

type GuildMaster struct {
	Name  string
	Image string
}

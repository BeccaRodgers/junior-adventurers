package view

import "time"

type GuildData struct {
	ID             int
	Name           string
	Type           string
	Capacity       int
	Founded        time.Time
	MeetingPlace   string
	MeetingTime    string
	Email          string
	GuildMaster    GuildMaster
	Leaders        []Leader
	Members        []Member
	WaitingListLen int
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

type GuildEnquiriesData struct {
	ID          int
	Name        string
	Type        string
	Capacity    int
	NumMembers  int
	Enquiries   []Member
	WaitingList []Member
}

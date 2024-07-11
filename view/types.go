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
	Leaders        []MemberData
	Members        []MemberData
	WaitingListLen int
}

type MemberData struct {
	ID      int
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
	Enquiries   []MemberData
	WaitingList []MemberData
}

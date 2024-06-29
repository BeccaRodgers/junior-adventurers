package views

type GuildData struct {
	Name        string
	GuildMaster string
	Leaders     []Leader
	Members     []Member
}

type Leader struct {
	Member
}

type Member struct {
	Name    string
	Age     int
	Species string
}

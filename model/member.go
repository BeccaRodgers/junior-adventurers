package model

import "time"

type Member struct {
	id      MemberID
	name    MemberName
	dob     time.Time
	species MemberSpecies
}

func (m *Member) ID() MemberID {
	return m.id
}

func (m *Member) Name() MemberName {
	return m.name
}

func (m *Member) Dob() time.Time {
	return m.dob
}

func (m *Member) Age() int {
	year, month, day := m.dob.Date()
	currentYear, currentMonth, currentDay := time.Now().Date()

	years := currentYear - year
	months := currentMonth - month
	days := currentDay - day
	if days < 0 {
		months--
	}
	if months < 0 {
		years--
	}
	return years
}

func (m *Member) Species() MemberSpecies {
	return m.species
}

func (m *Member) Serialize() MemberSerialization {
	return MemberSerialization{
		ID:      m.id,
		Name:    m.name,
		DOB:     m.dob,
		Species: m.species,
	}
}

type MemberSerialization struct {
	ID      MemberID
	Name    MemberName
	DOB     time.Time
	Species MemberSpecies
}

func (s MemberSerialization) Deserialize() *Member {
	return &Member{
		id:      s.ID,
		name:    s.Name,
		dob:     s.DOB,
		species: s.Species,
	}
}

type MemberID int64
type MemberName string
type MemberSpecies int64

const (
	Human MemberSpecies = iota
	Dwarf
	Werewolf
)

func (s MemberSpecies) String() string {
	switch s {
	case Human:
		return "Human"
	case Dwarf:
		return "Dwarf"
	case Werewolf:
		return "Werewolf"
	default:
		return "Unknown Species"
	}
}

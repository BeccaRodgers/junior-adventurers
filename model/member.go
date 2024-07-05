package model

import "time"

type Member struct {
	id      MemberID
	name    MemberName
	dob     time.Time
	species MemberSpecies
	image   MemberImage
}

func (m *Member) ID() MemberID {
	return m.id
}

func (m *Member) Name() MemberName {
	return m.name
}

func (m *Member) DOB() time.Time {
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

func (m *Member) Image() MemberImage {
	return m.image
}

func (m *Member) Serialize() MemberSerialization {
	return MemberSerialization{
		ID:      m.id,
		Name:    m.name,
		DOB:     m.dob,
		Species: m.species,
		Image:   m.image,
	}
}

type MemberSerialization struct {
	ID      MemberID
	Name    MemberName
	DOB     time.Time
	Species MemberSpecies
	Image   MemberImage
}

func (s MemberSerialization) Deserialize() *Member {
	return &Member{
		id:      s.ID,
		name:    s.Name,
		dob:     s.DOB,
		species: s.Species,
		image:   s.Image,
	}
}

type MemberID int64
type MemberName string
type MemberSpecies int64
type MemberImage string

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

func MemberSpeciesValues() []MemberSpecies {
	return []MemberSpecies{
		Human,
		Dwarf,
		Werewolf,
	}
}

package model

import (
	"fmt"
	"time"
)

type Member struct {
	id      MemberID
	name    MemberName
	dob     time.Time
	species SpeciesID
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

func (m *Member) Species() SpeciesID {
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
	Species SpeciesID
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

func NewMember(id MemberID, name MemberName, dob time.Time, speciesID SpeciesID) (*Member, error) {
	// TODO validate inputs
	name, err := name.Normalize()
	if err != nil {
		return nil, fmt.Errorf("invalid member name: %w", err)
	}

	// TODO enforce rules.

	return &Member{
		id:      id,
		name:    name,
		dob:     dob,
		species: speciesID,
		image:   "",
	}, nil
}

type MemberID int64
type SpeciesID int64
type MemberImage string

const (
	Unknown SpeciesID = iota
	Human
	Dwarf
	Werewolf
)

func (s SpeciesID) String() string {
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

func UnmarshalMemberSpecies(species string) SpeciesID {
	switch species {
	case "Human":
		return Human
	case "Dwarf":
		return Dwarf
	case "Werewolf":
		return Werewolf
	default:
		return Unknown
	}
}

func MemberSpeciesValues() []SpeciesID {
	return []SpeciesID{
		Human,
		Dwarf,
		Werewolf,
	}
}

package model

type SpeciesID int64

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

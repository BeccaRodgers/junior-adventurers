package model

type Guild struct {
	id   GuildID
	name GuildName
}

func (g *Guild) Name() GuildName {
	return g.name
}

func (g *Guild) ID() GuildID {
	return g.id
}

func (g *Guild) Serialize() GuildSerialization {
	return GuildSerialization{
		ID:   g.id,
		Name: g.name,
	}
}

type GuildSerialization struct {
	ID   GuildID
	Name GuildName
}

func (s GuildSerialization) Deserialize() *Guild {
	return &Guild{
		id:   s.ID,
		name: s.Name,
	}
}

type GuildID int64
type GuildName string

package model

import "artk.dev/crud"

type GuildRepository interface {
	crud.Repository[*Guild, GuildID, GuildSerialization]
}

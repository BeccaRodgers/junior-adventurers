package memdb

import (
	"artk.dev/crud"
	"fmt"
	"junior-adventurers/model"
)

var _ model.GuildRepository = &GuildRepository{}

type GuildRepository struct {
	crud.InMemoryRepository[*model.Guild, model.GuildID, model.GuildSerialization]
}

func NewGuildRepository() *GuildRepository {
	r := &GuildRepository{}
	r.Errors.NotFound = func(id model.GuildID) string {
		return fmt.Sprintf("guild not found with ID %v", id)
	}

	r.Reset()
	return r
}

package fixtures

import (
	"context"
	"errors"
	"junior-adventurers/model"
)

func FoundersGuildID() model.GuildID {
	return 1
}

func FoundersGuild() *model.Guild {
	return model.GuildSerialization{
		ID:   FoundersGuildID(),
		Name: "Founder's Guild",
	}.Deserialize()
}

func InsertGuilds(ctx context.Context, guilds model.GuildRepository) error {
	var errs []error
	for _, guild := range []*model.Guild{
		FoundersGuild(),
	} {
		if err := guilds.Insert(ctx, guild); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

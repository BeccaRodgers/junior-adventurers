package fixtures

import (
	"context"
	"errors"
	"junior-adventurers/model"
	"time"
)

func FoundersGuildID() model.GuildID {
	return 1
}
func FoundersGuildName() model.GuildName {
	return "Founder's Guild"
}

func FoundersGuildFoundingDate() time.Time {
	return time.Date(2000, time.April, 20, 0, 0, 0, 0, time.UTC)
}

func FoundersGuildSerialization() model.GuildSerialization {
	return model.GuildSerialization{
		ID:           FoundersGuildID(),
		Name:         FoundersGuildName(),
		FoundingDate: FoundersGuildFoundingDate(),
		MeetingPlace: "Hall of the Mountain King",
		MeetingTime:  "Thursday, 7pm",
		Email:        "foundersguild@example.com",
		Members: []model.MemberID{
			AngelaID(),
			BobID(),
			CharlotteID(),
			DavidID(),
		},
		Leaders: []model.MemberID{
			AlexID(),
			BeyonceID(),
		},
	}
}

func FoundersGuild() *model.Guild {
	return FoundersGuildSerialization().Deserialize()
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

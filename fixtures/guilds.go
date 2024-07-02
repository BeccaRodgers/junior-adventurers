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

func FoundersGuild() *model.Guild {
	return model.GuildSerialization{
		ID:           FoundersGuildID(),
		Name:         "Founder's Guild",
		FoundingDate: time.Date(2000, time.April, 20, 0, 0, 0, 0, time.UTC),
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

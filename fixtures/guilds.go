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

func FoundersGuildCapacity() model.GuildCapacity {
	return 20
}

func FoundersGuildFoundingDate() time.Time {
	return time.Date(2000, time.April, 20, 0, 0, 0, 0, time.UTC)
}

func FoundersGuildMeetingPlace() model.GuildMeetingPlace {
	return "Hall of the Mountain King"
}

func FoundersGuildMeetingTime() model.GuildMeetingTime {
	return "Thursday, 19:00-20:30"
}

func FoundersGuildEmail() model.GuildEmail {
	return "foundersguild@example.com"
}

func FoundersGuildGuildMaster() model.MemberID {
	return BeyonceID()
}

func FoundersGuildMembers() []model.MemberID {
	return []model.MemberID{
		AngelaID(),
		BobID(),
		CharlotteID(),
		DavidID(),
	}
}

func FoundersGuildLeaders() []model.MemberID {
	return []model.MemberID{
		AlexID(),
		BeyonceID(),
	}
}

func FoundersGuildEnquiries() []model.MemberID {
	return []model.MemberID{
		CarlosID(),
		ErikaID(),
	}
}

func FoundersGuildWaitingList() []model.MemberID {
	return []model.MemberID{
		FredID(),
	}
}

func FoundersGuildSerialization() model.GuildSerialization {
	return model.GuildSerialization{
		ID:           FoundersGuildID(),
		Name:         FoundersGuildName(),
		Capacity:     FoundersGuildCapacity(),
		FoundingDate: FoundersGuildFoundingDate(),
		MeetingPlace: FoundersGuildMeetingPlace(),
		MeetingTime:  FoundersGuildMeetingTime(),
		Email:        FoundersGuildEmail(),
		GuildMaster:  FoundersGuildGuildMaster(),
		Members:      FoundersGuildMembers(),
		Leaders:      FoundersGuildLeaders(),
		Enquiries:    FoundersGuildEnquiries(),
		WaitingList:  FoundersGuildWaitingList(),
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

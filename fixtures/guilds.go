package fixtures

import (
	"context"
	"errors"
	"junior-adventurers/model"
	"time"
)

func FledglingFoundersGuildID() model.GuildID {
	return 1
}

func FalconFoundersGuildID() model.GuildID {
	return 2
}

func FoundersGuildName() model.GuildName {
	return "Founder's Guild"
}

func FledglingFoundersGuildType() model.GuildType {
	return model.Fledgling
}

func FalconFoundersGuildType() model.GuildType {
	return model.Falcon
}

func FledglingFoundersGuildCapacity() model.GuildCapacity {
	return 20
}

func FalconFoundersGuildCapacity() model.GuildCapacity {
	return 30
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

func FledglingFoundersGuildGuildMaster() model.MemberID {
	return BeyonceID()
}

func FalconFoundersGuildGuildMaster() model.MemberID {
	return AlexID()
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

func FoundersGuildEnquiries() model.GuildEnquiries {
	return map[model.MemberID]model.EnquiryStatus{
		CarlosID(): model.Enquired,
		ErikaID():  model.Enquired,
		FredID():   model.WaitingList,
	}
}

func FledglingFoundersGuildSerialization() model.GuildSerialization {
	return model.GuildSerialization{
		ID:           FledglingFoundersGuildID(),
		Name:         FoundersGuildName(),
		GuildType:    FledglingFoundersGuildType(),
		Capacity:     FledglingFoundersGuildCapacity(),
		FoundingDate: FoundersGuildFoundingDate(),
		MeetingPlace: FoundersGuildMeetingPlace(),
		MeetingTime:  FoundersGuildMeetingTime(),
		Email:        FoundersGuildEmail(),
		GuildMaster:  FledglingFoundersGuildGuildMaster(),
		Members:      FoundersGuildMembers(),
		Leaders:      FoundersGuildLeaders(),
		Enquiries:    FoundersGuildEnquiries(),
	}
}

func FalconFoundersGuildSerialization() model.GuildSerialization {
	return model.GuildSerialization{
		ID:           FalconFoundersGuildID(),
		Name:         FoundersGuildName(),
		GuildType:    FalconFoundersGuildType(),
		Capacity:     FalconFoundersGuildCapacity(),
		FoundingDate: FoundersGuildFoundingDate(),
		MeetingPlace: FoundersGuildMeetingPlace(),
		MeetingTime:  FoundersGuildMeetingTime(),
		Email:        FoundersGuildEmail(),
		GuildMaster:  FalconFoundersGuildGuildMaster(),
		Members:      FoundersGuildMembers(),
		Leaders:      FoundersGuildLeaders(),
		Enquiries:    FoundersGuildEnquiries(),
	}
}

func FledglingFoundersGuild() *model.Guild {
	return FledglingFoundersGuildSerialization().Deserialize()
}

func FalconFoundersGuild() *model.Guild {
	return FalconFoundersGuildSerialization().Deserialize()
}

func InsertGuilds(ctx context.Context, guilds model.GuildRepository) error {
	var errs []error
	for _, guild := range []*model.Guild{
		FledglingFoundersGuild(),
		FalconFoundersGuild(),
	} {
		if err := guilds.Insert(ctx, guild); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

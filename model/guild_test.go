package model_test

import (
	"github.com/stretchr/testify/assert"
	"junior-adventurers/fixtures"
	"testing"
)

func Test_Guild_DeserializeSerialize(t *testing.T) {
	guildSerialization := fixtures.FoundersGuildSerialization()

	guild := guildSerialization.Deserialize()
	guildSerialized := guild.Serialize()

	assert.Equal(t, guildSerialization, guildSerialized)
}

func Test_Guild_ID(t *testing.T) {
	guild := fixtures.FoundersGuild()

	id := guild.ID()

	assert.Equal(t, fixtures.FoundersGuildID(), id)
}

func Test_Guild_Name(t *testing.T) {
	guild := fixtures.FoundersGuild()

	name := guild.Name()

	assert.Equal(t, fixtures.FoundersGuildName(), name)
}

func Test_Guild_GuildType(t *testing.T) {
	guild := fixtures.FoundersGuild()

	guildType := guild.Type()

	assert.Equal(t, fixtures.FoundersGuildType(), guildType)
}

func Test_Guild_Capacity(t *testing.T) {
	guild := fixtures.FoundersGuild()

	capacity := guild.Capacity()

	assert.Equal(t, fixtures.FoundersGuildCapacity(), capacity)
}

func Test_Guild_FoundingDate(t *testing.T) {
	guild := fixtures.FoundersGuild()

	date := guild.FoundingDate()

	assert.Equal(t, fixtures.FoundersGuildFoundingDate(), date)
}

func Test_Guild_MeetingPlace(t *testing.T) {
	guild := fixtures.FoundersGuild()

	place := guild.MeetingPlace()

	assert.Equal(t, fixtures.FoundersGuildMeetingPlace(), place)
}

func Test_Guild_MeetingTime(t *testing.T) {
	guild := fixtures.FoundersGuild()

	time := guild.MeetingTime()

	assert.Equal(t, fixtures.FoundersGuildMeetingTime(), time)
}

func Test_Guild_Email(t *testing.T) {
	guild := fixtures.FoundersGuild()

	email := guild.Email()

	assert.Equal(t, fixtures.FoundersGuildEmail(), email)
}

func Test_Guild_Members(t *testing.T) {
	guild := fixtures.FoundersGuild()

	members := guild.Members()

	assert.Equal(t, fixtures.FoundersGuildMembers(), members)
}

func Test_Guild_Leaders(t *testing.T) {
	guild := fixtures.FoundersGuild()

	leaders := guild.Leaders()

	assert.Equal(t, fixtures.FoundersGuildLeaders(), leaders)
}

func Test_Guild_GuildMaster(t *testing.T) {
	guild := fixtures.FoundersGuild()

	guildMaster := guild.GuildMaster()

	assert.Equal(t, fixtures.FoundersGuildGuildMaster(), guildMaster)
}

func Test_Guild_Enquiries(t *testing.T) {
	guild := fixtures.FoundersGuild()

	enquiries := guild.Enquiries()

	assert.Equal(t, fixtures.FoundersGuildEnquiries(), enquiries)
}

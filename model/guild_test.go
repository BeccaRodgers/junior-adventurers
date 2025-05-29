package model_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"junior-adventurers/fixtures"
	"junior-adventurers/model"
	"testing"
)

func Test_Guild_DeserializeSerialize(t *testing.T) {
	guildSerialization := fixtures.FledglingFoundersGuildSerialization()

	guild := guildSerialization.Deserialize()
	guildSerialized := guild.Serialize()

	assert.Equal(t, guildSerialization, guildSerialized)
}

func Test_Guild_ID(t *testing.T) {
	guild := fixtures.FledglingFoundersGuild()

	id := guild.ID()

	assert.Equal(t, fixtures.FledglingFoundersGuildID(), id)
}

func Test_Guild_Name(t *testing.T) {
	guild := fixtures.FledglingFoundersGuild()

	name := guild.Name()

	assert.Equal(t, fixtures.FoundersGuildName(), name)
}

func Test_Guild_GuildType(t *testing.T) {
	guild := fixtures.FledglingFoundersGuild()

	guildType := guild.Type()

	assert.Equal(t, fixtures.FledglingFoundersGuildType(), guildType)
}

func Test_Guild_Capacity(t *testing.T) {
	guild := fixtures.FledglingFoundersGuild()

	capacity := guild.Capacity()

	assert.Equal(t, fixtures.FledglingFoundersGuildCapacity(), capacity)
}

func Test_Guild_FoundingDate(t *testing.T) {
	guild := fixtures.FledglingFoundersGuild()

	date := guild.FoundingDate()

	assert.Equal(t, fixtures.FoundersGuildFoundingDate(), date)
}

func Test_Guild_MeetingPlace(t *testing.T) {
	guild := fixtures.FledglingFoundersGuild()

	place := guild.MeetingPlace()

	assert.Equal(t, fixtures.FoundersGuildMeetingPlace(), place)
}

func Test_Guild_MeetingTime(t *testing.T) {
	guild := fixtures.FledglingFoundersGuild()

	time := guild.MeetingTime()

	assert.Equal(t, fixtures.FoundersGuildMeetingTime(), time)
}

func Test_Guild_Email(t *testing.T) {
	guild := fixtures.FledglingFoundersGuild()

	email := guild.Email()

	assert.Equal(t, fixtures.FoundersGuildEmail(), email)
}

func Test_Guild_Members(t *testing.T) {
	guild := fixtures.FledglingFoundersGuild()

	members := guild.Members()

	assert.Equal(t, fixtures.FoundersGuildMembers(), members)
}

func Test_Guild_Leaders(t *testing.T) {
	guild := fixtures.FledglingFoundersGuild()

	leaders := guild.Leaders()

	assert.Equal(t, fixtures.FoundersGuildLeaders(), leaders)
}

func Test_Guild_GuildMaster(t *testing.T) {
	guild := fixtures.FledglingFoundersGuild()

	guildMaster := guild.GuildMaster()

	assert.Equal(t, fixtures.FledglingFoundersGuildGuildMaster(), guildMaster)
}

func Test_Guild_Enquiries(t *testing.T) {
	guild := fixtures.FledglingFoundersGuild()

	enquiries := guild.Enquiries()

	assert.Equal(t, fixtures.FoundersGuildEnquiries(), enquiries)
}

func TestGuild_AddToEnquiries(t *testing.T) {
	guild := fixtures.FledglingFoundersGuild()
	enquiries := guild.Enquiries()
	_, ok := enquiries[fixtures.GeorgeID()]
	require.False(t, ok)

	guild.AddToEnquiries(fixtures.GeorgeID())

	enquiries = guild.Enquiries()
	status, ok := enquiries[fixtures.GeorgeID()]
	assert.True(t, ok)
	assert.Equal(t, model.Enquired, status)

}

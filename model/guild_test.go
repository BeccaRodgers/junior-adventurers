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

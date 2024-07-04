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

func Test_Guild_Name(t *testing.T) {
	guild := fixtures.FoundersGuild()
	
	name := guild.Name()

	assert.Equal(t, fixtures.FoundersGuildName(), name)
}
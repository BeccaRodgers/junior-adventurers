package model_test

import (
	"github.com/stretchr/testify/assert"
	"junior-adventurers/fixtures"
	"testing"
)

func Test_GuildDeserializeSerialize(t *testing.T) {
	guildSerialization := fixtures.FoundersGuildSerialization()

	guild := guildSerialization.Deserialize()
	guildSerialized := guild.Serialize()

	assert.Equal(t, guildSerialization, guildSerialized)
}

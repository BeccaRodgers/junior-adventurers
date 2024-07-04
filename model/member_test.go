package model_test

import (
	"github.com/stretchr/testify/assert"
	"junior-adventurers/fixtures"
	"testing"
)

func Test_MemberDeserializeSerialize(t *testing.T) {
	memberSerialization := fixtures.AngelaSerialization()

	member := memberSerialization.Deserialize()
	memberSerialized := member.Serialize()

	assert.Equal(t, memberSerialization, memberSerialized)
}

func Test_Member_ID(t *testing.T) {
	member := fixtures.Angela()

	id := member.ID()

	assert.Equal(t, fixtures.AngelaID(), id)
}

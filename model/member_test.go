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

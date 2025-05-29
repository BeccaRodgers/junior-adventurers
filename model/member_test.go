package model_test

import (
	"github.com/stretchr/testify/assert"
	"junior-adventurers/fixtures"
	"junior-adventurers/model"
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

func Test_Member_Name(t *testing.T) {
	member := fixtures.Angela()

	name := member.Name()

	assert.Equal(t, fixtures.AngelaName(), name)
}

func Test_Member_DOB(t *testing.T) {
	member := fixtures.Angela()

	dob := member.DOB()

	assert.Equal(t, fixtures.AngelaDOB(), dob)
}

func Test_Member_Species(t *testing.T) {
	member := fixtures.Angela()

	species := member.Species()

	assert.Equal(t, fixtures.AngelaSpecies(), species)
}

func Test_Member_Image(t *testing.T) {
	member := fixtures.Beyonce()

	image := member.Image()

	assert.Equal(t, fixtures.BeyonceImage(), image)
}

func Test_Member_NextAvailableID(t *testing.T) {
	// Fixtures will have already taken IDs, so start after these.
	// Calling the function "uses" that ID, so we need to start from one after.
	start := model.NextAvailableID() + 1
	for i := start; i < start+10; i++ {
		assert.Equal(t, i, model.NextAvailableID())
	}
}

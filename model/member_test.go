package model_test

import (
	"github.com/stretchr/testify/assert"
	"junior-adventurers/fixtures"
	"junior-adventurers/model"
	"testing"
	"time"
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

func TestMember_NewMember(t *testing.T) {
	for _, tc := range []struct {
		name       string
		memberName model.MemberName
		dob        time.Time
		speciesID  model.SpeciesID
		expected   *model.Member
		ok         bool
	}{
		{
			name:       "Valid inputs",
			memberName: "Becca",
			dob:        time.Date(2015, 6, 10, 0, 0, 0, 0, time.UTC),
			speciesID:  model.Human,
			expected: model.MemberSerialization{
				Name:    "Becca",
				DOB:     time.Date(2015, 6, 10, 0, 0, 0, 0, time.UTC),
				Species: model.Human,
			}.Deserialize(),
			ok: true,
		},
		{
			name:       "Invalid name",
			memberName: "",
			dob:        time.Date(2015, 6, 10, 0, 0, 0, 0, time.UTC),
			speciesID:  model.Human,
			expected:   nil,
			ok:         false,
		},
		{
			name:       "Empty DOB",
			memberName: "Becca",
			dob:        time.Time{},
			speciesID:  model.Human,
			expected:   nil,
			ok:         false,
		},
		{
			name:       "Future DOB",
			memberName: "Becca",
			dob:        time.Now().Add(time.Hour),
			speciesID:  model.Human,
			expected:   nil,
			ok:         false,
		},
		{
			name:       "Invalid Species",
			memberName: "Becca",
			dob:        time.Date(2015, 6, 10, 0, 0, 0, 0, time.UTC),
			speciesID:  model.Unknown,
			expected:   nil,
			ok:         false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := model.NewMember(tc.memberName, tc.dob, tc.speciesID)
			if tc.ok {
				assert.Empty(t, err)
				assert.Equal(t, tc.expected.Name(), actual.Name())
				assert.Equal(t, tc.expected.DOB(), actual.DOB())
				assert.Equal(t, tc.expected.Species(), actual.Species())
				assert.Equal(t, tc.expected.Image(), actual.Image())
			} else {
				assert.NotEmpty(t, err)
				assert.Empty(t, actual)
			}
		})
	}
}

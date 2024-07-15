package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemberName_Normalize(t *testing.T) {
	for _, tc := range []struct {
		name     string
		original MemberName
		expected MemberName
		ok       bool
	}{
		{
			name:     "Single word, already normalized",
			original: "Becca",
			expected: "Becca",
			ok:       true,
		},
		{
			name:     "Single word, extra leading whitespace",
			original: " Becca",
			expected: "Becca",
			ok:       true,
		},
		{
			name:     "Empty is invalid",
			original: "",
			expected: "",
			ok:       false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.original.Normalize()
			if tc.ok {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, actual)
			} else {
				assert.Error(t, err)
				assert.Empty(t, actual)
			}
		})
	}
}

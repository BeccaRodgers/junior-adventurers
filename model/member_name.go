package model

import (
	"errors"
	"strings"
)

type MemberName string

func (n MemberName) Normalize() (MemberName, error) {
	name := string(n)

	// Remove extra whitespace.
	words := strings.Fields(name)
	name = strings.Join(words, " ")
	
	if len(name) == 0 {
		return "", errors.New("cannot be empty")
	}

	return MemberName(name), nil
}

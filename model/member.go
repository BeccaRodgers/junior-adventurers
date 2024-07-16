package model

import (
	"artk.dev/apperror"
	"errors"
	"fmt"
	"time"
)

type Member struct {
	id      MemberID
	name    MemberName
	dob     time.Time
	species SpeciesID
	image   MemberImage
}

func (m *Member) ID() MemberID {
	return m.id
}

func (m *Member) Name() MemberName {
	return m.name
}

func (m *Member) DOB() time.Time {
	return m.dob
}

func (m *Member) Age() int {
	year, month, day := m.dob.Date()
	currentYear, currentMonth, currentDay := time.Now().Date()

	years := currentYear - year
	months := currentMonth - month
	days := currentDay - day
	if days < 0 {
		months--
	}
	if months < 0 {
		years--
	}
	return years
}

func (m *Member) Species() SpeciesID {
	return m.species
}

func (m *Member) Image() MemberImage {
	return m.image
}

func (m *Member) Serialize() MemberSerialization {
	return MemberSerialization{
		ID:      m.id,
		Name:    m.name,
		DOB:     m.dob,
		Species: m.species,
		Image:   m.image,
	}
}

type MemberSerialization struct {
	ID      MemberID
	Name    MemberName
	DOB     time.Time
	Species SpeciesID
	Image   MemberImage
}

func (s MemberSerialization) Deserialize() *Member {
	return &Member{
		id:      s.ID,
		name:    s.Name,
		dob:     s.DOB,
		species: s.Species,
		image:   s.Image,
	}
}

func NewMember(id MemberID, name MemberName, dob time.Time, speciesID SpeciesID) (*Member, *NewMemberError) {
	// TODO validate inputs
	errs := &NewMemberError{}
	name, err := name.Normalize()
	if err != nil {
		errs.Name = fmt.Errorf("invalid member name: %w", err)
	}

	if err = ValidateDOB(dob); err != nil {
		errs.DOB = fmt.Errorf("invalid date of birth: %w", err)
	}

	if errs != nil {
		return nil, errs
	}
	// TODO enforce rules.

	return &Member{
		id:      id,
		name:    name,
		dob:     dob,
		species: speciesID,
		image:   "",
	}, nil
}

type MemberID int64
type MemberImage string

var _ error = NewMemberError{}

type NewMemberError struct {
	Name error
	DOB  error
}

func (n NewMemberError) Error() string {
	return errors.Join(n.Name, n.DOB).Error()
}

func (n NewMemberError) Validation() bool {
	return true
}

func (n NewMemberError) Kind() apperror.Kind {
	return apperror.ValidationError
}

func ValidateDOB(dob time.Time) error {
	if dob.IsZero() {
		return errors.New("cannot be empty")
	}

	if dob.After(time.Now()) {
		return errors.New("cannot be in the future")
	}
	return nil
}

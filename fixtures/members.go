package fixtures

import (
	"context"
	"errors"
	"junior-adventurers/model"
	"time"
)

var AngelaId = model.NextAvailableID()

func AngelaID() model.MemberID {
	return AngelaId
}

func AngelaName() model.MemberName {
	return "Angela"
}

func AngelaDOB() time.Time {
	return time.Date(2014, 1, 7, 0, 0, 0, 0, time.UTC)
}

func AngelaSpecies() model.SpeciesID {
	return model.Human
}

func AngelaSerialization() model.MemberSerialization {
	return model.MemberSerialization{
		ID:      AngelaID(),
		Name:    AngelaName(),
		DOB:     AngelaDOB(),
		Species: AngelaSpecies(),
	}
}

func Angela() *model.Member {
	return AngelaSerialization().Deserialize()
}

var BobId = model.NextAvailableID()

func BobID() model.MemberID {
	return BobId
}

func Bob() *model.Member {
	return model.MemberSerialization{
		ID:      BobID(),
		Name:    "Bob",
		DOB:     time.Date(2011, 8, 9, 0, 0, 0, 0, time.UTC),
		Species: model.Human,
	}.Deserialize()
}

var CharlotteId = model.NextAvailableID()

func CharlotteID() model.MemberID {
	return CharlotteId
}

func Charlotte() *model.Member {
	return model.MemberSerialization{
		ID:      CharlotteID(),
		Name:    "Charlotte",
		DOB:     time.Date(2012, 11, 17, 0, 0, 0, 0, time.UTC),
		Species: model.Werewolf,
	}.Deserialize()
}

var DavidId = model.NextAvailableID()

func DavidID() model.MemberID {
	return DavidId
}

func David() *model.Member {
	return model.MemberSerialization{
		ID:      DavidID(),
		Name:    "David",
		DOB:     time.Date(2012, 8, 10, 0, 0, 0, 0, time.UTC),
		Species: model.Dwarf,
	}.Deserialize()
}

var BeyonceId = model.NextAvailableID()

func BeyonceID() model.MemberID {
	return BeyonceId
}

func BeyonceImage() model.MemberImage {
	return "Beyonce.png"
}

func Beyonce() *model.Member {
	return model.MemberSerialization{
		ID:      BeyonceID(),
		Name:    "Beyoncé",
		DOB:     time.Date(1982, 4, 10, 0, 0, 0, 0, time.UTC),
		Species: model.Dwarf,
		Image:   BeyonceImage(),
	}.Deserialize()
}

var AlexId = model.NextAvailableID()

func AlexID() model.MemberID {
	return AlexId
}

func Alex() *model.Member {
	return model.MemberSerialization{
		ID:      AlexID(),
		Name:    "Alex",
		DOB:     time.Date(1984, 9, 10, 0, 0, 0, 0, time.UTC),
		Species: model.Human,
	}.Deserialize()
}

var CarlosId = model.NextAvailableID()

func CarlosID() model.MemberID {
	return CarlosId
}

func Carlos() *model.Member {
	return model.MemberSerialization{
		ID:      CarlosID(),
		Name:    "Carlos",
		DOB:     time.Date(1984, 9, 10, 0, 0, 0, 0, time.UTC),
		Species: model.Werewolf,
	}.Deserialize()
}

var ErikaId = model.NextAvailableID()

func ErikaID() model.MemberID {
	return ErikaId
}

func Erika() *model.Member {
	return model.MemberSerialization{
		ID:      ErikaID(),
		Name:    "Erika",
		DOB:     time.Date(2012, 6, 10, 0, 0, 0, 0, time.UTC),
		Species: model.Human,
	}.Deserialize()
}

var FredId = model.NextAvailableID()

func FredID() model.MemberID {
	return FredId
}

func Fred() *model.Member {
	return model.MemberSerialization{
		ID:      FredID(),
		Name:    "Fred",
		DOB:     time.Date(2016, 6, 10, 0, 0, 0, 0, time.UTC),
		Species: model.Human,
	}.Deserialize()
}

var GeorgeId = model.NextAvailableID()

func GeorgeID() model.MemberID {
	return GeorgeId
}

func George() *model.Member {
	return model.MemberSerialization{
		ID:      GeorgeID(),
		Name:    "George",
		DOB:     time.Date(2015, 6, 10, 0, 0, 0, 0, time.UTC),
		Species: model.Human,
	}.Deserialize()
}

func InsertMembers(ctx context.Context, guilds model.MemberRepository) error {
	var errs []error
	for _, member := range []*model.Member{
		Angela(),
		Bob(),
		Charlotte(),
		David(),
		Beyonce(),
		Alex(),
		Carlos(),
		Erika(),
		Fred(),
		George(),
	} {
		if err := guilds.Insert(ctx, member); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

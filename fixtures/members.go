package fixtures

import (
	"context"
	"errors"
	"junior-adventurers/model"
	"time"
)

func AngelaID() model.MemberID {
	return 1
}

func AngelaName() model.MemberName {
	return "Angela"
}

func AngelaDOB() time.Time {
	return time.Date(2014, 1, 7, 0, 0, 0, 0, time.UTC)
}

func AngelaSpecies() model.MemberSpecies {
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

func BobID() model.MemberID {
	return 2
}

func Bob() *model.Member {
	return model.MemberSerialization{
		ID:      BobID(),
		Name:    "Bob",
		DOB:     time.Date(2011, 8, 9, 0, 0, 0, 0, time.UTC),
		Species: model.Human,
	}.Deserialize()
}

func CharlotteID() model.MemberID {
	return 3
}

func Charlotte() *model.Member {
	return model.MemberSerialization{
		ID:      CharlotteID(),
		Name:    "Charlotte",
		DOB:     time.Date(2012, 11, 17, 0, 0, 0, 0, time.UTC),
		Species: model.Werewolf,
	}.Deserialize()
}

func DavidID() model.MemberID {
	return 4
}

func David() *model.Member {
	return model.MemberSerialization{
		ID:      DavidID(),
		Name:    "David",
		DOB:     time.Date(2012, 8, 10, 0, 0, 0, 0, time.UTC),
		Species: model.Dwarf,
	}.Deserialize()
}

func BeyonceID() model.MemberID {
	return 5
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

func AlexID() model.MemberID {
	return 6
}

func Alex() *model.Member {
	return model.MemberSerialization{
		ID:      AlexID(),
		Name:    "Alex",
		DOB:     time.Date(1984, 9, 10, 0, 0, 0, 0, time.UTC),
		Species: model.Human,
	}.Deserialize()
}

func CarlosID() model.MemberID {
	return 7
}

func Carlos() *model.Member {
	return model.MemberSerialization{
		ID:      CarlosID(),
		Name:    "Carlos",
		DOB:     time.Date(1984, 9, 10, 0, 0, 0, 0, time.UTC),
		Species: model.Werewolf,
	}.Deserialize()
}

func ErikaID() model.MemberID {
	return 8
}

func Erika() *model.Member {
	return model.MemberSerialization{
		ID:      ErikaID(),
		Name:    "Erika",
		DOB:     time.Date(2010, 6, 10, 0, 0, 0, 0, time.UTC),
		Species: model.Human,
	}.Deserialize()
}

func FredID() model.MemberID {
	return 9
}

func Fred() *model.Member {
	return model.MemberSerialization{
		ID:      FredID(),
		Name:    "Fred",
		DOB:     time.Date(2010, 6, 10, 0, 0, 0, 0, time.UTC),
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
	} {
		if err := guilds.Insert(ctx, member); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

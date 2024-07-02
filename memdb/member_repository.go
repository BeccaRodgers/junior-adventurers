package memdb

import (
	"artk.dev/crud"
	"fmt"
	"junior-adventurers/model"
)

var _ model.MemberRepository = &MemberRepository{}

type MemberRepository struct {
	crud.InMemoryRepository[*model.Member, model.MemberID, model.MemberSerialization]
}

func NewMemberRepository() *MemberRepository {
	r := &MemberRepository{}
	r.Errors.NotFound = func(id model.MemberID) string {
		return fmt.Sprintf("member not found with ID %v", id)
	}

	r.Reset()
	return r
}

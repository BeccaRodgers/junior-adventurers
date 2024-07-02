package model

import "artk.dev/crud"

type MemberRepository interface {
	crud.Repository[*Member, MemberID, MemberSerialization]
}

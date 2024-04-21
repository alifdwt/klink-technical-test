package mapper

type Mapper struct {
	UserMapper   UserMapping
	MemberMapper MemberMapping
}

func NewMapper() *Mapper {
	return &Mapper{
		UserMapper:   NewUserMapper(),
		MemberMapper: NewMemberMapper(),
	}
}

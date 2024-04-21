package responses

type UserResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Privilege string `json:"privilege"`
}

type UserWithRelationResponse struct {
	ID        int              `json:"id"`
	Username  string           `json:"username"`
	Privilege string           `json:"privilege"`
	Member    []MemberResponse `json:"member"`
}

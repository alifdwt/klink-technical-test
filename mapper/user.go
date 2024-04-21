package mapper

import (
	"technical-test/domain/responses"
	"technical-test/models"
)

type userMapper struct {
}

func NewUserMapper() *userMapper {
	return &userMapper{}
}

func (m *userMapper) ToUserResponse(request *models.User) *responses.UserResponse {
	return &responses.UserResponse{
		ID:        request.ID,
		Username:  request.Username,
		Privilege: request.Privilege,
	}
}

func (m *userMapper) ToUserWithRelationResponse(request *models.User) *responses.UserWithRelationResponse {
	var members []responses.MemberResponse
	if request.Member != nil {
		for _, member := range request.Member {
			members = append(members, *NewMemberMapper().ToMemberResponse(&member))
		}
	}

	return &responses.UserWithRelationResponse{
		ID:       request.ID,
		Username: request.Username,
		Member:   members,
	}
}

func (m *userMapper) ToUserWithRelationResponses(requests []*models.User) []responses.UserWithRelationResponse {
	var userResponses []responses.UserWithRelationResponse

	for _, request := range requests {
		response := m.ToUserWithRelationResponse(request)
		userResponses = append(userResponses, *response)
	}

	if len(userResponses) > 0 {
		return userResponses
	}

	return nil
}

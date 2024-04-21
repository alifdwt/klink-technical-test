package mapper

import (
	"technical-test/domain/responses"
	"technical-test/models"
)

type memberMapper struct{}

func NewMemberMapper() *memberMapper {
	return &memberMapper{}
}

func (m *memberMapper) ToMemberResponse(request *models.Member) *responses.MemberResponse {
	return &responses.MemberResponse{
		ID:           request.ID,
		UserID:       request.UserID,
		GenderID:     request.GenderID,
		LevelID:      request.LevelID,
		NamaDepan:    request.NamaDepan,
		NamaBelakang: request.NamaBelakang,
		Birthdate:    request.Birthdate,
		JoinDate:     request.JoinDate,
		UpdateAt:     request.UpdatedAt,
	}
}

func (m *memberMapper) ToMemberWithRelationResponse(request *models.Member) *responses.MemberWithRelationResponse {
	return &responses.MemberWithRelationResponse{
		ID:           request.ID,
		UserID:       request.UserID,
		GenderID:     request.GenderID,
		LevelID:      request.LevelID,
		NamaDepan:    request.NamaDepan,
		NamaBelakang: request.NamaBelakang,
		Birthdate:    request.Birthdate,
		JoinDate:     request.JoinDate,
		UpdateAt:     request.UpdatedAt,
		User:         *NewUserMapper().ToUserResponse(&request.User),
		Gender:       responses.GenderResponse(request.Gender),
		Level:        responses.LevelResponse(request.Level),
	}
}

func (m *memberMapper) ToMemberResponses(requests *[]models.Member) []responses.MemberResponse {
	var memberResponses []responses.MemberResponse

	for _, request := range *requests {
		response := m.ToMemberResponse(&request)
		memberResponses = append(memberResponses, *response)
	}

	if len(memberResponses) > 0 {
		return memberResponses
	}

	return nil
}

func (m *memberMapper) ToMemberWithRelationResponses(requests *[]models.Member) []responses.MemberWithRelationResponse {
	var memberResponses []responses.MemberWithRelationResponse

	for _, request := range *requests {
		response := m.ToMemberWithRelationResponse(&request)
		memberResponses = append(memberResponses, *response)
	}

	if len(memberResponses) > 0 {
		return memberResponses
	}

	return []responses.MemberWithRelationResponse{}
}

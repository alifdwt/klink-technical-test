package mapper

import (
	"technical-test/domain/responses"
	"technical-test/models"
)

type UserMapping interface {
	ToUserResponse(request *models.User) *responses.UserResponse
	ToUserWithRelationResponse(request *models.User) *responses.UserWithRelationResponse
	ToUserWithRelationResponses(requests []*models.User) []responses.UserWithRelationResponse
}

type MemberMapping interface {
	ToMemberResponse(request *models.Member) *responses.MemberResponse
	ToMemberWithRelationResponse(request *models.Member) *responses.MemberWithRelationResponse
	ToMemberResponses(requests *[]models.Member) []responses.MemberResponse
	ToMemberWithRelationResponses(requests *[]models.Member) []responses.MemberWithRelationResponse
}

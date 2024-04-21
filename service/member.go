package service

import (
	"technical-test/domain/requests/member"
	"technical-test/domain/responses"
	"technical-test/mapper"
	"technical-test/pkg/logger"
	"technical-test/repository"
)

type memberService struct {
	repository repository.MemberRepository
	log        logger.Logger
	mapper     mapper.MemberMapping
}

func NewMemberService(member repository.MemberRepository, log logger.Logger, mapper mapper.MemberMapping) *memberService {
	return &memberService{
		repository: member,
		log:        log,
		mapper:     mapper,
	}
}

func (s *memberService) CreateMember(userId int, genderId int, levelId int, request member.CreateMemberRequest) (*responses.MemberResponse, error) {
	res, err := s.repository.CreateMember(userId, genderId, levelId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToMemberResponse(res)

	return mapper, nil
}

func (s *memberService) GetMemberAll() (*[]responses.MemberResponse, error) {
	res, err := s.repository.GetMemberAll()
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToMemberResponses(res)

	return &mapper, nil
}

func (s *memberService) GetMemberByUserId(userId int) (*[]responses.MemberWithRelationResponse, error) {
	res, err := s.repository.GetMemberByUserId(userId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToMemberWithRelationResponses(res)

	return &mapper, nil
}

func (s *memberService) GetMemberById(id int) (*responses.MemberWithRelationResponse, error) {
	res, err := s.repository.GetMemberById(id)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToMemberWithRelationResponse(res)

	return mapper, nil
}

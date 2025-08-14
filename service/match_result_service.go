package service

import (
	"ayo-indonesia/dto"
	"ayo-indonesia/model"
	"ayo-indonesia/repository"
)

type MatchResultService interface {
	Create(req dto.CreateMatchResultRequest) (dto.MatchResultResponse, error)
	GetAll() ([]dto.MatchResultResponse, error)
	GetByID(id uint) (dto.MatchResultResponse, error)
	Update(id uint, req dto.CreateMatchResultRequest) (dto.MatchResultResponse, error)
	Delete(id uint) error
}

type matchResultService struct {
	matchResultRepo repository.MatchResultRepository
}

func NewMatchResultService(matchResultRepo repository.MatchResultRepository) MatchResultService {
	return &matchResultService{
		matchResultRepo: matchResultRepo,
	}
}

func (s *matchResultService) Create(req dto.CreateMatchResultRequest) (dto.MatchResultResponse, error) {
	result := model.MatchResult{
		MatchID:      req.MatchID,
		HomeScore:    req.HomeScore,
		AwayScore:    req.AwayScore,
		WinnerTeamID: req.WinnerTeamID,
	}
	err := s.matchResultRepo.Create(&result)
	if err != nil {
		return dto.MatchResultResponse{}, err
	}
	return toMatchResultResponse(result), nil
}

func (s *matchResultService) GetAll() ([]dto.MatchResultResponse, error) {
	results, err := s.matchResultRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var res []dto.MatchResultResponse
	for _, r := range results {
		res = append(res, toMatchResultResponse(r))
	}
	return res, nil
}

func (s *matchResultService) GetByID(id uint) (dto.MatchResultResponse, error) {
	result, err := s.matchResultRepo.GetByID(id)
	if err != nil {
		return dto.MatchResultResponse{}, err
	}
	return toMatchResultResponse(result), nil
}

func (s *matchResultService) Update(id uint, req dto.CreateMatchResultRequest) (dto.MatchResultResponse, error) {
	result := model.MatchResult{
		MatchID:      req.MatchID,
		HomeScore:    req.HomeScore,
		AwayScore:    req.AwayScore,
		WinnerTeamID: req.WinnerTeamID,
	}
	err := s.matchResultRepo.Update(id, result)
	if err != nil {
		return dto.MatchResultResponse{}, err
	}
	updated, _ := s.matchResultRepo.GetByID(id)
	return toMatchResultResponse(updated), nil
}

func (s *matchResultService) Delete(id uint) error {
	result, err := s.matchResultRepo.GetByID(id)
	if err != nil {
		return err
	}
	return s.matchResultRepo.Delete(&result)
}

func toMatchResultResponse(m model.MatchResult) dto.MatchResultResponse {
	return dto.MatchResultResponse{
		ID:           m.ID,
		MatchID:      m.MatchID,
		HomeScore:    m.HomeScore,
		AwayScore:    m.AwayScore,
		WinnerTeamID: m.WinnerTeamID,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

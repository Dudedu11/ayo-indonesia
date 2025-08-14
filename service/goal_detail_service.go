package service

import (
	"ayo-indonesia/dto"
	"ayo-indonesia/model"
	"ayo-indonesia/repository"
	"errors"
)

type GoalDetailService interface {
	Create(req dto.CreateGoalDetailRequest) (dto.GoalDetailResponse, error)
	GetAll() ([]dto.GoalDetailResponse, error)
	GetByID(id uint) (dto.GoalDetailResponse, error)
	Update(id uint, req dto.CreateGoalDetailRequest) (dto.GoalDetailResponse, error)
	Delete(id uint) error
}

type goalDetailService struct {
	goalDetailRepo  repository.GoalDetailRepository
	matchRepo       repository.MatchRepository
	playerRepo      repository.PlayerRepository
	teamRepo        repository.TeamRepository
	matchResultRepo repository.MatchResultRepository
}

func NewGoalDetailService(
	goalDetailRepo repository.GoalDetailRepository,
	matchRepo repository.MatchRepository,
	playerRepo repository.PlayerRepository,
	teamRepo repository.TeamRepository,
	matchResultRepo repository.MatchResultRepository,
) GoalDetailService {
	return &goalDetailService{
		goalDetailRepo:  goalDetailRepo,
		matchRepo:       matchRepo,
		playerRepo:      playerRepo,
		teamRepo:        teamRepo,
		matchResultRepo: matchResultRepo,
	}
}

func (s *goalDetailService) Create(req dto.CreateGoalDetailRequest) (dto.GoalDetailResponse, error) {
	match, err := s.matchRepo.GetByID(req.MatchID)
	if err != nil {
		return dto.GoalDetailResponse{}, errors.New("match not found")
	}
	if _, err := s.playerRepo.GetByID(req.PlayerID); err != nil {
		return dto.GoalDetailResponse{}, errors.New("player not found")
	}
	if _, err := s.teamRepo.GetByID(req.TeamID); err != nil {
		return dto.GoalDetailResponse{}, errors.New("team not found")
	}

	goal := model.GoalDetail{
		MatchID:  req.MatchID,
		PlayerID: req.PlayerID,
		TeamID:   req.TeamID,
		Minute:   req.Minute,
	}

	if err := s.goalDetailRepo.Create(&goal); err != nil {
		return dto.GoalDetailResponse{}, err
	}

	result, err := s.matchResultRepo.GetByMatchID(req.MatchID)
	if err != nil {
		homeScore := 0
		awayScore := 0
		if match.HomeTeamID == req.TeamID {
			homeScore = 1
		} else {
			awayScore = 1
		}
		result = model.MatchResult{
			MatchID:      req.MatchID,
			HomeScore:    homeScore,
			AwayScore:    awayScore,
			WinnerTeamID: &req.TeamID,
		}
		if err := s.matchResultRepo.Create(&result); err != nil {
			return dto.GoalDetailResponse{}, err
		}
	} else {
		if req.TeamID == match.HomeTeamID {
			result.HomeScore++
		} else if req.TeamID == match.AwayTeamID {
			result.AwayScore++
		}

		if result.HomeScore > result.AwayScore {
			result.WinnerTeamID = &match.HomeTeamID
		} else if result.HomeScore < result.AwayScore {
			result.WinnerTeamID = &match.AwayTeamID
		} else {
			result.WinnerTeamID = nil
		}
		if err := s.matchResultRepo.Update(result.ID, result); err != nil {
			return dto.GoalDetailResponse{}, err
		}
	}

	return toGoalDetailResponse(goal), nil
}

func (s *goalDetailService) GetAll() ([]dto.GoalDetailResponse, error) {
	goals, err := s.goalDetailRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var res []dto.GoalDetailResponse
	for _, g := range goals {
		res = append(res, toGoalDetailResponse(g))
	}

	return res, nil
}

func (s *goalDetailService) GetByID(id uint) (dto.GoalDetailResponse, error) {
	goal, err := s.goalDetailRepo.GetByID(id)
	if err != nil {
		return dto.GoalDetailResponse{}, err
	}
	return toGoalDetailResponse(goal), nil
}

func (s *goalDetailService) Update(id uint, req dto.CreateGoalDetailRequest) (dto.GoalDetailResponse, error) {
	if _, err := s.matchRepo.GetByID(req.MatchID); err != nil {
		return dto.GoalDetailResponse{}, errors.New("match not found")
	}
	if _, err := s.playerRepo.GetByID(req.PlayerID); err != nil {
		return dto.GoalDetailResponse{}, errors.New("player not found")
	}
	if _, err := s.teamRepo.GetByID(req.TeamID); err != nil {
		return dto.GoalDetailResponse{}, errors.New("team not found")
	}

	goal := model.GoalDetail{
		MatchID:  req.MatchID,
		PlayerID: req.PlayerID,
		TeamID:   req.TeamID,
		Minute:   req.Minute,
	}

	if err := s.goalDetailRepo.Update(id, goal); err != nil {
		return dto.GoalDetailResponse{}, err
	}

	updated, _ := s.goalDetailRepo.GetByID(id)
	return toGoalDetailResponse(updated), nil
}

func (s *goalDetailService) Delete(id uint) error {
	goal, err := s.goalDetailRepo.GetByID(id)
	if err != nil {
		return err
	}
	return s.goalDetailRepo.Delete(&goal)
}

func toGoalDetailResponse(g model.GoalDetail) dto.GoalDetailResponse {
	return dto.GoalDetailResponse{
		ID:        g.ID,
		MatchID:   g.MatchID,
		PlayerID:  g.PlayerID,
		TeamID:    g.TeamID,
		Minute:    g.Minute,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}

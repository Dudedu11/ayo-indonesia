package service

import (
	"ayo-indonesia/dto"
	"ayo-indonesia/model"
	"ayo-indonesia/repository"
	"errors"
)

type MatchService interface {
	Create(req dto.CreateMatchRequest) (dto.MatchResponse, error)
	GetAll() ([]dto.GetMatchResponse, error)
	GetByID(id uint) (dto.MatchResponse, error)
	Update(id uint, req dto.CreateMatchRequest) (dto.MatchResponse, error)
	Delete(id uint) error
}

type matchService struct {
	matchRepo       repository.MatchRepository
	teamRepo        repository.TeamRepository
	matchResultRepo repository.MatchResultRepository
	goalDetailRepo  repository.GoalDetailRepository
	playerRepo      repository.PlayerRepository
}

func NewMatchService(matchRepo repository.MatchRepository,
	teamRepo repository.TeamRepository,
	matchResultRepo repository.MatchResultRepository,
	goalDetailRepo repository.GoalDetailRepository,
	playerRepo repository.PlayerRepository,
) MatchService {
	return &matchService{
		matchRepo:       matchRepo,
		teamRepo:        teamRepo,
		matchResultRepo: matchResultRepo,
		goalDetailRepo:  goalDetailRepo,
		playerRepo:      playerRepo,
	}
}

func (s *matchService) Create(req dto.CreateMatchRequest) (dto.MatchResponse, error) {
	if _, err := s.teamRepo.GetByID(req.HomeTeamID); err != nil {
		return dto.MatchResponse{}, errors.New("home team not found")
	}
	if _, err := s.teamRepo.GetByID(req.AwayTeamID); err != nil {
		return dto.MatchResponse{}, errors.New("away team not found")
	}

	match := model.Match{
		Date:       req.Date,
		Time:       req.Time,
		HomeTeamID: req.HomeTeamID,
		AwayTeamID: req.AwayTeamID,
	}

	err := s.matchRepo.Create(&match)
	if err != nil {
		return dto.MatchResponse{}, err
	}

	return toMatchResponse(match), nil
}

func (s *matchService) GetAll() ([]dto.GetMatchResponse, error) {
	matches, err := s.matchRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var res []dto.GetMatchResponse
	for _, m := range matches {
		match, _ := s.toGetMatchResponse(m)
		res = append(res, match)
	}

	return res, nil
}

func (s *matchService) GetByID(id uint) (dto.MatchResponse, error) {
	match, err := s.matchRepo.GetByID(id)
	if err != nil {
		return dto.MatchResponse{}, err
	}

	return toMatchResponse(match), nil
}

func (s *matchService) Update(id uint, req dto.CreateMatchRequest) (dto.MatchResponse, error) {
	if _, err := s.teamRepo.GetByID(req.HomeTeamID); err != nil {
		return dto.MatchResponse{}, errors.New("home team not found")
	}
	if _, err := s.teamRepo.GetByID(req.AwayTeamID); err != nil {
		return dto.MatchResponse{}, errors.New("away team not found")
	}

	match := model.Match{
		Date:       req.Date,
		Time:       req.Time,
		HomeTeamID: req.HomeTeamID,
		AwayTeamID: req.AwayTeamID,
	}

	err := s.matchRepo.Update(id, match)
	if err != nil {
		return dto.MatchResponse{}, err
	}

	updated, _ := s.matchRepo.GetByID(id)

	return toMatchResponse(updated), nil
}

func (s *matchService) Delete(id uint) error {
	match, err := s.matchRepo.GetByID(id)
	if err != nil {
		return err
	}
	return s.matchRepo.Delete(&match)
}

func toMatchResponse(m model.Match) dto.MatchResponse {
	return dto.MatchResponse{
		ID:         m.ID,
		Date:       m.Date,
		Time:       m.Time,
		HomeTeamID: m.HomeTeamID,
		AwayTeamID: m.AwayTeamID,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

func (s *matchService) toGetMatchResponse(m model.Match) (dto.GetMatchResponse, error) {
	var response dto.GetMatchResponse
	homeTeam, err := s.teamRepo.GetByID(m.HomeTeamID)
	if err != nil {
		return response, errors.New("home team not found")
	}

	awayTeam, err := s.teamRepo.GetByID(m.AwayTeamID)
	if err != nil {
		return response, errors.New("away team not found")
	}

	var matchDetail *dto.DetailMatchResult
	matchResult, err := s.matchResultRepo.GetByMatchID(m.ID)
	if err == nil {
		matchDetail = &dto.DetailMatchResult{
			HomeScore: matchResult.HomeScore,
			AwayScore: matchResult.AwayScore,
		}
		if matchResult.WinnerTeamID != nil {
			if team, err := s.teamRepo.GetByID(*matchResult.WinnerTeamID); err == nil {
				matchDetail.Winner = team.Name
			}
		}else{
			matchDetail.Winner = "Draw"
		}

		goals, err := s.goalDetailRepo.GetByMatchID(m.ID)
		if err != nil {
			return response, errors.New("failed get goal detail")
		}

		var goalDetail []dto.GoalDetail
		for _, goal := range goals {
			team, _ := s.teamRepo.GetByID(goal.TeamID)
			player, _ := s.playerRepo.GetByID(goal.PlayerID)
			goalDetail = append(goalDetail, dto.GoalDetail{
				Team: team.Name,
				Player: player.Name,
				Minute: goal.Minute,
			})
		}
		matchDetail.GoalDetail = goalDetail
	}

	return dto.GetMatchResponse{
		ID:                m.ID,
		Date:              m.Date,
		Time:              m.Time,
		HomeTeam:          homeTeam.Name,
		AwayTeam:          awayTeam.Name,
		DetailMatchResult: matchDetail,
	}, nil
}

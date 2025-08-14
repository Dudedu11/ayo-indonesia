package service

import (
	"ayo-indonesia/dto"
	"ayo-indonesia/model"
	"ayo-indonesia/repository"
	"fmt"
)

type TeamService interface {
	Create(req dto.CreateTeamRequest) (dto.TeamResponse, error)
	GetAll() ([]dto.GetTeamResponse, error)
	GetByID(id uint) (dto.GetTeamResponse, error) 
	Update(id uint, req dto.CreateTeamRequest) (dto.TeamResponse, error)
	Delete(id uint) error
}

type teamService struct {
	repo            repository.TeamRepository
	matchResultRepo repository.MatchResultRepository
}

func NewTeamService(repo repository.TeamRepository, matchResultRepo repository.MatchResultRepository) TeamService {
	return &teamService{
		repo:            repo,
		matchResultRepo: matchResultRepo,
	}
}

func (s *teamService) Create(req dto.CreateTeamRequest) (dto.TeamResponse, error) {
	team := model.Team{
		Name:        req.Name,
		LogoURL:     req.LogoURL,
		FoundedYear: req.FoundedYear,
		BaseAddress: req.BaseAddress,
		BaseCity:    req.BaseCity,
	}
	err := s.repo.Create(&team)
	if err != nil {
		return dto.TeamResponse{}, err
	}
	return toTeamResponse(team), nil
}

func (s *teamService) GetAll() ([]dto.GetTeamResponse, error) {
	teams, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	fmt.Println(teams)
	var res []dto.GetTeamResponse
	for _, t := range teams {
		teamRes, _ := s.toGetTeamResponse(t)
		res = append(res, teamRes)
	}
	return res, nil
}

func (s *teamService) GetByID(id uint) (dto.GetTeamResponse, error) {
	team, err := s.repo.GetByID(id)
	if err != nil {
		return dto.GetTeamResponse{}, err
	}
	res, _ := s.toGetTeamResponse(team)
	return res, nil
}

func (s *teamService) Update(id uint, req dto.CreateTeamRequest) (dto.TeamResponse, error) {
	team := model.Team{
		Name:        req.Name,
		LogoURL:     req.LogoURL,
		FoundedYear: req.FoundedYear,
		BaseAddress: req.BaseAddress,
		BaseCity:    req.BaseCity,
	}
	err := s.repo.Update(id, team)
	if err != nil {
		return dto.TeamResponse{}, err
	}
	updatedTeam, _ := s.repo.GetByID(id)
	return toTeamResponse(updatedTeam), nil

}

func (s *teamService) Delete(id uint) error {
	team, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(&team)
}

func toTeamResponse(t model.Team) dto.TeamResponse {
	return dto.TeamResponse{
		ID:          t.ID,
		Name:        t.Name,
		LogoURL:     t.LogoURL,
		FoundedYear: t.FoundedYear,
		BaseAddress: t.BaseAddress,
		BaseCity:    t.BaseCity,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func (s *teamService) toGetTeamResponse(t model.Team) (dto.GetTeamResponse, error) {
	summary := &dto.SummaryMatch{}

	winStats, err := s.matchResultRepo.GetWinStats(t.ID)
	if err == nil {
		summary.AwayWin = winStats.AwayWins
		summary.HomeWin = winStats.HomeWins
	}

	return dto.GetTeamResponse{
		ID:           t.ID,
		Name:         t.Name,
		LogoURL:      t.LogoURL,
		FoundedYear:  t.FoundedYear,
		BaseAddress:  t.BaseAddress,
		BaseCity:     t.BaseCity,
		SummaryMatch: summary,
	}, nil
}


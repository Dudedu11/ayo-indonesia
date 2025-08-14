package service

import (
	"ayo-indonesia/constant"
	"ayo-indonesia/dto"
	"ayo-indonesia/model"
	"ayo-indonesia/repository"
	"errors"
)

type PlayerService interface {
	Create(req dto.CreatePlayerRequest) (dto.PlayerResponse, error)
	GetAll() ([]dto.PlayerResponse, error)
	GetByID(id uint) (dto.PlayerResponse, error)
	Update(id uint, req dto.CreatePlayerRequest) (dto.PlayerResponse, error)
	Delete(id uint) error
	GetTopScore() ([]dto.TopScoreResponse, error) 
}

type playerService struct {
	playerRepo     repository.PlayerRepository
	goalDetailRepo repository.GoalDetailRepository
}

func NewPlayerService(playerRepo repository.PlayerRepository, goalDetailRepo repository.GoalDetailRepository) PlayerService {
	return &playerService{
		playerRepo:     playerRepo,
		goalDetailRepo: goalDetailRepo,
	}
}

func (s *playerService) Create(req dto.CreatePlayerRequest) (dto.PlayerResponse, error) {
	if !constant.IsValidPosition(req.Position) {
		return dto.PlayerResponse{}, errors.New("invalid position: must be one of penyerang, gelandang, bertahan, penjaga gawang")
	}

	if s.playerRepo.IsNumberUsed(req.TeamID, req.Number, nil) {
		return dto.PlayerResponse{}, errors.New("number already used in this team")
	}

	player := model.Player{
		Name:     req.Name,
		Height:   req.Height,
		Weight:   req.Weight,
		Position: req.Position,
		Number:   req.Number,
		TeamID:   req.TeamID,
	}

	err := s.playerRepo.Create(&player)
	if err != nil {
		return dto.PlayerResponse{}, err
	}

	return toPlayerResponse(player), nil
}

func (s *playerService) GetAll() ([]dto.PlayerResponse, error) {
	players, err := s.playerRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var res []dto.PlayerResponse
	for _, p := range players {
		res = append(res, toPlayerResponse(p))
	}
	return res, nil
}

func (s *playerService) GetByID(id uint) (dto.PlayerResponse, error) {
	player, err := s.playerRepo.GetByID(id)
	if err != nil {
		return dto.PlayerResponse{}, err
	}
	return toPlayerResponse(player), nil
}

func (s *playerService) Update(id uint, req dto.CreatePlayerRequest) (dto.PlayerResponse, error) {
	if s.playerRepo.IsNumberUsed(req.TeamID, req.Number, &id) {
		return dto.PlayerResponse{}, errors.New("number already used in this team")
	}
	player := model.Player{
		Name:     req.Name,
		Height:   req.Height,
		Weight:   req.Weight,
		Position: req.Position,
		Number:   req.Number,
		TeamID:   req.TeamID,
	}
	err := s.playerRepo.Update(id, player)
	if err != nil {
		return dto.PlayerResponse{}, err
	}
	updated, _ := s.playerRepo.GetByID(id)
	return toPlayerResponse(updated), nil
}

func (s *playerService) Delete(id uint) error {
	player, err := s.playerRepo.GetByID(id)
	if err != nil {
		return err
	}
	return s.playerRepo.Delete(&player)
}

func toPlayerResponse(p model.Player) dto.PlayerResponse {
	return dto.PlayerResponse{
		ID:        p.ID,
		Name:      p.Name,
		Height:    p.Height,
		Weight:    p.Weight,
		Position:  p.Position,
		Number:    p.Number,
		TeamID:    p.TeamID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (s *playerService) GetTopScore() ([]dto.TopScoreResponse, error) {
	goals, err := s.goalDetailRepo.GetTopScore()
	if err != nil {
		return nil, err
	}

	var res []dto.TopScoreResponse
	for _, goal := range goals {
		player, _ := s.playerRepo.GetByID(goal.PlayerID)
		res = append(res, dto.TopScoreResponse{
			Name: player.Name,
			Goal: goal.Goals,
		})

	}

	return res, nil
}

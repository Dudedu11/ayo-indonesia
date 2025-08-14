package repository

import (
	"ayo-indonesia/config"
	"ayo-indonesia/model"
)

type MatchRepository interface {
	Create(match *model.Match) error
	GetAll() ([]model.Match, error)
	GetByID(id uint) (model.Match, error)
	Update(id uint, match model.Match) error
	Delete(match *model.Match) error
}

type matchRepository struct{}

func NewMatchRepository() MatchRepository {
	return &matchRepository{}
}

func (r *matchRepository) Create(match *model.Match) error {
	return config.DB.Create(match).Error
}

func (r *matchRepository) GetAll() ([]model.Match, error) {
	var matches []model.Match
	err := config.DB.Where("deleted_at IS NULL").Find(&matches).Error
	return matches, err
}

func (r *matchRepository) GetByID(id uint) (model.Match, error) {
	var match model.Match
	err := config.DB.First(&match, id).Error
	return match, err
}

func (r *matchRepository) Update(id uint, updated model.Match) error {
	var match model.Match
	if err := config.DB.First(&match, id).Error; err != nil {
		return err
	}
	match.Date = updated.Date
	match.Time = updated.Time
	match.HomeTeamID = updated.HomeTeamID
	match.AwayTeamID = updated.AwayTeamID
	return config.DB.Save(&match).Error
}

func (r *matchRepository) Delete(match *model.Match) error {
	return config.DB.Delete(match).Error 
}

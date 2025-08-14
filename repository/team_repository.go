package repository

import (
	"ayo-indonesia/config"
	"ayo-indonesia/model"
)

type TeamRepository interface {
	Create(team *model.Team) error
	GetAll() ([]model.Team, error)
	GetByID(id uint) (model.Team, error)
	Update(id uint, team model.Team) error
	Delete(team *model.Team) error
}

type teamRepository struct{}

func NewTeamRepository() TeamRepository {
	return &teamRepository{}
}

func (r *teamRepository) Create(team *model.Team) error {
	return config.DB.Create(team).Error
}

func (r *teamRepository) GetAll() ([]model.Team, error) {
	var teams []model.Team
	err := config.DB.Where("deleted_at IS NULL").Find(&teams).Error
	return teams, err
}

func (r *teamRepository) GetByID(id uint) (model.Team, error) {
	var team model.Team
	err := config.DB.First(&team, id).Error
	return team, err
}

func (r *teamRepository) Update(id uint, updated model.Team) error {
	var team model.Team
	if err := config.DB.First(&team, id).Error; err != nil {
		return err
	}

	team.Name = updated.Name
	team.LogoURL = updated.LogoURL
	team.FoundedYear = updated.FoundedYear
	team.BaseAddress = updated.BaseAddress
	team.BaseCity = updated.BaseCity

	return config.DB.Save(&team).Error
}

func (r *teamRepository) Delete(team *model.Team) error {
	return config.DB.Delete(team).Error
}
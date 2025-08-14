package repository

import (
	"ayo-indonesia/config"
	"ayo-indonesia/model"
)

type PlayerRepository interface {
	Create(player *model.Player) error
	GetAll() ([]model.Player, error)
	GetByID(id uint) (model.Player, error)
	Update(id uint, player model.Player) error
	Delete(player *model.Player) error
	IsNumberUsed(teamID uint, number int, excludeID *uint) bool
}

type playerRepository struct{}

func NewPlayerRepository() PlayerRepository {
	return &playerRepository{}
}

func (r *playerRepository) Create(player *model.Player) error {
	return config.DB.Create(player).Error
}

func (r *playerRepository) GetAll() ([]model.Player, error) {
	var players []model.Player
	err := config.DB.Where("deleted_at IS NULL").Find(&players).Error
	return players, err
}

func (r *playerRepository) GetByID(id uint) (model.Player, error) {
	var player model.Player
	err := config.DB.First(&player, id).Error
	return player, err
}

func (r *playerRepository) Update(id uint, updated model.Player) error {
	var player model.Player
	if err := config.DB.First(&player, id).Error; err != nil {
		return err
	}
	player.Name = updated.Name
	player.Height = updated.Height
	player.Weight = updated.Weight
	player.Position = updated.Position
	player.Number = updated.Number
	player.TeamID = updated.TeamID
	return config.DB.Save(&player).Error
}

func (r *playerRepository) Delete(player *model.Player) error {
	return config.DB.Delete(player).Error
}

func (r *playerRepository) IsNumberUsed(teamID uint, number int, excludeID *uint) bool {
	var count int64
	query := config.DB.Model(&model.Player{}).Where("team_id = ? AND number = ? AND deleted_at IS NULL", teamID, number)
	if excludeID != nil {
		query = query.Where("id != ?", *excludeID)
	}
	query.Count(&count)
	return count > 0
}
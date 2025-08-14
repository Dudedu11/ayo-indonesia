package repository

import (
	"ayo-indonesia/config"
	"ayo-indonesia/dto"
	"ayo-indonesia/model"
)

type GoalDetailRepository interface {
	Create(goalDetail *model.GoalDetail) error
	GetAll() ([]model.GoalDetail, error)
	GetByID(id uint) (model.GoalDetail, error)
	GetByMatchID(matchID uint) ([]model.GoalDetail, error)
	Update(id uint, updated model.GoalDetail) error
	Delete(goalDetail *model.GoalDetail) error
	GetTopScore() ([]dto.TopScoreGoalResponse, error)
}

type goalDetailRepository struct{}

func NewGoalDetailRepository() GoalDetailRepository {
	return &goalDetailRepository{}
}

func (r *goalDetailRepository) Create(goalDetail *model.GoalDetail) error {
	return config.DB.Create(goalDetail).Error
}

func (r *goalDetailRepository) GetAll() ([]model.GoalDetail, error) {
	var goalDetails []model.GoalDetail
	err := config.DB.Where("deleted_at IS NULL").Find(&goalDetails).Error
	return goalDetails, err
}

func (r *goalDetailRepository) GetByID(id uint) (model.GoalDetail, error) {
	var goalDetail model.GoalDetail
	err := config.DB.First(&goalDetail, id).Error
	return goalDetail, err
}

func (r *goalDetailRepository) Update(id uint, updated model.GoalDetail) error {
	var goalDetail model.GoalDetail
	if err := config.DB.First(&goalDetail, id).Error; err != nil {
		return err
	}
	goalDetail.MatchID = updated.MatchID
	goalDetail.PlayerID = updated.PlayerID
	goalDetail.TeamID = updated.TeamID
	goalDetail.Minute = updated.Minute
	return config.DB.Save(&goalDetail).Error
}

func (r *goalDetailRepository) Delete(goalDetail *model.GoalDetail) error {
	return config.DB.Delete(goalDetail).Error
}

func (r *goalDetailRepository) GetByMatchID(matchID uint) ([]model.GoalDetail, error) {
	var goalDetails []model.GoalDetail
	err := config.DB.
		Where("match_id = ? AND deleted_at IS NULL", matchID).
		Order("minute ASC").
		Find(&goalDetails).Error
	return goalDetails, err
}

func (r *goalDetailRepository) GetTopScore() ([]dto.TopScoreGoalResponse, error) {
	var results []dto.TopScoreGoalResponse
	err := config.DB.
		Model(&model.GoalDetail{}).
		Select("player_id, COUNT(*) as goals").
		Where("deleted_at IS NULL").
		Group("player_id").
		Order("goals DESC").
		Scan(&results).Error

	return results, err
}

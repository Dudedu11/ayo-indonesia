package repository

import (
	"ayo-indonesia/config"
	"ayo-indonesia/dto"
	"ayo-indonesia/model"
)

type MatchResultRepository interface {
	Create(result *model.MatchResult) error
	GetAll() ([]model.MatchResult, error)
	GetByID(id uint) (model.MatchResult, error)
	GetByMatchID(matchID uint) (model.MatchResult, error) 
	Update(id uint, updated model.MatchResult) error
	Delete(result *model.MatchResult) error
	GetWinStats(teamID uint) (dto.WinStatsResponse, error)
}

type matchResultRepository struct{}

func NewMatchResultRepository() MatchResultRepository {
	return &matchResultRepository{}
}

func (r *matchResultRepository) Create(result *model.MatchResult) error {
	return config.DB.Create(result).Error
}

func (r *matchResultRepository) GetAll() ([]model.MatchResult, error) {
	var results []model.MatchResult
	err := config.DB.Where("deleted_at IS NULL").Find(&results).Error
	return results, err
}

func (r *matchResultRepository) GetByID(id uint) (model.MatchResult, error) {
	var result model.MatchResult
	err := config.DB.First(&result, id).Error
	return result, err
}

func (r *matchResultRepository) GetByMatchID(matchID uint) (model.MatchResult, error) {
	var result model.MatchResult
	err := config.DB.Where("match_id = ?", matchID).First(&result).Error
	return result, err
}

func (r *matchResultRepository) Update(id uint, updated model.MatchResult) error {
	var result model.MatchResult
	if err := config.DB.First(&result, id).Error; err != nil {
		return err
	}
	result.MatchID = updated.MatchID
	result.HomeScore = updated.HomeScore
	result.AwayScore = updated.AwayScore
	result.WinnerTeamID = updated.WinnerTeamID
	return config.DB.Save(&result).Error
}

func (r *matchResultRepository) Delete(result *model.MatchResult) error {
	return config.DB.Delete(result).Error
}

func (r *matchResultRepository) GetWinStats(teamID uint) (dto.WinStatsResponse, error) {
    var result dto.WinStatsResponse
    result.TeamID = teamID

    err := config.DB.
        Table("match_results").
        Select(`
            SUM(CASE WHEN matches.home_team_id = ? AND match_results.home_score > match_results.away_score THEN 1 ELSE 0 END) AS home_wins,
            SUM(CASE WHEN matches.away_team_id = ? AND match_results.away_score > match_results.home_score THEN 1 ELSE 0 END) AS away_wins
        `, teamID, teamID).
        Joins("JOIN matches ON matches.id = match_results.match_id").
        Scan(&result).Error

    return result, err
}


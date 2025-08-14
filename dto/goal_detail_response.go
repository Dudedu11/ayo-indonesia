package dto

import "time"

type GoalDetailResponse struct {
	ID        uint      `json:"id"`
	MatchID   uint      `json:"match_id"`
	PlayerID  uint      `json:"player_id"`
	TeamID    uint      `json:"team_id"`
	Minute    int       `json:"minute"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TopScoreGoalResponse struct {
	PlayerID uint `json:"player_id"`
	Goals int    `json:"goals"`
}
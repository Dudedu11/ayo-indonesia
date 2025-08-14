package model

import "gorm.io/gorm"

type GoalDetail struct {
	gorm.Model
	MatchID  uint `json:"match_id" `
	PlayerID uint `json:"player_id"`
	TeamID   uint `json:"team_id"`
	Minute   int  `json:"minute"`
}

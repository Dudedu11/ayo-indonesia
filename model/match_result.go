package model

import "gorm.io/gorm"

type MatchResult struct {
	gorm.Model
	MatchID      uint  `json:"match_id"`
	HomeScore    int   `json:"home_score"`
	AwayScore    int   `json:"away_score"`
	WinnerTeamID *uint `json:"winner_team_id"`
}

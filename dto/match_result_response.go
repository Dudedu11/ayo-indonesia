package dto

import "time"

type MatchResultResponse struct {
	ID           uint      `json:"id"`
	MatchID      uint      `json:"match_id"`
	HomeScore    int       `json:"home_score"`
	AwayScore    int       `json:"away_score"`
	WinnerTeamID *uint     `json:"winner_team_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type WinStatsResponse struct {
    TeamID    uint `json:"team_id"`
    HomeWins  int  `json:"home_wins"`
    AwayWins  int  `json:"away_wins"`
}
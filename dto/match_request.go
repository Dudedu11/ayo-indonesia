package dto

import "time"

type CreateMatchRequest struct {
	Date       time.Time `json:"date"`
	Time       string    `json:"time"`
	HomeTeamID uint      `json:"home_team_id"`
	AwayTeamID uint      `json:"away_team_id"`
}

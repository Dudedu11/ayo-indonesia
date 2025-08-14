package model

import (
	"time"

	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	Date       time.Time `json:"date"`
	Time       string    `json:"time"`
	HomeTeamID uint      `json:"home_team_id"`
	AwayTeamID uint      `json:"away_team_id"`
}
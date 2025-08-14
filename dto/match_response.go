package dto

import "time"

type MatchResponse struct {
	ID         uint      `json:"id"`
	Date       time.Time `json:"date"`
	Time       string    `json:"time"`
	HomeTeamID uint      `json:"home_team_id"`
	AwayTeamID uint      `json:"away_team_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetMatchResponse struct {
	ID                uint               `json:"id"`
	Date              time.Time          `json:"date"`
	Time              string             `json:"time"`
	HomeTeam          string             `json:"home_team"`
	AwayTeam          string             `json:"away_team"`
	DetailMatchResult *DetailMatchResult `json:"detail_match_result,omitempty"`
}

type DetailMatchResult struct {
	HomeScore  int    `json:"home_score"`
	AwayScore  int    `json:"away_score"`
	Winner     string `json:"winner"`
	GoalDetail []GoalDetail `json:"goal_detail,omitempty"`
}

type GoalDetail struct {
	Player string `json:"player"`
	Team   string `json:"team"`
	Minute int    `json:"minute"`
}

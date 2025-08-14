package dto

type CreateMatchResultRequest struct {
	MatchID      uint  `json:"match_id"`
	HomeScore    int   `json:"home_score"`
	AwayScore    int   `json:"away_score"`
	WinnerTeamID *uint `json:"winner_team_id"`
}
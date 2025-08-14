package dto

type CreateGoalDetailRequest struct {
	MatchID  uint `json:"match_id"`
	PlayerID uint `json:"player_id"`
	TeamID   uint `json:"team_id"`
	Minute   int  `json:"minute"`
}

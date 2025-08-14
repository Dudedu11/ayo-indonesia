package dto

type CreatePlayerRequest struct {
	Name     string  `json:"name"`
	Height   float64 `json:"height"`
	Weight   float64 `json:"weight"`
	Position string  `json:"position"`
	Number   int     `json:"number"`
	TeamID   uint    `json:"team_id"`
}
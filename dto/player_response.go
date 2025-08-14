package dto

import "time"

type PlayerResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Height    float64   `json:"height"`
	Weight    float64   `json:"weight"`
	Position  string    `json:"position"`
	Number    int       `json:"number"`
	TeamID    uint      `json:"team_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TopScoreResponse struct {
	Name string `json:"name"`
	Goal int    `json:"goal"`
}

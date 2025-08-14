package model

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name      string `json:"name"`
	Height    float64 `json:"height"`
	Weight    float64 `json:"wight"`
	Position  string `json:"position"`
	Number    int    `json:"number"`
	TeamID    uint   `json:"team_id"`
}
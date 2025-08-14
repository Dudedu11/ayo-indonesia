package dto

import "time"

type TeamResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	LogoURL     string    `json:"logo_url"`
	FoundedYear int       `json:"founded_year"`
	BaseAddress string    `json:"base_address"`
	BaseCity    string    `json:"base_city"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetTeamResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	LogoURL     string `json:"logo_url"`
	FoundedYear int    `json:"founded_year"`
	BaseAddress string `json:"base_address"`
	BaseCity    string `json:"base_city"`
	SummaryMatch *SummaryMatch `json:"summary_match,omitempty"`
}

type SummaryMatch struct {
	HomeWin int `json:"home_win"`
	AwayWin int `json:"away_win"`
}

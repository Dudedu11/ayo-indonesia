package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name        string `json:"name"`
	LogoURL     string `json:"logo_url"`
	FoundedYear int    `json:"founded_year"`
	BaseAddress string `json:"base_address"`
	BaseCity    string `json:"base_city"`
}

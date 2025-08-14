package dto

type CreateTeamRequest struct {
	Name        string `json:"name"`
	LogoURL     string `json:"logo_url"`
	FoundedYear int    `json:"founded_year"`
	BaseAddress string `json:"base_address"`
	BaseCity    string `json:"base_city"`
}

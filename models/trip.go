package models

import "time"

type Trip struct {
	ID          string    `json:"id"`
	Destination string    `json:"destination"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PlanningForm struct {
	PlanningMethod  string     `json:"planningMethod"`
	Locations       []Location `json:"locations"`
	StartDate       string     `json:"startDate"`
	EndDate         string     `json:"endDate"`
	Duration        int        `json:"duration"`
	BudgetLevel     string     `json:"budgetLevel,omitempty"`
	TravelIntensity string     `json:"travelIntensity,omitempty"`
	Interests       []string   `json:"interests,omitempty"`
}

type Location struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	LocationType string `json:"locationType"`
	Country      string `json:"country,omitempty"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Emoji        string `json:"emoji,omitempty"`
}

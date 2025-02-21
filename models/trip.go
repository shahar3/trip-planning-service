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
	BudgetLevel     string     `json:"budgetLevel"`
	TravelIntensity string     `json:"travelIntensity"`
	Interests       []string   `json:"interests"`
}

type Location struct {
	Value string `json:"value"`
}

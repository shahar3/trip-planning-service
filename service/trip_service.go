package service

import (
	"github.com/shahar3/trip-planning-service/models"
	"github.com/shahar3/trip-planning-service/repository"
)

// TripService defines the business logic interface for trip operations.
type TripService interface {
	CreateTrip(trip *models.Trip) error
	GetTrip(id string) (*models.Trip, error)
	UpdateTrip(trip *models.Trip) error
	DeleteTrip(id string) error
}

// tripServiceImpl is a concrete implementation of TripService.
type tripServiceImpl struct {
	repo repository.TripRepository
}

// NewTripService creates a new TripService instance.
func NewTripService(repo repository.TripRepository) TripService {
	return &tripServiceImpl{
		repo: repo,
	}
}

// CreateTrip creates a new trip. If the planning method is "ai",
// it calls the AI service to generate an itinerary and then updates the trip.
func (s *tripServiceImpl) CreateTrip(trip *models.Trip) error {
	// Persist the trip in the repository.
	if err := s.repo.CreateTrip(trip); err != nil {
		return err
	}

	// If the user chose AI-driven planning, delegate to the AI service.
	//if trip.PlanningMethod == "ai" && s.aiClient != nil {
	//	itinerary, err := s.aiClient.GenerateItinerary(trip)
	//	if err != nil {
	//		return fmt.Errorf("failed to generate itinerary: %w", err)
	//	}
	//	trip.Itinerary = itinerary
	//
	//	// Update the trip with the generated itinerary.
	//	if err := s.repo.UpdateTrip(trip); err != nil {
	//		return fmt.Errorf("failed to update trip with itinerary: %w", err)
	//	}
	//}

	return nil
}

// GetTrip retrieves a trip by its ID.
func (s *tripServiceImpl) GetTrip(id string) (*models.Trip, error) {
	return s.repo.GetTrip(id)
}

// UpdateTrip updates an existing trip.
func (s *tripServiceImpl) UpdateTrip(trip *models.Trip) error {
	return s.repo.UpdateTrip(trip)
}

// DeleteTrip deletes a trip by its ID.
func (s *tripServiceImpl) DeleteTrip(id string) error {
	return s.repo.DeleteTrip(id)
}

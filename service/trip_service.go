package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/shahar3/trip-planning-service/config"
	"github.com/shahar3/trip-planning-service/constants"
	"github.com/shahar3/trip-planning-service/models"
	"github.com/shahar3/trip-planning-service/pkg/kafka"
	"github.com/shahar3/trip-planning-service/repository"
)

// TripService defines the business logic interface for trip operations.
type TripService interface {
	CreateTrip(trip *models.PlanningForm) error
	GetTrip(id string) (*models.Trip, error)
	UpdateTrip(trip *models.Trip) error
	DeleteTrip(id string) error
}

// tripServiceImpl is a concrete implementation of TripService.
type tripServiceImpl struct {
	repo        repository.TripRepository
	kafkaClient *kafka.Client
}

// NewTripService creates a new TripService instance.
func NewTripService(repo repository.TripRepository, cfg *config.Config) TripService {
	kafkaInstance := kafka.NewKafkaClient([]string{cfg.Kafka.Broker}, constants.TripPlanningKafkaTopic)
	return &tripServiceImpl{
		repo:        repo,
		kafkaClient: kafkaInstance,
	}
}

// CreateTrip creates a new trip. If the planning method is "ai",
// it calls the AI service to generate an itinerary and then updates the trip.
func (s *tripServiceImpl) CreateTrip(form *models.PlanningForm) error {
	if form.PlanningMethod == constants.PlanningMethodAi {
		// Marshal the form to JSON and send to the AI service
		payload, err := json.Marshal(form)
		if err != nil {
			return fmt.Errorf("failed to marshal planning form: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := s.kafkaClient.SendMessage(ctx, []byte(constants.TripPlanningKafkaTopic), payload); err != nil {
			return fmt.Errorf("failed to send message to AI service: %w", err)
		}

		return nil
	} else if form.PlanningMethod == constants.PlanningMethodManual {
		return nil
	}

	return fmt.Errorf("unknown planning method: %s", form.PlanningMethod)
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

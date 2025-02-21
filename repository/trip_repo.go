package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/shahar3/trip-planning-service/models"
)

var ErrTripNotFound = errors.New("trip not found")

// TripRepository defines the interface for storing trips.
type TripRepository interface {
	CreateTrip(trip *models.Trip) error
	GetTrip(id string) (*models.Trip, error)
	UpdateTrip(trip *models.Trip) error
	DeleteTrip(id string) error
}

// InMemoryTripRepository is a simple in-memory implementation.
type InMemoryTripRepository struct {
	data map[string]*models.Trip
	mu   sync.RWMutex
}

func NewInMemoryTripRepository() *InMemoryTripRepository {
	return &InMemoryTripRepository{
		data: make(map[string]*models.Trip),
	}
}

func (r *InMemoryTripRepository) CreateTrip(trip *models.Trip) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	trip.ID = uuid.New().String()
	now := time.Now()
	trip.CreatedAt = now
	trip.UpdatedAt = now
	r.data[trip.ID] = trip
	return nil
}

func (r *InMemoryTripRepository) GetTrip(id string) (*models.Trip, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	trip, exists := r.data[id]
	if !exists {
		return nil, ErrTripNotFound
	}
	return trip, nil
}

func (r *InMemoryTripRepository) UpdateTrip(trip *models.Trip) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.data[trip.ID]
	if !exists {
		return ErrTripNotFound
	}
	trip.UpdatedAt = time.Now()
	r.data[trip.ID] = trip
	return nil
}

func (r *InMemoryTripRepository) DeleteTrip(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.data[id]
	if !exists {
		return ErrTripNotFound
	}
	delete(r.data, id)
	return nil
}

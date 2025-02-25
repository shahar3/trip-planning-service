package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shahar3/trip-planning-service/models"
	"github.com/shahar3/trip-planning-service/service"
)

type TripHandler struct {
	service service.TripService
}

func NewTripHandler(service service.TripService) *TripHandler {
	return &TripHandler{service: service}
}

// CreateTrip handles POST /trip requests.
func (h *TripHandler) CreateTrip(c *gin.Context) {
	var planningForm models.PlanningForm
	if err := c.ShouldBindJSON(&planningForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateTrip(&planningForm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, planningForm)
}

// GetTrip handles GET /trip/:id requests.
func (h *TripHandler) GetTrip(c *gin.Context) {
	//id := c.Param("id")
	//trip, err := h.service.GetTrip(id)
	//if err != nil {
	//	if err == repository.ErrTripNotFound {
	//		c.JSON(http.StatusNotFound, gin.H{"error": "Trip not found"})
	//		return
	//	}
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//c.JSON(http.StatusOK, trip)
}

// UpdateTrip handles PUT /trip/:id requests.
func (h *TripHandler) UpdateTrip(c *gin.Context) {
	//id := c.Param("id")
	//var trip models.Trip
	//if err := c.ShouldBindJSON(&trip); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//trip.ID = id // ensure ID from URL is used
	//if err := h.Repo.UpdateTrip(&trip); err != nil {
	//	if err == repository.ErrTripNotFound {
	//		c.JSON(http.StatusNotFound, gin.H{"error": "Trip not found"})
	//		return
	//	}
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//c.JSON(http.StatusOK, trip)
}

// DeleteTrip handles DELETE /trip/:id requests.
func (h *TripHandler) DeleteTrip(c *gin.Context) {
	//id := c.Param("id")
	//if err := h.Repo.DeleteTrip(id); err != nil {
	//	if err == repository.ErrTripNotFound {
	//		c.JSON(http.StatusNotFound, gin.H{"error": "Trip not found"})
	//		return
	//	}
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{"message": "Trip deleted"})
}

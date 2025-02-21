package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shahar3/trip-planning-service/handlers"
)

func SetupRoutes(router *gin.Engine, tripHandler *handlers.TripHandler) {
	api := router.Group("/api")
	{
		api.POST("/trip", tripHandler.CreateTrip)
		api.GET("/trip/:id", tripHandler.GetTrip)
		api.PUT("/trip/:id", tripHandler.UpdateTrip)
		api.DELETE("/trip/:id", tripHandler.DeleteTrip)
	}
}

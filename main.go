package main

import (
	"github.com/fitnis/appointment-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/appointments")
	{
		api.POST("/schedule", handlers.ScheduleAppointment)
		api.GET("/schedule", handlers.GetAppointments)
		api.DELETE("/schedule/:appointmentId", handlers.CancelAppointment)
	}

	router.Run(":8081")
}

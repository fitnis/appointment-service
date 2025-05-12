package main

import (
	"github.com/fitnis/appointment-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/appointments")
	{
		handlers.RegisterAppointmentRoutes(api)
	}

	router.Run(":8081")
}

package handlers

import (
	"net/http"

	"github.com/fitnis/appointment-service/models"
	"github.com/fitnis/appointment-service/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// curl -X POST -H "Content-Type: application/json" -d '{"patientId":"123","date":"2025-04-01","time":"14:00","doctor":"Dr. House"}' http://localhost:8081/appointments/schedule
func ScheduleAppointment(c *gin.Context) {
	var req models.AppointmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	id := uuid.New().String()
	resp := services.ScheduleAppointment(id, req)
	c.JSON(http.StatusCreated, resp)
}

// curl http://localhost:8081/appointments/schedule
func GetAppointments(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetAppointments())
}

// curl -X DELETE http://localhost:8081/appointments/schedule/abc-123
func CancelAppointment(c *gin.Context) {
	id := c.Param("appointmentId")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		return
	}
	services.CancelAppointment(id)
	c.Status(http.StatusNoContent)
}

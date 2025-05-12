package services

import (
	"github.com/fitnis/appointment-service/models"
)

var appointments []models.AppointmentRequest
var appointmentCounter = 0

func ScheduleAppointment(req models.AppointmentRequest) models.GenericResponse {
	appointments = append(appointments, req)
	appointmentCounter++
	return models.GenericResponse{Message: "Appointment scheduled"}
}

func GetAppointments() []models.AppointmentRequest {
	return appointments
}

func CancelAppointment(id string) {
	// Mock delete: just clear the list
	appointments = []models.AppointmentRequest{}
}

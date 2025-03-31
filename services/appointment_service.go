package services

import (
	"github.com/fitnis/appointment-service/models"
)

var appointments = make(map[string]models.AppointmentRequest)

func ScheduleAppointment(id string, req models.AppointmentRequest) models.GenericResponse {
	appointments[id] = req
	return models.GenericResponse{Message: "Appointment scheduled"}
}

func GetAppointments() []models.AppointmentRequest {
	var list []models.AppointmentRequest
	for _, a := range appointments {
		list = append(list, a)
	}
	return list
}

func CancelAppointment(id string) {
	delete(appointments, id)
}

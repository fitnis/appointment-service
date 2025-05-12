package models

type AppointmentRequest struct {
	PatientID string `json:"patientId"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Doctor    string `json:"doctor"`
}
type GenericResponse struct {
	Message string `json:"message"`
}

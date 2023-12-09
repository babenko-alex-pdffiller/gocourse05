package cmd

import (
	"fmt"
)

type MedicalRecord interface {
	GetMedicalHistory() string
}

func PrintMedicalHistory(mr MedicalRecord) {
	fmt.Println(mr.GetMedicalHistory())
}

type PatientRecord struct {
	// ...інші поля
}

func (pr PatientRecord) GetPatientHistory() string {
	return "Patient medical history"
}

// HealthcareService - інтерфейс для охорони здоров'я
type HealthcareService interface {
	RegisterPatient(p clinic.Patient) error
	GetPatientByID(id string) (clinic.Patient, error)
}

func performHealthCheck(hs HealthcareService, patientID string) {
	patient, err := hs.GetPatientByID(patientID)
	if err != nil {
		fmt.Println("Error getting patient:", err)
		return
	}
	fmt.Printf("Performing health check on patient: %s\n", patient.Name)
	// Додаткова логіка перевірки здоров'я
}

func printDetails(i interface{}) {
	switch v := i.(type) {
	case Patient:
		fmt.Printf("Patient Name: %s, Age: %d\n", v.Name, v.Age)
	case *Clinic:
		fmt.Println("Clinic with patients:", len(v.patients))
	default:
		fmt.Println("Unknown type")
	}
}

func printInfo(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

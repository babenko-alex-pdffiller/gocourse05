package funcs

import (
	"fmt"
	"gocourse05/pkg/clinic"
	"gocourse05/pkg/patient"
	"log"
)

type Visitor interface {
	PrintVisits()
}

// HealthcareService - інтерфейс для охорони здоров'я
type HealthcareService interface {
	GetPatientByID(id string) (clinic.Patient, error)
}

func PerformHealthCheck(hs HealthcareService, patientID string) {
	p, err := hs.GetPatientByID(patientID)
	if err != nil {
		fmt.Println("Error getting patient:", err)
		return
	}
	fmt.Printf("Performing health check on patient: %s\n", p.GetName())
	// Додаткова логіка перевірки здоров'я
}

func PrintDetails(objects ...any) {
	for i := range objects {
		switch v := objects[i].(type) {
		case patient.Patient:
			fmt.Printf("Patient Name: %s, Age: %d\n", v.Name, v.Age)
		case *clinic.Clinic:
			fmt.Println("Clinic with patients:", len(v.Patients()))
		default:
			fmt.Println("Unknown type")
		}
	}
}

// ProcessPatient uses contracts
func ProcessPatient(object any) error {
	if visitor, ok := object.(Visitor); !ok {
		return fmt.Errorf(`object doesn't implement Visitor interface, it's %T`, object)
	} else {
		visitor.PrintVisits()
	}

	return nil
}

// ProcessClinic doesn't use contracts
func ProcessClinic(object any) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(fmt.Errorf(`RECOVERED: object doesn't implement Visitor interface: %s`, err))
		}
	}()

	c := object.(clinic.Clinic)
	for _, p := range c.Patients() {
		fmt.Printf("%v\n", p)
	}
}

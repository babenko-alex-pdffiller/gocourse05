package main

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

func performHealthCheck(hs HealthcareService, patientID string) {
	p, err := hs.GetPatientByID(patientID)
	if err != nil {
		fmt.Println("Error getting patient:", err)
		return
	}
	fmt.Printf("Performing health check on patient: %s\n", p.GetName())
	// Додаткова логіка перевірки здоров'я
}

func printDetails(objects ...any) {
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

func processPatient(object any) error {
	if visitor, ok := object.(Visitor); !ok {
		return fmt.Errorf(`object doesn't implement Visitor interface, it's %T`, object)
	} else {
		visitor.PrintVisits()
	}

	return nil
}

func main() {
	p1 := patient.Patient{`100`, `Frank`, 32, `A+`, []string{`Visit 1`, `Visit 2`, `Visit 2`}}
	p2 := patient.VIPPatient{patient.Patient{`101`, `Sam`, 27, `B-`, []string{`Visit 21`, `Visit 22`}}}
	c1 := &clinic.Clinic{}

	// Просте використання інтерфейсу
	if err := c1.RegisterPatient(p1); err != nil {
		log.Fatal(err)
	}
	if err := c1.RegisterPatient(p2); err != nil {
		log.Fatal(err)
	}

	// Поліморфізм з Інтерфейсами
	performHealthCheck(c1, p1.GetId())

	// Type Assertions та Type Switches
	printDetails(p1, p2, c1)

	// Інтерфейси як Контракти
	if err := processPatient(p1); err != nil {
		log.Fatal(err)
	}
	if err := processPatient(p2); err != nil {
		log.Println(err)
	}
}

package clinic

import "fmt"

type Patient interface {
	Id() string
	GetPatientDetails() string
}

// Clinic - структура, що містить мапу пацієнтів
type Clinic struct {
	patients map[string]Patient
}

func (c *Clinic) RegisterPatient(p Patient) error {
	// Логіка реєстрації пацієнта
	c.patients[p.Id()] = p
	return nil
}

func (c *Clinic) GetPatientByID(id string) (Patient, error) {
	// Логіка пошуку пацієнта
	patient, exists := c.patients[id]
	if !exists {
		return nil, fmt.Errorf("patient not found")
	}

	return patient, nil
}

func (c *Clinic) PatientExists(id string) bool {
	_, exists := c.patients[id]
	return exists
}

func (c *Clinic) ProvideHealthcare(ps Patient) {
	// використання інтерфейсу замість конкретного типу з пакету `patient`
}

// independent package

package clinic

import "fmt"

type Patient interface { // prevent Cycle import. The save in "patient" package
	GetId() string
	GetName() string
	GetPatientDetails() string
}

type Clinic struct {
	patients map[string]Patient
}

func (c *Clinic) RegisterPatient(p Patient) error {
	// Логіка реєстрації пацієнта
	if c.patients == nil {
		c.patients = make(map[string]Patient)
	}
	c.patients[p.GetId()] = p
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

func (c *Clinic) Patients() map[string]Patient {
	return c.patients
}

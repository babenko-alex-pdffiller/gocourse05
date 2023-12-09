package patient

type Clinic interface {
	PatientExists(id string) bool
}

// Patient - структура, що представляє пацієнта
type Patient struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	BloodType string `json:"blood_type"`
}

func (p Patient) GetPatientDetails() string {
	return ``
}

func (p Patient) CheckPresence(c Clinic) bool {
	return c.PatientExists(p.Id())
}

func (p Patient) Id() string {
	return p.ID
}

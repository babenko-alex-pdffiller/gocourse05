// independent package

package patient

import (
	"fmt"
	"strings"
)

type Clinic interface { // prevent Cycle import. The save in "clinic" package
	PatientExists(id string) bool
}

type VIPPatient struct {
	Patient
}

func (p VIPPatient) GetPatientDetails() string {
	return fmt.Sprintf(`%s: %s`, `VIP`, p.Patient.GetPatientDetails())
}

func (p VIPPatient) PrintVisits() {
	fmt.Printf("%s visits:\n- %s;", p.Name, strings.Join(p.Visits, ";\n- "))
}

// Patient - структура, що представляє пацієнта
type Patient struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	BloodType string   `json:"blood_type"`
	Visits    []string `json:"visits"`
}

func (p Patient) GetPatientDetails() string {
	return fmt.Sprintf(`%s, %d`, p.Name, p.Age)
}

// CheckPresence shows use of Clinic interface
func (p Patient) CheckPresence(c Clinic) bool {
	return c.PatientExists(p.GetId())
}

func (p Patient) GetId() string {
	return p.ID
}

func (p Patient) GetName() string {
	return p.Name
}

package log

import "gocourse05/pkg/clinic"

var someCity map[string]clinic.Clinic

func init() {
	someCity = make(map[string]clinic.Clinic)
}

func AddClinic(key string, val clinic.Clinic) {
	someCity[key] = val
}

func Clinics() map[string]clinic.Clinic {
	return someCity
}

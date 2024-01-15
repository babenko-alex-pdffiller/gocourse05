package city

import "gocourse05/pkg/clinic"

var SomeCity City

type City struct {
	clinics map[string]clinic.Clinic
}

func init() {
	SomeCity = City{make(map[string]clinic.Clinic)}
}

func (c *City) AddClinic(key string, val clinic.Clinic) {
	c.clinics[key] = val
}

func (c *City) Clinics() map[string]clinic.Clinic {
	return c.clinics
}

package camera

import (
	"fmt"
	"time"
)

type Camera interface {
	TakePhoto() Photo
}

type ZooCamera struct {
	ID int
	CameraLocation
}

func (zc ZooCamera) GetId() int {
	return zc.ID
}

func (zc ZooCamera) TakePhoto() Photo {
	return Photo{}
}

type InfraredCamera struct {
	ZooCamera
}

func (ic InfraredCamera) TakePhoto() Photo {
	return Photo{
		Body: fmt.Sprintf("Infrared Photo #%.3d", ic.ID),
		Time: time.Now(),
		Location: CameraLocation{
			Latitude:  ic.Latitude,
			Longitude: ic.Longitude,
		},
	}
}

func (ic InfraredCamera) GetMetaData() string {
	return fmt.Sprintf("%s: %g, %g", "Infrared", ic.Latitude, ic.Longitude)
}

type FlashCamera struct {
	ZooCamera
}

func (fc FlashCamera) TakePhoto() Photo {
	return Photo{
		Body: fmt.Sprintf("Photo with Flash #%.3d", fc.ID),
		Time: time.Now(),
		Location: CameraLocation{
			Latitude:  fc.Latitude,
			Longitude: fc.Longitude,
		},
	}
}

type CameraLocation struct {
	Latitude  float64
	Longitude float64
}

type Photo struct {
	Body     string
	Time     time.Time
	Location CameraLocation
}

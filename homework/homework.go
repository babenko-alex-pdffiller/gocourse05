package main

import (
	"fmt"
	"gocourse05/core/camera"
	"math/rand/v2"
)

type ID interface {
	GetId() int
}

type Metadata interface {
	GetMetaData() string
}

func main() {
	camerasList := buildCamerasList()

	data := map[int]string{}
	photoID := 1
	for _, zooCamera := range camerasList {
		photoData := ""
		// Check if camera is instance of ID interface
		if IDCamera, ok := zooCamera.(ID); ok {
			photoData += fmt.Sprintf("Camera ID: %d, ", IDCamera.GetId())
		}

		photo := zooCamera.TakePhoto()
		photoData += photo.Body

		if metadataCamera, ok := zooCamera.(Metadata); ok {
			photoData += fmt.Sprintf(", Metadata:  %s", metadataCamera.GetMetaData())
		} else if photoData != "" {
			photoData += fmt.Sprintf(", Metadata is empty")
		}

		if photoData != "" {
			data[photoID] = photoData
			photoID++
		}
	}

	sendDataToServer(data)
}

var sendDataToServer = func(data map[int]string) {
	dataToServer := ""
	for _, v := range data {
		dataToServer += fmt.Sprintf("{%s}, ", v)
	}

	if dataToServer == "" {
		fmt.Println("Sent empty request to server")
	} else {
		fmt.Printf("Sent to server: %s\n", dataToServer)
	}
}

var buildCamerasList = func() []camera.Camera {
	camerasList := []camera.Camera{}
	for i := range [10]int{} {
		if i%2 == 0 {
			camerasList = append(
				camerasList,
				camera.InfraredCamera{
					ZooCamera: camera.ZooCamera{
						ID:             i + 1,
						CameraLocation: camera.CameraLocation{Latitude: rand.Float64(), Longitude: rand.Float64()},
					},
				},
			)
		} else {
			camerasList = append(
				camerasList,
				camera.FlashCamera{
					ZooCamera: camera.ZooCamera{
						ID:             i + 1,
						CameraLocation: camera.CameraLocation{Latitude: rand.Float64(), Longitude: rand.Float64()},
					},
				},
			)
		}
	}

	return camerasList
}

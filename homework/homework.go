package main

import (
	"fmt"
	"gocourse05/core/camera"
	"math/rand/v2"
)

type IDer interface {
	ID() int
}

type Metadater interface {
	MetaData() string
}

func main() {
	cameras := buildCameras()

	data := collectData(&cameras)

	dataToServer := sendDataToServer(data)

	if dataToServer == "" {
		fmt.Println("Nothin to send, empty request")
	}
}

func collectData(cameras *[]camera.Camera) map[int]string {
	data := map[int]string{}
	photoID := 1
	for _, zooCamera := range *cameras {
		photoData := ""
		// Check if camera is instance of IDer interface
		if сameraID, ok := zooCamera.(IDer); ok {
			photoData += fmt.Sprintf("Camera ID: %d, ", сameraID.ID())
		}

		photo := zooCamera.TakePhoto()
		photoData += photo.Body

		if metadataCamera, ok := zooCamera.(Metadater); ok {
			photoData += fmt.Sprintf(", Metadata:  %s", metadataCamera.MetaData())
		} else if photoData != "" {
			photoData += fmt.Sprintf(", Metadata is empty")
		}

		if photoData != "" {
			data[photoID] = photoData
			photoID++
		}
	}

	return data
}

func sendDataToServer(data map[int]string) string {
	dataToServer := ""
	for _, v := range data {
		dataToServer += fmt.Sprintf("{%s}, ", v)
	}

	if dataToServer != "" {
		fmt.Printf("Sent to server: %s\n", dataToServer)
	}

	return dataToServer
}

func buildCameras() []camera.Camera {
	cameras := []camera.Camera{}
	for i := range [10]int{} {
		cameraID := i + 1
		if i%2 == 0 {
			cameras = append(cameras, camera.InfraredCamera{ZooCamera: createZooCamera(cameraID)})
		} else {
			cameras = append(cameras, camera.FlashCamera{ZooCamera: createZooCamera(cameraID)})
		}
	}

	return cameras
}

func createZooCamera(id int) camera.ZooCamera {
	return camera.ZooCamera{
		Identifier:     id,
		CameraLocation: camera.CameraLocation{Latitude: rand.Float64(), Longitude: rand.Float64()},
	}
}

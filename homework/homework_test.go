package main

import (
	"gocourse05/core/camera"
	"testing"
)

func TestCollectData(t *testing.T) {
	// Arrange
	cameras := []camera.Camera{
		camera.InfraredCamera{
			ZooCamera: camera.ZooCamera{
				Identifier:     10,
				CameraLocation: camera.CameraLocation{Latitude: 0.1, Longitude: 0.1},
			},
		},
		camera.FlashCamera{
			ZooCamera: camera.ZooCamera{
				Identifier:     11,
				CameraLocation: camera.CameraLocation{Latitude: 0.1, Longitude: 0.1},
			},
		},
	}
	// Act
	data := collectData(&cameras)
	// Assert
	if len(data) != 2 {
		t.Errorf("Expected 2 cameras, got %d", len(data))
	}
	if data[1] != "Camera ID: 10, Infrared Photo #010, Metadata is empty" {
		t.Errorf("Unexpected data in first item, got %s", data[1])
	}
	if data[2] != "Camera ID: 11, Photo with Flash #011, Metadata is empty" {
		t.Errorf("Unexpected data in second item, got %s", data[2])
	}
}

func TestSendDataToServer(t *testing.T) {
	// Arrange
	data := map[int]string{
		1: "Camera ID: 10, Infrared Photo #010, Metadata is empty",
		2: "Camera ID: 11, Photo with Flash #011, Metadata is empty",
	}
	expectedData := "{Camera ID: 10, Infrared Photo #010, Metadata is empty}, {Camera ID: 11, Photo with Flash #011, Metadata is empty}, "
	// Act
	dataToServer := sendDataToServer(data)
	// Assert
	if dataToServer != expectedData {
		t.Errorf("Unexpected data, got %s", dataToServer)
	}
}

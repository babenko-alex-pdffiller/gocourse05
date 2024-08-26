package main

import (
	"fmt"
	"gocourse05/core/camera"
	"strings"
	"testing"
)

var lastSentData string

func mockSendDataToServer(data map[int]string) {
	dataToServer := ""
	for _, v := range data {
		dataToServer += fmt.Sprintf("{%s}, ", v)
	}

	lastSentData = dataToServer
}

func TestMainFunctionSuccessful(t *testing.T) {
	// Arrange
	lastSentData = ""
	originalSendDataToServer := sendDataToServer
	sendDataToServer = mockSendDataToServer
	defer func() { sendDataToServer = originalSendDataToServer }()

	/// Act
	main()
	// Assert
	if lastSentData == "" {
		t.Error("Expected data to be sent to the server, but got empty string")
	}

	if !strings.Contains(lastSentData, "Camera ID: 1") || !strings.Contains(lastSentData, "Camera ID: 2") {
		t.Error("Expected data to contain camera IDs 1 and 2, but it was not found")
	}
}

func mockBuildCamerasList() []camera.Camera {
	camerasList := []camera.Camera{}

	return camerasList
}

func TestMainFunctionWithUnexpectedFormat(t *testing.T) {
	// Arrange
	originalBuildCamerasList := buildCamerasList
	buildCamerasList = mockBuildCamerasList
	defer func() {
		buildCamerasList = originalBuildCamerasList
		if r := recover(); r != nil {
			t.Errorf("Main function panicked with: %v", r)
		}
	}()

	/// Act
	main()
}

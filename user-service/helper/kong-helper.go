package helper

import (
	"fmt"
	"gopkg.in/resty.v1"
)

func RegisterKong() {
	client := resty.New()

	// -----------> add service
	serviceData := map[string]string{
		"name": "user-service",
		"url":  "http://192.168.103.243:3000",
		"path": "/user-service",
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(serviceData).
		Post("http://localhost:8001/services/")

	if err != nil {
		fmt.Println("Error add service:", err)
		return
	}

	fmt.Println("Response add service:", resp)

	//---------------> add router
	routeData := map[string]interface{}{
		"name":          "user-service",
		"service":       "user-service",
		"paths":         []string{"/user-service"},
		"methods":       []string{"GET", "POST"},
		"protocols":     []string{"http", "https"},
		"hosts":         []string{"192.168.103.243"},
		"strip_path":    true,
		"preserve_host": true,
	}

	respPath, errPath := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(routeData).
		Post("http://localhost:8001/services/user-service/routes")

	if errPath != nil {
		fmt.Println("Error add route:", errPath)
		return
	}

	fmt.Println("Response add route:", respPath)
}

package helper

import (
	"fmt"
	"gopkg.in/resty.v1"
)

func RegisterKong() {
	client := resty.New()

	//------------> add service
	serviceData := map[string]string{
		"name": "order-service",
		"url":  "http://192.168.103.243:3001",
		"path": "/order-service",
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
		"name":          "order-service",
		"service":       "order-service",
		"paths":         []string{"/order-service"},
		"methods":       []string{"GET", "POST"},
		"protocols":     []string{"http", "https"},
		"hosts":         []string{"192.168.103.243"},
		"strip_path":    true,
		"preserve_host": true,
	}

	respPath, errPath := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(routeData).
		Post("http://localhost:8001/services/order-service/routes")

	if errPath != nil {
		fmt.Println("Error add route:", errPath)
		return
	}

	fmt.Println("Response add route:", respPath)
}

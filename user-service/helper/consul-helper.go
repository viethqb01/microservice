package helper

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
)

func RegisterServiceWithConsul() {
	// add client
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	// add service
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "user-service"
	registration.Name = "user-service"
	address := "192.168.103.243"
	registration.Address = address
	registration.Port = 3000
	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/healthcheck", address, 3000)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"
	err = consul.Agent().ServiceRegister(registration)
	if err != nil {
		fmt.Println("Error add consul:", err)
		return
	}
}

func LookupServiceWithConsul(serviceID string) (string, error) {
	config := consulapi.DefaultConfig()
	client, err := consulapi.NewClient(config)
	if err != nil {
		return "", err
	}
	services, err := client.Agent().Services()
	if err != nil {
		return "", err
	}
	srvc := services[serviceID]
	address := srvc.Address
	port := srvc.Port
	return fmt.Sprintf("http://%s:%v", address, port), nil
}

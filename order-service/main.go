package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"gopkg.in/resty.v1"
	"net/http"
	"order-service/helper"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", Hello)
	e.GET("/healthcheck", Healthcheck)
	e.GET("/order/list", OrderList)

	// Register Kong
	helper.RegisterKong()

	// Register consul
	helper.RegisterServiceWithConsul()

	// Start server
	e.Logger.Fatal(e.Start(":3001"))
}

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Welcome to Order Service!")
}

func Healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "Good!")
}

func OrderList(c echo.Context) error {
	type Item struct {
		OrderId string `json:"orderId"`
		Price   int    `json:"price"`
	}

	type User struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
	}

	type Order struct {
		User  User   `json:"user"`
		Items []Item `json:"items"`
	}

	item1 := Item{
		OrderId: "123",
		Price:   100000,
	}

	item2 := Item{
		OrderId: "456",
		Price:   200000,
	}

	order := Order{}

	order.Items = append(order.Items, item1)
	order.Items = append(order.Items, item2)

	// call user service - get user info from user service
	add, _ := helper.LookupServiceWithConsul("user-service")

	client := resty.New()
	res, _ := client.R().
		Get(fmt.Sprintf("%s%s", add, "/user/info"))

	if res.String() != "" {
		err := json.Unmarshal([]byte(res.String()), &order.User)
		if err != nil {
			fmt.Printf("Error unmarshal user info: %+v", err)
			return nil
		}
	}

	return c.JSON(http.StatusOK, order)
}

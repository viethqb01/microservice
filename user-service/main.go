package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"user-service/helper"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", Hello)
	e.GET("/healthcheck", Healthcheck)
	e.GET("/user/info", UserInfo)

	// Register Kong
	helper.RegisterKong()

	// Register consul
	helper.RegisterServiceWithConsul()

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Welcome to User Service!")
}

func Healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "Good!")
}

func UserInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":           1,
		"name":         "Hoang Quoc Bao Viet",
		"email":        "viethqb01@gmail.com",
		"phone_number": "0355755697",
	})
}

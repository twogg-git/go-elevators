package controllers

import (
	"go-elevators/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetElevators(c echo.Context) error {
	result := models.GetElevator()
	println("Call GetElevators...")
	return c.JSON(http.StatusOK, result)
}

func PostElevator(c echo.Context, e models.Elevator) error {
	result := models.PostElevator(e)
	println("Call PostElevators...")
	return c.JSON(http.StatusOK, result)
}

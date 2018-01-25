package controllers

import (
	"go-gorutinesfun/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetEmployees(c echo.Context) error {
	result := models.GetEmployee()
	println("GetEmployees...")
	return c.JSON(http.StatusOK, result)
}

func PostEmployee(c echo.Context, e models.Employee) error {
	result := models.PostEmployee(e)
	println("PostEmployee...")
	return c.JSON(http.StatusOK, result)
}

func PutEmployee(c echo.Context, e models.Employee) error {
	result := models.PutEmployee(e)
	println("PutEmployee...")
	return c.JSON(http.StatusOK, result)
}

func DeleteEmployee(c echo.Context, id int) error {
	result := models.DeleteEmployee(id)
	println("DeleteEmployee...")
	return c.JSON(http.StatusOK, result)
}

package main

import (
	"go-gorutinesfun/controllers"
	"go-gorutinesfun/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/*
https://echo.labstack.com/guide

GET http://localhost:8081/employees
POST http://localhost:8081/employee
{
	"employee_name":"cata",
	"Age":"24",
	"Salary":"750666"
}
PUT http://localhost:8081/175
{
	"employee_name":"cata",
	"Age":"24",
	"Salary":"750666"
}
DELETE http://localhost:8081/employee/177
GET http://localhost:8081/employee/177
*/

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "Welcome mvc echo with mysql app using Golang")
	})

	e.GET("/employees", controllers.GetEmployees)

	e.POST("/employee", func(c echo.Context) error {
		emp := new(models.Employee)
		if err := c.Bind(emp); err != nil {
			return err
		}
		return controllers.PostEmployee(c, *emp)
	})

	e.PUT("/employee", func(c echo.Context) error {
		emp := new(models.Employee)
		if err := c.Bind(emp); err != nil {
			return err
		}
		return controllers.PutEmployee(c, *emp)
	})

	e.DELETE("/employee/:id", func(c echo.Context) error {
		id := c.Param("id")
		empId, _ := strconv.Atoi(id)
		return c.JSON(http.StatusNoContent, controllers.DeleteEmployee(c, empId))
	})

	e.Logger.Fatal(e.Start(":8081"))

}

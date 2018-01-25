package models

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"go-gorutinesfun/db"
)

type Employee struct {
	Id     string `json:"id"`
	Name   string `json:"employee_name"`
	Salary string `json: "employee_salary"`
	Age    string `json : "employee_age"`
}
type Employees struct {
	Employees []Employee `json:"employee"`
}

var con *sql.DB

func GetEmployee() Employees {
	con := db.CreateCon()
	//db.CreateCon()
	sqlStatement := "SELECT id,employee_name, employee_age, employee_salary FROM employee order by id"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := Employees{}

	for rows.Next() {
		employee := Employee{}

		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}
		result.Employees = append(result.Employees, employee)
	}
	con.Close()
	return result

}

func PostEmployee(e Employee) Employee {
	con := db.CreateCon()

	stmt, err := con.Prepare("INSERT INTO employee (employee_name, employee_age, employee_salary)VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(e.Name, e.Age, e.Salary)
	if err != nil {
		fmt.Println(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}

	sqlStatement := "SELECT id,employee_name, employee_age, employee_salary FROM employee WHERE id = ?"
	rows, err := con.Query(sqlStatement, id)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := Employee{}

	for rows.Next() {
		employee := Employee{}

		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Age, &employee.Salary)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}
		result = employee
	}

	con.Close()
	return result

}

func PutEmployee(e Employee) Employee {
	con := db.CreateCon()

	stmt, err := con.Prepare("UPDATE employee SET employee_name=?, employee_age = ?, employee_salary = ? WHERE id = ? ")
	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(e.Name, e.Age, e.Salary, e.Id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.RowsAffected())
	}

	sqlStatement := "SELECT id, employee_name, employee_age, employee_salary FROM employee WHERE id = ?"
	rows, err := con.Query(sqlStatement, e.Id)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := Employee{}

	for rows.Next() {
		employee := Employee{}
		err2 := rows.Scan(&employee.Id, &employee.Name, &employee.Age, &employee.Salary)
		if err2 != nil {
			fmt.Print(err2)
		}
		result = employee
	}

	con.Close()
	return result

}

func DeleteEmployee(id int) string {
	con := db.CreateCon()

	stmt, err := con.Prepare("DELETE FROM employee WHERE id = ?")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Delete id", id)

	res, err := stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.RowsAffected())
	}

	con.Close()
	return "OK"

}

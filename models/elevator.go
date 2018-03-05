package models

import (
	_ "database/sql"
	"fmt"
	"go-elevators/db"
	"strconv"
)

type Elevator struct {
	ElevatorId int    `json:"elevator_id"`
	Name       string `json:"name"`
	MaxSize    string `json: "max_size"`
	Status     bool   `json : "status"`
	UpdateAt   string `json : "updated_at"`
}
type Elevators struct {
	Elevators []Elevator `json:"elevator"`
}

func GetElevator() Elevators {
	con := db.CreateCon()
	sqlStatement := "SELECT elevator_id, name, max_size, status, updated_at FROM elevators"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	}
	defer rows.Close()
	result := Elevators{}

	for rows.Next() {
		elevator := Elevator{}
		err2 := rows.Scan(&elevator.ElevatorId, &elevator.Name, &elevator.MaxSize, &elevator.Status, &elevator.UpdateAt)
		if err2 != nil {
			fmt.Print(err2)
		}
		result.Elevators = append(result.Elevators, elevator)
	}
	con.Close()
	return result

}

func PostElevator(e Elevator) Elevator {
	con := db.CreateCon()

	stmt, err := con.Prepare("INSERT INTO elevators (name, max_size, status) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(e.Name, e.MaxSize, e.Status)
	if err != nil {
		fmt.Println(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}

	e.ElevatorId = strconv.Atoi(id)
	con.Close()

	return e
}

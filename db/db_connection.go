package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_HOST    = "127.0.0.1"
	DB_PORT    = "3300"
	DB_USER    = "root"
	DB_PASS    = "admin"
	DB_NAME    = "go-elevators"
	DB_CHARSET = "utf8"
)

var con *sql.DB

// Create mysql connection
func CreateCon() *sql.DB {
	//"mysql", "<username>:<pw>@tcp(<HOST>:<port>)/<dbname>")
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME+"?charset="+DB_CHARSET)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return db
}

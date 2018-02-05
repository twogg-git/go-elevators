package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_USER    = "root"
	DB_PASS    = "admin"
	DB_NAME    = "gofun"
	DB_CHARSET = "utf8"
)

/*Create mysql connection*/
func CreateCon() *sql.DB {
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@/"+DB_NAME+"?charset="+DB_CHARSET)
	if err != nil {
		fmt.Println(err.Error())
	}
	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return db
}

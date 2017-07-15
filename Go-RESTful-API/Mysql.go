package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:123456@/student?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}

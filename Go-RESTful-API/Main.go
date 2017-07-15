package main

import (
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	var err error
	db,err = sql.Open("mysql", "root:123456@/student")

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}

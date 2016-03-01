package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/ptd72") //pass dns to driver

	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	var str string
	err = db.QueryRow("select Wgraphid from windex").Scan(&str)

	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
	}

	log.Println(str)

}

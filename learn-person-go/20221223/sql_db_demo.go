package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

var (
	id   int
	name string
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_test")
	rows, err := db.Query("select id,name from users where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatalln(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

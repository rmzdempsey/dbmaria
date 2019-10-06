package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:maria@/rmddb")
	if err != nil {
		log.Fatalf("open error %v", err.Error())
	}
	defer db.Close()

	// Connect and check the server version

	results, err := db.Query("SELECT name FROM rmdt1")
	if err != nil {
		log.Fatalf("query error %v", err.Error())
	}
	for results.Next() {
		var name string
		// for each row, scan the result into our tag composite object
		err = results.Scan(&name)
		if err != nil {
			log.Fatalf("result reading error %v", err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf("Found: %v\n", name)
	}

}

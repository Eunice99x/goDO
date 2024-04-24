package db

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func DatabaseConnection(){
	dsn := "root:eunice99x@tcp(localhost:3306)/todos?parseTime=true"

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Check if the connection is successful
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Connected to MySQL database successfully")
}
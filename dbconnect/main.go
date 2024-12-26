package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float32
}

func main() {
	cfg := mysql.Config{
		User:   "mysqlu", //This values were hardcoded to avoid the need of
		Passwd: "pwd",    // exporting variables for each node prior to execute this script
		/*
					User:   os.Getenv("DBUSER"),
			        Passwd: os.Getenv("DBPASS"),
		*/
		Net:    "tcp",
		Addr:   "192.168.1.134:3306",
		DBName: "MUSIC",
	}

	// Get a database Handle
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

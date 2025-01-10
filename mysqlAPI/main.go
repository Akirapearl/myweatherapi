package main

// source: https://go.dev/doc/tutorial/database-access#single_row
import (
	"database/sql"
	"log"
	"net/http"

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
					User:   os.Getenv("DBUSER"), // export DBUSER=username
			        Passwd: os.Getenv("DBPASS"), // export DBPASS=password
		*/
		Net:    "tcp",
		Addr:   "192.168.1.134:3306",
		DBName: "MUSIC",
	}

	r := http.NewServeMux()
	r.HandleFunc("GET /albums", getAlbums)
	log.Print("Starting server on port :8080...")
	log.Fatal(http.ListenAndServe(":8090", r))
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	// Get a database Handle -- Initialize DB connection
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
}

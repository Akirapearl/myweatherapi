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
					User:   os.Getenv("DBUSER"), // export DBUSER=username
			        Passwd: os.Getenv("DBPASS"), // export DBPASS=password
		*/
		Net:    "tcp",
		Addr:   "192.168.1.134:3306",
		DBName: "MUSIC",
	}

	// Get a database Handle -- Initialize DB connection
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	albums, err := getAlbumsByArtist(db, "Linkin Park")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
}

func getAlbumsByArtist(db *sql.DB, name string) ([]Album, error) {
	// Found error on docs: db connection needs to be passed to this function in order to allow the query to be executed
	// otherwise, the db.Query returns undefined.

	// Album slice to hold data from returned rows
	var albums []Album

	rows, err := db.Query("SELECT * FROM Albums where artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("getAlbumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	// Loop through rows, using scan to assign column data to struct fields
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("getAlbumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getAlbumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

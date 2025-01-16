package main

// source: https://go.dev/doc/tutorial/database-access#single_row
import (
	"database/sql"
	"encoding/json"
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
		Net:    "tcp",
		Addr:   "192.168.1.134:3306",
		DBName: "MUSIC",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	r := http.NewServeMux()
	r.HandleFunc("GET /albums", getAlbums(db)) //calls function passing the values for the db connection
	r.HandleFunc("POST /albums/add", addAlbum(db))
	r.HandleFunc("PUT /albums/update", updateAlbum(db))
	log.Print("Starting server on port :8090...")
	log.Fatal(http.ListenAndServe(":8090", r))
}

func getAlbums(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var albums []Album
		rows, err := db.Query("SELECT * FROM Albums")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var alb Album
			if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			albums = append(albums, alb)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if len(albums) > 0 {
			if err := json.NewEncoder(w).Encode(albums); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {
			w.Write([]byte("No albums found"))
		}
	}
}

func addAlbum(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var alb Album
		/*Standardize expected response as a JSON*/
		err := json.NewDecoder(r.Body).Decode(&alb)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result, err := db.Exec("INSERT into Albums VALUES ((select max(ID)+1 from Albums a),?, ?, ?);", alb.Title, alb.Artist, alb.Price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		alb.ID = int(id)
		json.NewEncoder(w).Encode(alb)
	}
}

/*
curl http://localhost:8090/albums/add \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"Title": "The Modern Sound of Betty Carter","Artist": "Betty Carter","Price": 49.99}'

*/

func updateAlbum(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var alb Album
		err := json.NewDecoder(r.Body).Decode(&alb)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = db.Exec("UPDATE Albums SET Title = ? WHERE ID = ?", alb.Title, alb.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

/*
curl http://localhost:8090/albums/update \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"ID" : 9,"Title": "Test"}'

*/

package delete

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float32
}

func DeleteAlbum(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var alb Album
		err := json.NewDecoder(r.Body).Decode(&alb)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = db.Exec("DELETE FROM Albums WHERE ID = ?", alb.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

/*
  -- REQUIRES TO AVOID ESCAPING LINES
	curl -X DELETE http://localhost:8090/albums/delete -H "Content-Type: application/json" -d '{"ID" : 8}'
*/

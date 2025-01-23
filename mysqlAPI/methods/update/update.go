package update

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

func UpdateAlbum(db *sql.DB) http.HandlerFunc {
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

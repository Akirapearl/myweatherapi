package main

import (
	"log"
	"net/http"
)

/*
$ curl -X DELETE http://localhost:8090/albums/delete
called DELETE /albums/delete
*/

func main() {

	r := http.NewServeMux()
	r.HandleFunc("GET /albums", getAlbums)
	r.HandleFunc("POST /albums/add", addAlbum)
	r.HandleFunc("PUT /albums/update", updateAlbum)
	r.HandleFunc("DELETE /albums/delete", deleteAlbum)
	log.Print("Starting server on port :8090...")
	log.Fatal(http.ListenAndServe(":8090", r))

}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	log.Print("GET /albums")
	w.Write([]byte("called GET /albums"))
}

func addAlbum(w http.ResponseWriter, r *http.Request) {
	log.Print("POST /albums/add")
	w.Write([]byte("called POST /albums/add"))
}

func updateAlbum(w http.ResponseWriter, r *http.Request) {
	log.Print("PUT /albums/update")
	w.Write([]byte("called PUT /albums/update"))
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	log.Print("DELETE /albums/delete")
	w.Write([]byte("called DELETE /albums/delete"))
}

/*resource: https://go.dev/doc/tutorial/web-service-gin*/
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// GET
func getAlbums(c *gin.Context) {
	// gin.Context - carries request details, validates and serializes JSON, and more.
	// Context.IndentedJSON to serialize the struct into JSON and add it to the response.
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

// ---------------------------------------------
// POST
func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	fmt.Println("Hello World, this is my API")
	/*
		--  STATEMENT --
		You’ll build an API that provides access to a store selling vintage
		recordings on vinyl. So you’ll need to provide endpoints
		through which a client can get and add albums for users.
	*/

	// Endpoints
	// info: stored in memory, lost each time server is stopped -- upgrade path: DB/CSV file

	// initialize handler
	router := gin.Default()
	// calls GET method for /albums and assigns it to getAlbums function
	router.GET("/albums", getAlbums)
	// in Gin, the colon preceding an item in the path signifies that the item is a path parameter.
	router.GET("/albums/:id", getAlbumByID)
	// calls POST method for /albums and assigns it to postAlbums function
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")

}

/*
POST - suggested entry

curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

GET - byID
curl http://localhost:8080/albums/2

*/

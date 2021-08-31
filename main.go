package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Language string `json:"language"`
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBooks)
	router.DELETE("/books/:id", removeBookByID)

	router.Run("localhost:8080")
}

// books slice to seed record book data.
var books = []book{
	{ID: "1", Title: "Things Fall Apart", Author: "Chinua Achebe", Language: "English"},
	{ID: "2", Title: "Fairy Tales", Author: "Hans Christian Andersen", Language: "Danish"},
	{ID: "3", Title: "The Divine Comedy", Author: "Dante Alighieri", Language: "Italian"},
}

// getBooks responds with the list of all books as a JSON.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// postBooks adds a book from JSON received in request body.
func postBooks(c *gin.Context) {
	var newBook book

	// Call BindJSON to bind the received JSON to newBook
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// add the new book to the slice
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// getBookByID locate the book whose ID value matches the id
// parameter sent by the client, then returns that book as a response.
func getBookByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of books, looking for
	// a book whose ID value matches the parameter
	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func removeBookByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range books {
		if a.ID == id {

			books[i] = books[len(books)-1]
			books = books[:len(books)-1]

			c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
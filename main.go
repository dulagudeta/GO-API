package main

import (
	"github.com/gin-gonic/gin"
	//"net/http"
	"errors"
)

type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []book{
	{ID: "1", Title: "Fikir Eske Mekabir", Author: "Haddis Alemayehu", Price: 12.99},
	{ID: "2", Title: "Oromay", Author: "Bealu Girma", Price: 11.50},
	{ID: "3", Title: "Tobbya", Author: "Yismake Worku", Price: 13.75},
}

func getbooks(c *gin.Context) {
	c.IndentedJSON(200, books)
}

func bookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := getbookByID(id)
	if err != nil {
		return
	}
	c.IndentedJSON(200, book)
}
func getbookByID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func createbook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(201, newBook)
}
func main() {
	router := gin.Default()
	router.GET("/books", getbooks)
	router.GET("/books/:id", bookByID)
	router.POST("/books", createbook)
	router.Run("localhost:8080")
}

package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []Book{
	{ID: "1", Title: "Fikir Eske Mekabir", Author: "Haddis Alemayehu", Price: 12.99},
	{ID: "2", Title: "Oromay", Author: "Bealu Girma", Price: 11.50},
	{ID: "3", Title: "Tobbya", Author: "Yismake Worku", Price: 13.75},
}

//  GET ALL BOOKS 
func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

//  GET BOOK BY ID 
func getBookByIDHandler(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func getBookByID(id string) (*Book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

//  CREATE BOOK 
func createBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	if newBook.ID == "" || newBook.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID and Title are required"})
		return
	}

	for _, b := range books {
		if b.ID == newBook.ID {
			c.JSON(http.StatusConflict, gin.H{"error": "ID already exists"})
			return
		}
	}

	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

//  UPDATE BOOK 
func updateBook(c *gin.Context) {
	id := c.Param("id")

	var updatedData Book
	if err := c.BindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	for i, b := range books {
		if b.ID == id {
			books[i] = updatedData
			c.JSON(http.StatusOK, books[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

//  DELETE BOOK
func deleteBook(c *gin.Context) {
	id := c.Param("id")

	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}
// PATCH--PARTIAL UPDATE
func patchBook(c *gin.Context) {
	id := c.Param("id")

	var patchData map[string]interface{}
	if err := c.BindJSON(&patchData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	for i, b := range books {
		if b.ID == id {

			// Update only sent fields
			if title, ok := patchData["title"].(string); ok {
				books[i].Title = title
			}
			if author, ok := patchData["author"].(string); ok {
				books[i].Author = author
			}
			if price, ok := patchData["price"].(float64); ok {
				books[i].Price = price
			}
			if newID, ok := patchData["id"].(string); ok {
				books[i].ID = newID
			}

			c.JSON(http.StatusOK, books[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}


func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByIDHandler)
	router.POST("/books", createBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)
	router.PATCH("/books/:id", patchBook)


	router.Run("localhost:8080")
}

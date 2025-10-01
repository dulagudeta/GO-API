package main

import "github.com/gin-gonic/gin"

type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []book{
	{ID: "1", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Price: 10.99},
	{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", Price: 8.99},
	{ID: "3", Title: "1984", Author: "George Orwell", Price: 9.99},
}

func getbooks(c *gin.Context) {
	c.IndentedJSON(200, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getbooks)
	router.Run("localhost:8080")
}

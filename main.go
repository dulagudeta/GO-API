package main

import "github.com/gin-gonic/gin"

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

func main() {
	router := gin.Default()
	router.GET("/books", getbooks)
	router.Run("localhost:8080")
}

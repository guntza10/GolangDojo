package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []book{
	{
		ID:     "1",
		Name:   "Harry Potter",
		Author: "J.K. Rowling",
		Price:  15.9,
	},
	{
		ID:     "2",
		Name:   "One Piece",
		Author: "Oda Eiichirō",
		Price:  2.99,
	},
	{
		ID:     "3",
		Name:   "demon slayer",
		Author: "koyoharu gotouge",
		Price:  2.99,
	},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/book/:id", getBookById)
	router.POST("/books", addBook)
	router.PUT("/book/:id", updateBook)
	router.DELETE("/book/:id", deleteBookById)
	router.Run("localhost:8080")
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getBookById(c *gin.Context) {
	paramId := c.Param("id")
	for _, book := range books {
		if book.ID == paramId {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, "data not found")
}

func addBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	var editBook book
	if err := c.BindJSON(&editBook); err != nil {
		return
	}

	paramId := c.Param("id")
	for i := 0; i <= len(books)-1; i++ {
		if books[i].ID == paramId {
			books[i].Name = editBook.Name
			books[i].Author = editBook.Author
			books[i].Price = editBook.Price

			c.JSON(http.StatusOK, books[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, "data not found")
}

func deleteBookById(c *gin.Context) {
	paramId := c.Param("id")

	for i := 0; i < len(books)-1; i++ {
		if books[i].ID == paramId {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, "delete success")

			return
		}
	}

	c.JSON(http.StatusNotFound, "data not found")
}

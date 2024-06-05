package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "Book1", Author: "Author1", Quantity: 3},
	{ID: "2", Title: "Book2", Author: "Author2", Quantity: 4},
	{ID: "3", Title: "Book3", Author: "Author3", Quantity: 5},
}

func getBooks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, books)
}

func createBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	ctx.IndentedJSON(http.StatusCreated, newBook)
}

func getBookById(id string) (*Book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("Book not found with id " + id)
}

func bookById(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := getBookById(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")
	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Please enter a book id"})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not available"})
		return
	}

	book.Quantity -= 1
	ctx.IndentedJSON(http.StatusOK, book)
}

func returnBook(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")
	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Please enter a book id"})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	book.Quantity += 1
	ctx.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:3000")
}

package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"go-api-rest/src/models"
	"go-api-rest/src/repositories/booksR"
)

type APIEnv struct {
	DB *gorm.DB
}

const notFoundText = "there is no book with the informed ID"

func (a *APIEnv) GetBooks(c *gin.Context) {
	books, err := booksR.GetBooks(a.DB)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, books)
}

func (a *APIEnv) GetBookByID(c *gin.Context) {
	id := c.Params.ByName("id")
	book, exists, err := booksR.GetBookByID(id, a.DB)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, notFoundText)
		return
	}

	c.JSON(http.StatusOK, book)
}

func (a *APIEnv) CreateBook(c *gin.Context) {
	book := models.Book{}

	err := c.BindJSON(&book)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := a.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)
}

func (a *APIEnv) DeleteBook(c *gin.Context) {
	id := c.Params.ByName("id")

	_, exists, err := booksR.GetBookByID(id, a.DB)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	if !exists {
		c.JSON(http.StatusNotFound, notFoundText)
	}

	err = booksR.DeleteBook(id, a.DB)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "record deleted successfully")
}

func (a *APIEnv) UpdateBook(c *gin.Context) {
	id := c.Params.ByName("id")

	_, exists, err := booksR.GetBookByID(id, a.DB)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, notFoundText)
		return
	}

	updatedBook := models.Book{}

	err = c.BindJSON(&updatedBook)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := booksR.UpdateBook(a.DB, &updatedBook); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	a.GetBookByID(c)
}

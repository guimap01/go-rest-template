package books

import (
	"go-api-rest/src/database"
	"go-api-rest/src/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	api := &handlers.APIEnv{
		DB: database.GetDB(),
	}
	book := r.Group("/book")

	book.GET("/", api.GetBooks)
	book.GET("/:id", api.GetBookByID)
	book.POST("/", api.CreateBook)
	book.PUT("/:id", api.UpdateBook)
	book.DELETE("/:id", api.DeleteBook)
}

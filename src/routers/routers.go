package routers

import (
	"go-api-rest/src/routers/books"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	books.Routes(r)

	return r
}

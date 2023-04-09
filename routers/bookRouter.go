package routers

import (
	"sql-api-gorm/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)

	router.GET("/books", controllers.GetAllBook)

	router.PUT("/books/:id", controllers.UpdateBook)

	router.GET("/books/:id", controllers.GetBookByID)

	router.DELETE("/books/:id", controllers.DeleteBookById)

	return router
}

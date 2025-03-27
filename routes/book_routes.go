package routes

import (
	"rest_api_books/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoutes(r *gin.Engine) {
	book := r.Group("/api/book")
	{
		book.POST("/", controllers.CreateBook)
		book.GET("/", controllers.ListBooks)
		book.GET("/:id", controllers.GetBook)
		book.PUT("/:id", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "PUT Book"})
		})
		book.DELETE("/:id", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "DELETE Book"})
		})
	}
}

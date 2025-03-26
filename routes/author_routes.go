package routes

import (
	"rest_api_books/controllers"

	"github.com/gin-gonic/gin"
)

func AuthorRoutes(r *gin.Engine) {
	author := r.Group("/api/author")
	{
		author.POST("/", controllers.CreateAuthor)
		author.GET("/", controllers.ListAuthors)
		author.GET("/:id", controllers.GetAuthor)
		author.PUT("/:id", controllers.UpdateAuthor)
		author.DELETE("/:id", controllers.DeleteAuthor)
	}

}

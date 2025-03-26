package routes

import (
	"rest_api_books/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	controllers.InitializeControllers()
	AuthorRoutes(r)
}

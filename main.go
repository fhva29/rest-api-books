package main

import (
	"rest_api_books/config"
	"rest_api_books/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	routes.InitializeRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		return
	}

}

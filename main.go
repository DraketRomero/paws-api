package main

import (
	"github.com/gin-gonic/gin"
	"lexdev_api.com/paws/src/routes"
)

func main() {
	var server = routes.NewRoutes(gin.Default())

	server.RegisterRoutes()

	server.Server.Run("localhost:8153")
}

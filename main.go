package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	df "lexdev_api.com/paws/src/database/infrastructure/out/databasefactory"
	"lexdev_api.com/paws/src/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontro archivo .env")
	}

	var server = routes.NewRoutes(gin.Default())

	df.ConnecToDatabase()

	server.RegisterRoutes()

	port := os.Getenv("BCK_PORT")
	host := fmt.Sprintf(`localhost:%v`, port)
	server.Server.Run(host)
}

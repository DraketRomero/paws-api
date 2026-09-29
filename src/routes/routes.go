package routes

import (
	"github.com/gin-gonic/gin"
	"lexdev_api.com/paws/src/default/infrastructure/out"
)

type Routes struct {
	Server *gin.Engine
}

func NewRoutes(server *gin.Engine) Routes {
	return Routes{
		Server: server,
	}
}

func (routes *Routes) RegisterRoutes() {
	mainServer := routes.Server.Group("/v1")

	mainServer.GET("/", out.DefaultController)
}

package http

import (
	"github.com/gin-gonic/gin"
	"lexdev_api.com/paws/src/apiresponse/domain"
)

func APIResponse[T any](c *gin.Context, statusCode domain.StatusCode, statusMsg string, message string, data []T) {
	var response domain.ApiResponse[T] = domain.NewApiResponse(statusMsg, message, data)

	c.IndentedJSON(int(statusCode), response)
}

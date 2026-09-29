package out

import (
	"github.com/gin-gonic/gin"
	"lexdev_api.com/paws/src/APIResponse/domain"
	apiresponse "lexdev_api.com/paws/src/APIResponse/infrastructure/in/http"
)

func DefaultController(c *gin.Context) {
	apiresponse.APIResponse[any](c, domain.SUCCESS, "success", "Paws-API v1", nil)
}

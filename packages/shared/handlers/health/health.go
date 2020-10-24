package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(appName string) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "Hello from " + appName,
		})
	}
}

package component

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// catch panic
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Message": fmt.Sprintf("Internal Server Error: %v", err),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

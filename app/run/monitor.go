package run

import (
	"gin_demo/config"

	"github.com/gin-gonic/gin"
)

// Start project: gin.run(Port)
func Monitor(r *gin.Engine) {
	r.Run(":" + config.GetInterfacePort())
}

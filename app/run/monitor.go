package run

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Start project: gin.run(Port)
func Monitor(r *gin.Engine) {

	r.Run(":" + os.Getenv("PORT"))
}

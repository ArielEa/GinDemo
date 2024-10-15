package component

import (
	"gin_demo/app/routes"
	"gin_demo/public"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

// autoload
func Autoload() {
	public.SetEnvMode()

	r.Use(routes.CORSComponent())
}

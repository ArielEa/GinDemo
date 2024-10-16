package component

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var environemnt string

var cache_engine *gin.Engine

/**
* Set Cache
 */
func SetCache(r *gin.Engine) {
	environemnt := gin.Mode()

	cache_engine = r

	fmt.Printf("environment is %v \n", environemnt)
}

/**
* 環境を区別する。
 */
func DistinguishEnv() {

}

package public

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type Method struct {
	GET    string
	POST   string
	PUT    string
	DELETE string
}

var (
	method Method

	me sync.RWMutex
)

func InitializeMethod() {
	me.Lock()
	defer me.Unlock()

	method = Method{
		GET:    "GET",
		POST:   "POST",
		PUT:    "PUT",
		DELETE: "DELETE",
	}
}

func GetMethodConf() Method {

	InitializeMethod()

	me.RLock()
	defer me.RUnlock()

	return method
}

func GetMethod(c *gin.Context) string {
	method := c.Request.Method
	return methodToString(method)
}

func methodToString(method string) string {

	methodConfig := GetMethodConf()

	switch method {
	case "GET":
		return methodConfig.GET
	case "POST":
		return methodConfig.POST
	case "PUT":
		return methodConfig.PUT
	case "DELETE":
		return methodConfig.DELETE
	default:
		return "Unsupported request method."
	}
}

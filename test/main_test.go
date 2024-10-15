package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	return r
}

/*
	 *
	 *
		a. 解释：
			1. setupRouter(): 创建并返回一个新的 Gin 路由，用于测试。这样可以避免在 main() 中的路由影响测试代码。
			2. httptest.NewRecorder(): 用于记录 HTTP 响应，模拟真实的 HTTP 响应。
			3. http.NewRequest(): 创建一个新的 HTTP 请求，此处模拟了一个 GET/POST 请求。
			4. router.ServeHTTP(): 使用 Gin 框架处理请求并将结果写入 httptest.NewRecorder()。
			5. assert.Equal(t, ...): 使用 testify 库的 assert 来进行断言，检查状态码和响应是否符合预期。
		b. 测试方法和路径
			go test -v
			go test ./test -v
		c. 结果和返回
			demo(example):
				ok  	gin_demo/test	0.300s
*/
func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello, World!", w.Body.String())
}

// 测试 /ping 路由, 稍微简单的版本
// func TestPingRoute(t *testing.T) {
// 	// 创建路由
// 	router := setupRouter()

// 	// 使用 httptest 创建一个 HTTP 请求
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/ping", nil)

// 	// 使用 Gin 处理请求
// 	router.ServeHTTP(w, req)

// 	// 断言状态码为 200
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	// 断言返回的响应体为 "Hello, World!"
// 	assert.Equal(t, "Hello, World!", w.Body.String())
// }

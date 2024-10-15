package routes

/**
* 需要在这个文件中写明 GET，POST，PUT的请求方式
 */
import (
	"fmt"
	"gin_demo/public"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var r = *gin.Default()

/*
*
　Cross Error
　CORS ERROR を防止するため作ります
*/
func CORSComponent() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

/**
* 通过此任务识别路由
 */
func RequestMorroring(r *gin.Engine) {
	/**
	* 动态处理请求
	* r.Get, r.Post, r.Put など
	 */
	r.Any("/*path", func(c *gin.Context) {

		path := c.Param("path")

		// ip := c.ClientIP()

		public.LogWriter("Received a request for /ping")

		c.JSON(http.StatusOK, gin.H{
			"message": "pag",
			"method":  public.GetMethod(c),
			"path":    path,
			"mode":    gin.Mode(),
			"msg":     "eechanged",
		})
	})
}

func GetRequestAction() {

}

func PostRequestAction() {

}

func PutRequestAction() {

}

func FileRequestAuthAction(c *gin.Context) {
	// 从表单中获取文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Failed to upload file: %v", err)
		return
	}
	defer file.Close()

	// 获取文件的详细信息
	filename := header.Filename
	fmt.Printf("Uploaded File: %s\n", filename)

	// 保存文件到服务器
	err = c.SaveUploadedFile(header, "./"+header.Filename)

	if err != nil {
		log.Fatal(err)
	}

	// 返回成功消息
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded successfully!", filename))
}

func IpRequestCheck(c *gin.Context) string {
	ip := c.ClientIP()

	return string(ip)
}

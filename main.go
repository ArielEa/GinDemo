package main

import (
	"gin_demo/app/component"
	"gin_demo/app/routes"
	"gin_demo/app/run"
	"gin_demo/public"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

/**
* メインファイル。
* こちが入口です。(main)
* ギン　エンジン　(gin engine) を使用する必要がある場合は、ここから渡してください。
* そうしないと、パス（path）が一致しません。
* ぜひご注意ください。
 */
func main() {
	component.Register()

	public.GetDbConfiguration()

	routes.RequestMorroring(r)

	run.Monitor(r)
}

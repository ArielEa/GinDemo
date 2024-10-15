package main

import (
	"gin_demo/app/component"
	"gin_demo/app/routes"
	"gin_demo/app/run"

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
	// componentsについて
	component.Register()

	// インタフェース請求について
	routes.RequestMorroring(r)

	// プロジェクトを起動する
	run.Monitor(r)
}

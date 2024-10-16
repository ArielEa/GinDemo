package component

import (
	"gin_demo/app/routes"
	"gin_demo/config"
	"gin_demo/public"
	"os"

	"github.com/gin-gonic/gin"
)

var register_engine *gin.Engine

// register components
func Register(r *gin.Engine) error {

	register_engine = r

	CorsMonitor()

	os.Setenv("PORT", config.GetInterfacePort())

	// まず、このメソッドを呼び出します。
	ConfigMonitor()

	AutoloadMonitor()

	// DBについて
	DbMonitor()

	return nil
}

func CorsMonitor() {
	register_engine.Use(routes.CorsComponent())
}

/**
* 自動起動の機能
* config 機能を検査した後で、起動します。
 */
func AutoloadMonitor() {
	Autoload()
}

/**
* プロジェクトが初めて起動する時を運行する。
* Mysql entity
 */
func InstallMonitor() {

}

/**
* プロジェクトを監視する。
* データはDBに保存する。
 */
func ListenerMonitor() {

}

/**
* configなど読みます。
 */
func ConfigMonitor() (string, error) {
	configuration := Analyzed()

	return configuration, nil
}

/**
* DB
 */
func DbMonitor() {
	// DBについて、でもこの機能はcomponentsに入れてる予定があります。
	public.GetDbConfiguration()
}

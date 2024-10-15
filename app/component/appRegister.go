package component

import "gin_demo/public"

// register components
func Register() error {
	// まず、このメソッドを呼び出します。
	ConfigMonitor()

	AutoloadMonitor()

	// DBについて
	DbMonitor()

	return nil
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

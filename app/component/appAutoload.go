package component

import (
	"gin_demo/public"
)

// autoload
func Autoload() {
	public.SetEnvMode()
}

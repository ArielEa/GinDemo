package public

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Environment struct {
	GinMode string `json:"gin_mode"`
}

func LoadLocalEnvConfig() (Environment, error) {
	var environment Environment

	file, err := os.Open("./config/.env.json")

	if err != nil {
		return environment, err
	}

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	err = json.Unmarshal(byteValue, &environment)

	if err != nil {
		return environment, err
	}

	return environment, nil
}

func SetEnvMode() {
	env, err := LoadLocalEnvConfig()

	if err != nil {
		log.Fatalf("error loading environment: %v \n", err)
	}

	switch env.GinMode {
	case "debug":
		gin.SetMode(gin.DebugMode)

	case "test":
		gin.SetMode(gin.TestMode)

	case "release":
		gin.SetMode(gin.ReleaseMode)

	default:
		gin.SetMode(gin.DebugMode)
	}
}

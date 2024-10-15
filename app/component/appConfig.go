package component

// analysis global config and setup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var globalConfigFileName string = "config.json"

type basic_configuration struct {
	Port          int  `json:"port"`
	Db            bool `json:"db"`
	Redis         bool `json:"redis"`
	Cache         bool `json:"cache"`
	Log           bool `json:"log"`
	Env           bool `json:"environment_checker"`
	ErrorMonitor  bool `json:"error_warning"`
	ConsoleError  bool `json:"console_error"`
	AutoDbChecker bool `json:"auto_db_checker"`
	Listener      bool `json:"listener"`
	Test          bool `json:"test"`
}

var configuration_data basic_configuration

func Analyzed() string {

	GetConfigData()

	fmt.Println(configuration_data)

	fmt.Printf("Test 1 is : %v \n", configuration_data.Test1)

	return ""
}

func GetConfigData() (string, error) {
	pwdDir, _ := os.Getwd()

	fileDir := pwdDir + "/config/" + globalConfigFileName

	file, err := os.Open(fileDir)

	if err != nil {
		log.Fatalf("ファイルが存在しません。　%v \n", err)
	}

	defer file.Close()

	var configValue []byte

	configValue, jsonErr := ioutil.ReadAll(file)

	if jsonErr != nil {
		log.Fatalf("ファイル内でJSONエラー発生しました。　%v", jsonErr)
	}

	json.Unmarshal(configValue, &configuration_data)

	var configJsonResult interface{}

	if err := json.Unmarshal(configValue, &configJsonResult); err != nil {
		log.Fatalf("JSONの解析に失敗しました。%v", err)
	}

	configJsonString, stringErr := json.Marshal(configJsonResult)

	if stringErr != nil {
		log.Fatalf("JSONのコンパイル失敗しました。 %v", stringErr)
	}

	fmt.Println(string(configJsonString))

	return string(configJsonString), nil
}

package public

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

var fileDir string = "config"

func GetFile() {

}

func AutoFileConfiguration() (string, error) {
	// pc, file, line, ok
	_, file, _, ok := runtime.Caller(1)

	if !ok {
		return "", fmt.Errorf("no caller information")
	}

	fileConfigurationName, _ := GetMatrixConfig(file, string(FileMode))

	return GetConfiguration(fileConfigurationName)
}

func GetConfiguration(configName string) (string, error) {
	dir, err := os.Getwd()

	if err != nil {
		return "パスを読み取れません", fmt.Errorf("パスを読み取れません: %v", err)
	}

	fullDirPath := dir + "/" + fileDir + "/" + configName + ".json"

	file, err := os.Open(fullDirPath)

	if err != nil {
		return "ファイルを読み取れません", fmt.Errorf("ファイルを読み取れません: %v", err)
	}

	defer file.Close()

	jsonValue, err := ioutil.ReadAll(file)

	if err != nil {
		return "ファイル内でエラー発生しました", fmt.Errorf("ファイル内でエラー発生しました: %v", err)
	}

	var result interface{}

	if err := json.Unmarshal(jsonValue, &result); err != nil {
		return "", fmt.Errorf("JSONの解析に失敗しました: %v", err)
	}

	jsonString, err := json.Marshal(result)

	if err != nil {
		return "", fmt.Errorf("JSON コンパイル失敗でした: %v", err)
	}

	return string(jsonString), nil
}

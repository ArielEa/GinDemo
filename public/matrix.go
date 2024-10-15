package public

import (
	"fmt"
	"strings"
)

var file_suffix string = "json"

type FileNames string

type ConfigMode string

var tmp_name string

const (
	FileMode ConfigMode = "file"
)

/*
*
全てのマトリックス(matrix value)の値をここに書きます
*/
const (
	db_file     FileNames = "db"
	env_file    FileNames = ".env"
	config_file FileNames = "config"
	port_file   FileNames = "port"
)

func CheckMatrixConfig(name string) (string, error) {

	switch name {
	case string(db_file):
		return string(db_file), nil

	case string(env_file):
		return string(env_file), nil

	case string(config_file):
		return string(config_file), nil

	case string(port_file):
		return string(port_file), nil

	default:
		return "", fmt.Errorf("no matching filenames for string: %v", name)
	}
}

/*
*
  - この方法を使用して、渡されたパラメーターである「mode」によって、
    使用するメソッド（methods）を区別します。
*/
func GetMatrixConfig(name string, mode string) (string, error) {

	tmp_name = name

	switch mode {
	case string(FileMode):
		return GetFileNames()
	default:
		return "null", nil
	}
}

func GetFileNames() (string, error) {

	filename := tmp_name

	/**
	 * intercept インターセプト
	 * 文字列最後の　”/” 後の情報をインターセプトします
	 */
	lastIndex := strings.LastIndex(filename, "/")

	if lastIndex == -1 {
		return "", fmt.Errorf("")
	}
	result := filename[lastIndex+1:]

	/**
	 * 文字列最後の「・」より前の情報をインターセプトします
	 */
	fileNameIndex := strings.LastIndex(result, ".")

	if fileNameIndex == -1 {
		return "", fmt.Errorf("")
	}

	fileConfigutationName := result[:fileNameIndex]

	checkName, _ := CheckMatrixConfig(fileConfigutationName)

	return checkName, nil
}

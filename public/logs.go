package public

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func GetLogsFile() string {

	mode := gin.Mode()

	filename := mode + ".log"

	return filename
}

func LogWriter(message string) {
	filename := GetLogsFile()

	// 一時的な保管場所として設定します。
	// 具体的なディレクトリーの設定は、configure（・config/ファイル.json など）内で行います。
	filepath := "/logs/" + filename

	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	dirpath := dir + filepath

	file, err := os.OpenFile(dirpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Printf("ファイル %v がないです! ご確認ください！", dirpath)
		return
	}

	defer file.Close()

	log.SetOutput(file)

	log.Println(message)
}

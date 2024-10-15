package public

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type db_configuration struct {
	User      string `json:"user"`
	Password  string `json:"password"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	Db        string `json:"dbname"`
	Charset   string `json:"charset"`
	ParseTime string `json:"parseTime"`
	Loc       string `json:"loc"`
}

var db *gorm.DB

var dbDetailMessage db_configuration

func GetDbConfiguration() {
	dbJson, err := AutoFileConfiguration()

	if err != nil {
		return
	}

	var env map[string]json.RawMessage

	env_err := json.Unmarshal([]byte(dbJson), &env)

	if env_err != nil {
		return
	}

	envJson, exist := env["prod"]

	if !exist {
		return
	}

	detail_err := json.Unmarshal(envJson, &dbDetailMessage)

	if detail_err != nil {
		return
	}

	SqlConnect()
}

func SqlConnect() {
	dsn := dbDetailMessage.User + ":" + dbDetailMessage.Password + "@tcp(" + dbDetailMessage.Host + ":" + dbDetailMessage.Port + ")/" + dbDetailMessage.Db +
		"?charset=" + dbDetailMessage.Charset +
		"&parseTime=" + dbDetailMessage.ParseTime +
		"&loc=" + url.QueryEscape(dbDetailMessage.Loc)

	var err error

	// この後は、gorm.Openに含まれるため、db.closeを使用しないでください。
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Host Connection Failed! Reason : %v", err)
	}

	SqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("Orignal Connection is Lost! %v", err)
	}

	if err := SqlDB.Ping(); err != nil {
		log.Fatalf("Connection Ping Failed! %v", err)
	}

	if err := SqlDB.Close(); err != nil {
		log.Fatalf("Failed to close DB: %v", err)
	}

	fmt.Println("db connection success")
}

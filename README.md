# マイ　プロジェクトー

#### セキュリティのお知らせ
- app_key: 1234
- secret: abcd
- version: v1.0
- sub_code: fee431abhd4772

#### Package Install パッゲージのインストール

##### Init Project プロジェクトの初期化
```
go mod init project_name

go run main.go
```

##### Install gin-gonic/gin moudle
```
go get -u github.com/gin-gonic/gin
```

##### Install Air (Hot Reload) - ホットリロード
###### 1、 git から
```
git clone https://github.com/cosmtrek/air.git

cd air

go build -o project_directly 

cd project

./air
```

###### 2、go を利用してインストールします、airを使ってる。
```
go install github.com/cosmtrek/air@latest

air init 

./air
```

###### 3、config ファイルなど
```tree
注意：　ファイルの名前は、点の後ろの名前は”ael”で設定します。
       aelファイルを解析します。
── config
│   ├── db.json  ## db.ael
│   ├── config.json ## config.ael
│   ├── .env.json ## .env.ael
```
db.json
```json
{
    "prod": {
      "user": "user id",
      "password": "password",
      "host": "host.url",
      "port": "3306",
      "dbname": "dbname",
      "charset": "utf8mb4",
      "parseTime": "True",
      "loc": "Asia/Tokyo"
    },
    "test": {
    },
    "development": {
    }
  }
  
```

.env.json 
```json
{
    "gin_mode": "debug" // debug, prod, test
}  
```

config.json
```json
{
    "port": 1234, // your custom port
    "db": true,
    "redis": false,
    "cache": true,
    "log": true,
    "environment_checker": true,
    "error_warning": true,
    "console_error":false,
    "auto_db_checker": true,
    "listener": false,
    "test": true
}
```

###### 4、ウェブサイトからダウンロードします
###### https://github.com/air-verse/air/releases ページに訪問して適切なバッションを選択し、ダウンロードします

This project is a simple web application built with Go.

## Project Structure

- [README.md](#readme-md)
- [main.go](#main-go)
- [config/](#config)
  - [config.go](#config-go)
  - [settings.go](#settings-go)
- [controllers/](#controllers)
  - [home_controller.go](#home-controller-go)
  - [user_controller.go](#user-controller-go)
- [models/](#models)
  - [user.go](#user-go)
  - [product.go](#product-go)

## File Descriptions

### <a id="readme-md"></a>README.md

プロジェクトのドキュメント、プロジェクトの概要、セットアップ手順、その他の重要な情報が含まれいます。

### <a id="main-go"></a>main.go

The entry point of the application. This file contains the main function that starts the server and initializes the application.

### <a id="config"></a>config/

Contains configuration files.

#### <a id="config-go"></a>config.go

Handles the loading and management of configuration settings.

#### <a id="settings-go"></a>settings.go

Contains specific configuration settings for different environments or services.

### <a id="controllers"></a>controllers/

Handles the application logic for different routes.

#### <a id="home-controller-go"></a>home_controller.go

Manages the logic for the home page and related routes.

#### <a id="user-controller-go"></a>user_controller.go

Manages user-related routes and functionality, such as user registration and login.

### <a id="models"></a>models/

Defines data models representing entities in the application.

#### <a id="user-go"></a>user.go

Defines the `User` model and its associated methods.

#### <a id="product-go"></a>product.go

Defines the `Product` model and its associated methods.

package main

import (
	"fmt"
	"order_demo/config"
	"order_demo/lib/auth"
	"order_demo/model"
	"order_demo/router"
	"order_demo/schema"
	"time"

	_ "order_demo/docs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// @title Order_demo API
// @version 1.0
// @description order_demo example api

// @contact.name Levi
// @contact.url https://github.com/levi840714

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiToken
// @in header
// @name Authorization
// @host localhost:8080
// schemes http
func main() {
	var err error

	//DB connect
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Config.Mysql.User, config.Config.Mysql.Password, config.Config.Mysql.Ip, config.Config.Mysql.Port, config.Config.Mysql.Db)
	model.DB, err = gorm.Open("mysql", connectStr)
	if err != nil {
		panic("DB connection failed!")
	}
	model.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(schema.AllSchema...)
	defer model.DB.Close()

	// jwt setting
	secretKey := config.Config.Jwt.Secret
	tokenLifeTime := time.Duration(config.Config.Jwt.LifeTime) * time.Minute

	auth.Init(secretKey, tokenLifeTime)

	// setup routers
	router := router.SetupRouters()
	router.Run()
}

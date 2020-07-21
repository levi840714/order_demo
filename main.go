package main

import (
	"fmt"
	"order_demo/config"
	"order_demo/model"
	"order_demo/router"
	"order_demo/schema"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	var err error

	//DB connect
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Config.Mysql.User, config.Config.Mysql.Password, config.Config.Mysql.Ip, config.Config.Mysql.Port, config.Config.Mysql.Db)
	model.DB, err = gorm.Open("mysql", connectStr)
	if err != nil {
		panic("DB connect fail!")
	}
	model.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(schema.AllSchema...)
	defer model.DB.Close()

	// setup routers
	router := router.SetupRouters()
	router.Run()
}

package models

import (
	"github.com/gladmo/wechat/settings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {

	username := settings.Get("mysql.username")
	pass := settings.Get("mysql.password")
	host := settings.Get("mysql.host")
	port := settings.Get("mysql.port")
	table := settings.Get("mysql.table")

	dsn := username + ":" + pass + "@tcp(" + host + ":" + port + ")/" + table + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	return db
}

func InitTables() {
	db := Connect().AutoMigrate(&Crawl{})
	db.AutoMigrate(&Text_joke{})
	db.AutoMigrate(&Img_joke{})

	defer db.Close()
}

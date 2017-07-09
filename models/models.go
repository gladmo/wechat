package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {

	dsn := "root:Aa13857026500@tcp(ubuntu-0706.mysql.rds.aliyuncs.com:3306)/jokes?charset=utf8mb4&parseTime=True&loc=Local"
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

package config

import (
	"exchange_app/global"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(127.0.0.1:3306)/inallxin?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if (err != nil) {
		log.Fatal("fail to init database")
	}

	global.Db = db
}
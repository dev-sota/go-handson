package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:password@tcp(127.0.0.1:3306)/go-handson?charset=utf8mb4&parseTime=True&loc=Local"

var db *gorm.DB

func Init() (err error) {
	db, err = gorm.Open(mysql.Open(dsn))
	return
}

func Get() *gorm.DB {
	return db
}

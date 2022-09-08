package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() *gorm.DB {
	d, err := gorm.Open("mysql", "root:root@/new_schema?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return d
}

func GetDB() *gorm.DB {
	return db
}

package model

import (
	"github.com/jinzhu/gorm"

	"quiz/pkg/config"
)

var db *gorm.DB

type User struct {
	ID       uint   `gorm:"primaryKey"json:"id"`
	Username string `gorm:""json:"username"`
	Password string `json:"password"`
}

func init() {
	db = config.Connect()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() {
	db.NewRecord(u)
	if db.Error != nil {
		panic(db.Error)
	}
	db.Create(&u)
	if db.Error != nil {
		panic(db.Error)
	}
}

func GetUserByUsername(username string) *User {
	var user User
	db.Where("Username=?", username).Find(&user)
	return &user
}

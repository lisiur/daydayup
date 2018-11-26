package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lisiur/daydayup/models"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.LoginUser{})
}

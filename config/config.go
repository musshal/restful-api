package config

import (
	"restful-api/structs"

	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3307)/restful_api?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})

	return db
}

package config

import (
	"restful-api/structs"

	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("root:root@tcp(127.0.0.1:3307)/restful_api")

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})

	return db
}

package config

import (
	"github.com/jinzhu/gorm"
	"github.com/yasir16/simpleAPI/structs"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "forge:@tcp(localhost:3306)/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(structs.User{})
	return db
}

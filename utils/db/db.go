package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_NAME = "db_book"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

func ConnectionDB() *gorm.DB {

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		defer db.Close()
		panic("failed to connect database")

	}

	return db

}

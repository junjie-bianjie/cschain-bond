package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dialect = "mysql"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open(dialect, "root:@tcp(localhost:4000)/test?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	if err != nil {
		// handle the error
		panic(err.Error())
	}
	return db
}

package utils

import (
	"cschain-bond/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dialect = "mysql"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open(dialect, "root:@tcp(localhost:4000)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logger.Error("connection tidb failed", logger.String("err", err.Error()))
		panic(err.Error())
	}
	db.SingularTable(true)
	return db
}

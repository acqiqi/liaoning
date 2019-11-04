package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"vgateway/database/mysql"
)

type DBService struct {
	MysqlDb *gorm.DB
}

var DB *DBService

func Open() (db *DBService, err error) {
	fmt.Println("begin database")
	DB = &DBService{}
	DB.MysqlDb, err = mysql.Open()
	if err != nil {
		return
	}
	return
}

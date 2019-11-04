package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"vgateway/common"
)

var DB gorm.DB

func Open() (db *gorm.DB, err error) {
	//打开数据库
	db, err = gorm.Open("mysql",
		common.GetConfig("db_username")+
			":"+common.GetConfig("db_password")+
			"@tcp("+common.GetConfig("db_host")+
			":"+common.GetConfig("db_port")+
			")/"+common.GetConfig("db_name")+
			"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("sqlOk")
	//设置表前缀
	gorm.DefaultTableNameHandler = func(Orm *gorm.DB, defaultTableName string) string {
		return "vhake_" + defaultTableName
	}
	// 禁用复数表名
	db.SingularTable(true)
	//WhereUser()
	//defer Orm.Close()
	DB = *db
	return
}

func init() {
}

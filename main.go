package main

import (
	"log"
	"vgateway/application"
	"vgateway/common"
	"vgateway/database"
	_ "vgateway/kernel"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	//melsecfxserial.FxCalculateWordStartAddress("c100")
	initService()
}

func initService() {
	database.Open()
	common.SetupCache()
	application.InitServer()
}

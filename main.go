package main

import (
	"log"
	"vgateway/application"
	"vgateway/common"
	"vgateway/database"
	_ "vgateway/kernel"
	"vgateway/library/agreement"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	//testAgreement()

	initService()
}

func testAgreement() {
	agreement.Pack.Config.DeviceType = agreement.DeviceTypePLC
	agreement.Setup()
	agreement.Pack.BaseFunc.DriveAddr = 0xffff
	agreement.Pack.BaseFunc.FuncBeginAddr = 1
	data := agreement.Pack.GetDOOnOffCode(true)
	log.Println("begin init")
	log.Println(data)
}

func initService() {
	database.Open()
	common.SetupCache()
	application.InitServer()
}

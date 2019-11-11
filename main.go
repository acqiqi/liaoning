package main

import (
	"log"
	"vgateway/application"
	"vgateway/common"
	"vgateway/database"
	_ "vgateway/kernel"
	"vgateway/kernel/serial"
	"vgateway/library/agreement"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	//testAgreement()
	//melsecfxserial.FxCalculateWordStartAddress("c100")
	var buf = make([]byte, 10)
	buf[0] = 0x30
	buf[1] = 0x30
	log.Println(buf)
	initService()
}

func testAgreement() {
	agreement.Pack.Config.DeviceType = agreement.DeviceTypePLC
	agreement.Setup()
	//agreement.Pack.BaseFunc.DriveAddr = 0xffff
	agreement.Pack.BaseFunc.FuncBeginAddr = 7
	data := agreement.Pack.GetDOOnOffCode(true)
	log.Println("写数据")
	log.Printf("%x", data)
	serial.SerialWrite(data)
}

func initService() {
	database.Open()
	common.SetupCache()
	application.InitServer()
}

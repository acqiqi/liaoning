package main

import (
	"log"
	"vgateway/application"
	"vgateway/application/service/mqtt"
	"vgateway/common"
	"vgateway/database"
	_ "vgateway/kernel"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	//melsecfxserial.FxCalculateWordStartAddress("c100")

	str := "{\"message_id\": \"123qwe\", \"data\": {\"vhake\": \"123\"}}"

	db := new(mqtt.SubscribeData)
	common.JsonDecode(str, &db)
	log.Println(db)

	log.Println(db.Data["vhake"])

	initService()

}

func initService() {
	database.Open()
	common.SetupCache()
	application.InitServer()
}

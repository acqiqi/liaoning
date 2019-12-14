package mqtt

import (
	"log"
	"vgateway/library/agreement"
)

var LibDriverOne = new(agreement.Obj)

// 处理回调信息
func handleSubData(data *SubscribeData) {
	switch data.Type {
	case "config": //配置类型
		if err := data.Config(); err != nil {
			log.Println(err.Error())
		}
		break
	case "all": //全局配置
		if err := data.All(); err != nil {
			log.Println(err.Error())
		}
		break
	case "read": //读数据模式
		if err := data.Read(); err != nil {
			log.Println(err.Error())
		}
		break
	case "write": //写模式
		if err := data.Write(); err != nil {
			log.Println(err.Error())
		}
		break
	case "ota": //OTA
		if err := data.OTA(); err != nil {
			log.Println(err.Error())
		}
		break
	case "rst": //Rst
		if err := data.Rst(); err != nil {
			log.Println(err.Error())
		}
		break
	case "lock": //功能锁
		if err := data.Lock(); err != nil {
			log.Println(err.Error())
		}
		break
	case "lib": //PLC协议配置
		if err := data.Lib(); err != nil {
			log.Println(err.Error())
		}
		break
	case "pt": //透传模式  如果开透传就不允许PLC协议配置
		if err := data.PT(); err != nil {
			log.Println(err.Error())
		}
		break
	default:
		log.Println("Subscribe DataType Error")
	}
}

func Pub() {

}

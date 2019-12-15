package mqtt

import (
	"log"
	"vgateway/common"
)

//发送消息
func Publish(topic string, qos byte, payload string) {
	Mqttclient.Publish(topic, qos, false, payload)
}

//返回错误信息
func PublishCallbackErr(data SubscribeData, code int, msg string) {
	db := new(PublishData)
	db.MessageId = data.MessageId
	db.Code = code
	db.Msg = msg
	str := common.JsonEncode(db)
	Publish(db.Topic, 1, str)
}

//返回成功信息
func PublishCallbackSuccess(data PublishData, msg string) {
	str := common.JsonEncode(data)
	log.Println(str)
	Publish(data.Topic, 1, str)
}

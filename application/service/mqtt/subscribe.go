package mqtt

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
	"log"
	"vgateway/kernel/serial"
	"vgateway/library/agreement/driver/melsecfxserial"
)

var SubscribeList []string

// SubscribeHandler
//var SubscribeHandler mqtt.MessageHandler

// sub消息回调
func SubscribeCallback(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("收到消息")
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	fmt.Println("MsgId--->", msg.MessageID())

	data := msg.Payload()
	str := string(data)

	//cdata,_ := melsecfxserial.BuildWriteBoolPacket(str, true)
	//cdata ,_ := melsecfxserial.BuildWriteWordCommand(str,[]byte{0x31,0x31,0x31,0x31,0x31,0x31,0x31,0x31,0x31,0x31,0x31,0x31})
	cdata, _ := melsecfxserial.BuildReadWordCommand(str, 10)
	log.Println(cdata)

	serial.SerialFlush()      //清空接收区
	serial.SerialWrite(cdata) //写入数据

	callbackdata, err := serial.ReadSerialOneData()
	if err != nil {
		log.Println("ganga")
	} else {
		log.Println(callbackdata)
	}
}

// 新增订阅
func AddSubscribe(subscride string) error {
	//查询订阅是否重复
	for _, v := range SubscribeList {
		if subscride == v {
			return errors.New("重复订阅")
		}
	}

	if Mqttclient.Subscribe(subscride, 2, SubscribeCallback); MqttToken.Wait() && MqttToken.Error() != nil {
		fmt.Println(MqttToken.Error())
		return MqttToken.Error()
	}
	SubscribeList = append(SubscribeList, subscride)
	fmt.Println("addSub----->list:", SubscribeList)
	return nil
}

func RemoveSubscride(subscride string) error {
	if Mqttclient.Subscribe(subscride, 2, SubscribeCallback); MqttToken.Wait() && MqttToken.Error() != nil {
		fmt.Println(MqttToken.Error())
		return MqttToken.Error()
	}
	// 删除订阅List
	for i, v := range SubscribeList {
		if subscride == v {
			SubscribeList = append(SubscribeList[:i], SubscribeList[i+1:]...)
		}
		fmt.Println("unSub----->list:", SubscribeList)
	}
	return nil
}

package mqtt

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
	"log"
	"vgateway/library/agreement"
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

	driver := new(agreement.Obj)                           //实例化工厂
	driver.DriverType = agreement.DriverTypeMelsecFxSerial //使用三菱fx串口
	driver.SerialNo = agreement.SerialNoPort0              //串口0
	if err := driver.Init(); err != nil {
		log.Println("driver init err")
		return
	}
	// 写bool 开关
	if err := driver.WriteBool(str, true); err != nil {
		log.Println("err writebool")
	}

	basr, err := driver.ReadBool(str, 1)
	if err != nil {
		log.Println(err.Error())
		log.Println("mdzz")
	} else {
		log.Println(basr)
		log.Println("ojbk")
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

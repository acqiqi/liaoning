package mqtt

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"vgateway/common"
)

var Mqttclient mqtt.Client
var MqttToken mqtt.Token

func Setup() {
	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	fmt.Println("MqttBegin")

	opts := mqtt.NewClientOptions().
		AddBroker(common.GetConfig("mqtt_server")).
		SetClientID(common.GetConfig("mqtt_client_id"))
	Mqttclient = mqtt.NewClient(opts)

	if MqttToken = Mqttclient.Connect(); MqttToken.Wait() && MqttToken.Error() != nil {
		panic(MqttToken.Error())

	}

	//MqttToken := Mqttclient.Publish("/group/vhake/pub", 0, false, "haha")
	MqttToken.Wait()
	//mqttclient.Publish("/group/vhake/pub", 0, false, "Sey")

	if MqttToken := Mqttclient.Subscribe("/device/pub/vhake", 0, SubscribeCallback); MqttToken.Wait() && MqttToken.Error() != nil {
		fmt.Println("error?", MqttToken.Error())
	}
}

package application

import (
	"vgateway/application/service/mqtt"
	"vgateway/application/service/restful"
)

func InitServer() {
	mqtt.Setup()
	restful.InitRestfulServer()

}

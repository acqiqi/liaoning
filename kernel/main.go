package kernel

import (
	"log"
	_ "vgateway/kernel/config"
	//_ "vgateway/kernel/serial"
)

func init() {
	log.Println("init kernel")
}

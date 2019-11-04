package serial

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
)

var IsSerialOpen bool = false

var Serial *serial.Port

func Setup() {

	c := &serial.Config{
		//Name: "/dev/tty.wchusbserial143220",
		Name: "/dev/ttyS1",

		Baud:     9600,
		Parity:   serial.ParityNone,
		StopBits: serial.Stop1,
		Size:     7}
	var err error
	Serial, err = serial.OpenPort(c)
	if err != nil {
		log.Println(err)
		IsSerialOpen = false
	} else {
		IsSerialOpen = true
	}
}

func SerialWrite(data []byte) {
	n, err := Serial.Write(data)
	if err != nil {
		log.Println(err)
		IsSerialOpen = false
	}
	fmt.Print(n)
}

func SerialRead() (data []byte) {
	buf := make([]byte, 128)
	n, err := Serial.Read(buf)
	if err != nil {
		log.Println(err)
		IsSerialOpen = false
	}
	log.Printf("%q", buf[:n])
	return buf
}

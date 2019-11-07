package serial

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"time"
)

var IsSerialOpen bool = false

var Serial *serial.Port

func init() {

	c := &serial.Config{
		Name: "/dev/ttymxc2",
		//Name: "/dev/ttyS1",

		Baud:     115200,
		Parity:   serial.ParityNone,
		StopBits: serial.Stop1,
		Size:     8}
	var err error
	Serial, err = serial.OpenPort(c)
	if err != nil {
		log.Println(err)
		IsSerialOpen = false
	} else {
		IsSerialOpen = true
	}
	go LoopRead()
}

func SerialWrite(data []byte) {
	n, err := Serial.Write(data)
	if err != nil {
		log.Println(err)
		IsSerialOpen = false
	}
	fmt.Print(n)
}

func LoopRead() {
	for i := 0; i < 100000; i++ {
		time.Sleep(time.Millisecond * 100)
		SerialRead()
	}

}

func SerialRead() (data []byte) {

	buf := make([]byte, 128)
	n, err := Serial.Read(buf)
	if err != nil {
		log.Println(err)
		IsSerialOpen = false
	}
	fmt.Printf("%s", buf[:n])
	//log.Printf("%s", buf[:n])
	return buf
}

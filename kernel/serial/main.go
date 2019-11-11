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
		//Name: "/dev/tty.wchusbserial146420",
		Name: "/dev/ttyS0",

		Baud:     9600,
		Parity:   serial.ParityEven,
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
	fmt.Printf("%x", buf[:n])
	//log.Printf("%s", buf[:n])
	return buf
}

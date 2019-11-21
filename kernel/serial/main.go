package serial

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tarm/serial"
	"log"
	"time"
)

var IsSerialOpen = false

//标记是否要准备读取
var readFlag = false

//
var readDaoda = false

var SerialBuff = make([]byte, 128)
var SerialLeh = 0

var Serial *serial.Port

func init() {

	c := &serial.Config{
		Name: "COM7",
		//Name: "/dev/tty.wchusbserial146420",
		//Name: "/dev/ttyS0",

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

// 一次读取数据
func ReadSerialOneData() (base []byte, err error) {
	SerialFlush() //清空接收区
	//最大等待1秒
	readFlag = true
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 100)
		if readDaoda {
			data := make([]byte, SerialLeh)
			for ii := 0; ii < SerialLeh; ii++ {
				data[ii] = SerialBuff[ii]
			}
			base = data
			readDaoda = false
			return
		}
	}
	return base, errors.New("err not msg")
}

func LoopRead() {
	for true {
		time.Sleep(time.Millisecond * 100)
		SerialRead()
	}

}

func SerialFlush() {
	SerialBuff = make([]byte, 128)
	SerialLeh = 0
	Serial.Flush()
}

func SerialRead() (data []byte, err error) {

	buf := make([]byte, 128)
	n, err := Serial.Read(buf)
	if err != nil {
		log.Println(err)
		IsSerialOpen = false
	} else {
	}
	//判断是否需要读取
	if readFlag {
		if n > 0 {
			SerialBuff = buf
			SerialLeh = n
			readDaoda = true
			readFlag = false
		} else {
			err = errors.New("not serial data")

		}
	}

	//log.Printf("%s", buf[:n])
	return
}

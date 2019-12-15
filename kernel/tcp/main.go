package tcp

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net"
	"time"
)

type ClientShortNetworkLink struct {
	Length     int
	Buffer     []byte
	MaxSize    int
	ServerIp   string
	ServerPort string
	TCPAddr    *net.TCPAddr
	TCPConn    *net.TCPConn
	Reader     *bufio.Reader
	IsOpen     bool
	isReadFlag bool
}

// 初始化
func (this *ClientShortNetworkLink) Init() (err error) {
	if this.IsOpen == true {
		return errors.New("Do not reopen TcpLink")
	}
	this.isReadFlag = false

	this.TCPAddr, err = net.ResolveTCPAddr("tcp", this.ServerIp+":"+this.ServerPort)
	if err != nil {
		log.Println(err)
	}

	this.TCPConn, err = net.DialTCP("tcp", nil, this.TCPAddr)
	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return
	}
	//defer this.TCPConn.Close()
	fmt.Println(this.TCPConn.LocalAddr().String() + " : Client connected!")
	this.IsOpen = true
	go this.onMessageReceived()

	return
}

//写数据
func (this *ClientShortNetworkLink) Write(data []byte) (err error) {
	if this.IsOpen == false {
		return errors.New("not open tcp")
	}
	//fmt.Printf("%x",data)
	fmt.Println()
	log.Println(hex.EncodeToString(data))
	_, err = this.TCPConn.Write(data)
	return
}

func (this *ClientShortNetworkLink) Read() (data []byte, err error) {
	if this.IsOpen == false {
		return nil, errors.New("not open tcp")
	}
	this.Length = 0
	this.isReadFlag = true
	time.Sleep(time.Millisecond * 100)
	//获取数据
	if this.Length > 0 {
		//切片
		d := make([]byte, this.Length)
		for i := 0; i < this.Length; i++ {
			d[i] = this.Buffer[i]
		}
		log.Println(d)
		data = d
		this.isReadFlag = false
		this.Length = 0
		return
	} else {
		this.isReadFlag = false
		this.Length = 0
		return nil, errors.New("not callback data")
	}
}

//读取监听线程
func (this *ClientShortNetworkLink) onMessageReceived() {
	if this.IsOpen == false {
		return //关闭监听通道
	}
	this.Reader = bufio.NewReader(this.TCPConn)
	for {
		//开启读取通道的时候才会读取
		if this.isReadFlag == true {
			buf := make([]byte, 1024)
			n, err := this.Reader.Read(buf)

			if err != nil {
				log.Println(err)
				this.Close()
				return
			} else {
				this.Buffer = buf[:n]
				this.Length = n
			}
		}
	}
}

//关闭网络通道
func (this *ClientShortNetworkLink) Close() {
	//this.TCPConn.CloseRead()
	//this.TCPConn.CloseWrite()
	this.TCPConn.Close()
	this.IsOpen = false
	this.Length = 0
	this.isReadFlag = false
}

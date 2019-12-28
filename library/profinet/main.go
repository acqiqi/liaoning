package profinet

import (
	"errors"
	"fmt"
	"vgateway/kernel/config"
	"vgateway/library/profinet/driver/melsecfxserial"
	"vgateway/library/profinet/driver/siemenss7tcp"
)

const (
	DriverTypeMelsecFxSerial = "MelsecFxSerial"
	DriverTypeSiemensS7Tcp   = "SiemensS7Tcp"
	SerialNoPort0            = 0
	SerialNoPort1            = 1
	SerialNoPort2            = 2
)

//封装接口
type Iprofinet interface {
	InitDriver() (err error)                                         //初始化
	WriteBool(address string, status bool) (err error)               //写bool
	ReadBool(address string, length uint) (status []bool, err error) //读bool
	WriteBytes(address string, value []byte) (err error)             //写bytes
	ReadBytes(address string, length uint) (base []byte, err error)  //读bytes
}

type Obj struct {
	DriverType    string `json:"driver_type"`    //设备类型
	DriverAddress string `json:"driver_address"` //ip 或者从站地址
	DriverPort    string `json:"driver_port"`    //端口号
	SerialNo      string `json:"serial_no"`      //使用串口号 只针对使用串口协议
	IsOpen        bool   `json:"is_open"`        //是否打开
	PlcFlag       int    `json:"plc_flag"`
	Iprofinet
}

var LibDriverOne = new(Obj)

//初始化操作
func init() {
	LibDriverOne.DriverType = config.ConfigObject.Lib.DriverType
	LibDriverOne.DriverAddress = config.ConfigObject.Lib.DriverAddress
	LibDriverOne.DriverPort = config.ConfigObject.Lib.DriverPort
	LibDriverOne.SerialNo = config.ConfigObject.Lib.SerialNo
	LibDriverOne.PlcFlag = config.ConfigObject.Lib.PlcFlag
	if err := LibDriverOne.Init(); err != nil {
		fmt.Println("初次初始化协议驱动失败：" + err.Error())
		return
	}
	if err := LibDriverOne.InitDriver(); err != nil {
		fmt.Println("初次初始化协议驱动失败2：" + err.Error())
		return
	}
}

//初始化
func (this *Obj) Init() (err error) {

	//初始化类型
	switch this.DriverType {
	case DriverTypeMelsecFxSerial:
		driver := new(melsecfxserial.MelsecFxSerial)
		driver.SerialNo = this.SerialNo
		this.Iprofinet = driver
		break
	case DriverTypeSiemensS7Tcp:
		driver := new(siemenss7tcp.SiemensS7Tcp)
		driver.DriverAddress = this.DriverAddress
		driver.DriverPort = this.DriverPort
		driver.PlcFlag = this.PlcFlag
		this.Iprofinet = driver
	default:
		return errors.New("设备类型不正确")
	}

	return
}

package agreement

import (
	"errors"
	"vgateway/library/agreement/driver/melsecfxserial"
)

const (
	DriverTypeMelsecFxSerial = "MelsecFxSerial"

	SerialNoPort0 = 0
	SerialNoPort1 = 1
	SerialNoPort2 = 2
)

//封装接口
type IAgreement interface {
	Begin() (err error)                                              //初始化
	WriteBool(address string, status bool) (err error)               //写bool
	ReadBool(address string, length uint) (status []bool, err error) //读bool
	WriteBytes(address string, value []byte) (err error)             //写bytes
	ReadBytes(address string, length uint) (base []byte, err error)  //读bytes
}

type Obj struct {
	DriverType    string //设备类型
	DriverAddress string //ip 或者从站地址
	DriverPort    string //端口号
	SerialNo      int    //使用串口号 只针对使用串口协议
	IsOpen        bool   //是否打开
	IAgreement
}

//初始化操作
func init() {

}

//初始化
func (this *Obj) Init() (err error) {

	//初始化类型
	switch this.DriverType {
	case DriverTypeMelsecFxSerial:
		driver := new(melsecfxserial.MelsecFxSerial)
		driver.SerialNo = this.SerialNo
		this.IAgreement = driver
		break
	default:
		return errors.New("设备类型不正确")
	}

	return
}

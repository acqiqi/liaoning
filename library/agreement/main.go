package agreement

import "fmt"

type BaseFunc struct {
	DriveAddr     uint16 //设备地址
	FuncCode      uint16 //功能码
	FuncBeginAddr uint32 //寄存器地址 只取一个地址只输入这个
	FuncEndAddr   uint32 //寄存器结束地址
}

type DrivePack struct {
	Config       ConfigPack
	PushData     []byte
	CallbackData []byte
	Obj          interface {
		GetDoReadCode(driveAddr uint16, funcaddr uint32) (callback []byte)              //获取读DO数据协议
		DecodeDoReadCode(uData []byte, uAddress uint32) (callback bool)                 //输入数据查询DO开关量 bool
		GetDOOnOffCode(driveAddr uint16, funcaddr uint32, onOff bool) (callback []byte) //设置DO输出开关
	}
	BaseFunc BaseFunc
}

type ConfigPack struct {
	DeviceType int //设备类型 PLC私有协议 modbus ambit私有协议等
	Drive      int //设备类型如果属于plc私有协议 这里输入协议标准
}

/**
设备类型 PLC 比较特殊
*/
const (
	DeviceTypePLC          = 0
	DeviceTypeMODBUSRTU    = 1
	DeviceTypeMODBUSTCP    = 2
	DeviceTypeMODBUSASCII  = 3
	DeviceTypeAMBITPRODUCT = 4
)

/**
驱动模式 PLC Ambit类型需要
*/
const (
	DriveTypeMitsubishiFX = 0
	DriveTypeMitsubishiQ  = 1
	DriveTypeSiemensS7    = 2
)

func (this *DrivePack) InitPack() {
	if this.Config.DeviceType == DeviceTypePLC {
		this.Obj = new(MitsubishiFX)
	}
}

func (this *DrivePack) GetDoReadCode() (data []byte) {
	data = this.Obj.GetDoReadCode(this.BaseFunc.DriveAddr, this.BaseFunc.FuncBeginAddr)
	fmt.Println(data)
	return
}

func (this *DrivePack) GetDOOnOffCode(onoff bool) (data []byte) {
	data = this.Obj.GetDOOnOffCode(this.BaseFunc.DriveAddr, this.BaseFunc.FuncBeginAddr, onoff)
	return
}

func (this *DrivePack) DecodeDoReadCode(data []byte) (onoff bool) {
	onoff = this.Obj.DecodeDoReadCode(data, this.BaseFunc.FuncBeginAddr)
	return
}

var Pack DrivePack

func Setup() {
	Pack.InitPack()
}

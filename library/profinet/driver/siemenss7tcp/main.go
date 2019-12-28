package siemenss7tcp

import (
	"errors"
	"log"
)

const (
	// 设备型号
	S1200     = 1
	S300      = 2
	S400      = 3
	S1500     = 4
	S200Smart = 5
	S200      = 6
)

type SiemensS7Tcp struct {
	DriverAddress string //ip 或者从站地址
	DriverPort    string //端口号
	PlcFlag       int    //plc类型 1200系列 300系列 400系列 1500系列PLC 200的smart系列 200系统，需要额外配置以太网模块
	WordLength    uint   //单个数据字节的长度，西门子为2，三菱，欧姆龙，modbusTcp就为1，AB PLC无效
}

//初始化
func (this *SiemensS7Tcp) InitDriver() (err error) {
	this.WordLength = 2
	if this.DriverPort == "" {
		this.DriverPort = "102"
	}
	switch this.PlcFlag {
	case S1200:
		plcHead1[21] = 0
		break
	case S300:
		plcHead1[21] = 2
		break
	case S400:
		plcHead1[21] = 3
		plcHead1[17] = 0x00
		break
	case S1500:
		plcHead1[21] = 0
		break
	case S200Smart:
		plcHead1 = plcHead1_200smart
		plcHead2 = plcHead2_200smart
		break
	case S200:
		plcHead1 = plcHead1_200
		plcHead2 = plcHead2_200
	default:
		log.Println("not plc flag")
		plcHead1[18] = 0
		break
	}
	return
}

/// <summary>
/// 写入PLC的一个位，例如"M100.6"，"I100.7"，"Q100.0"，"DB20.100.0"，如果只写了"M100"默认为"M100.0"
/// </summary>
/// <param name="address">起始地址，格式为"M100.6",  "I100.7",  "Q100.0",  "DB20.100.0"</param>
/// <param name="value">写入的数据，True或是False</param>
/// <returns>是否写入成功的结果对象</returns>
func (this *SiemensS7Tcp) WriteBool(address string, status bool) (err error) {
	cmd, err := BuildWriteBoolCommand(address, status)
	if err != nil {
		return
	}
	netb, err := WriteShortTcpBytes(cmd, this.DriverAddress, this.DriverPort)
	if err != nil {
		return
	}
	lehcallback := len(netb) - 1
	log.Println(netb)
	if netb[lehcallback] != 0xff {
		return errors.New("not ok write")
	}

	return
}

func (this *SiemensS7Tcp) ReadBool(address string, length uint) (status []bool, err error) {
	return
}
func (this *SiemensS7Tcp) WriteBytes(address string, value []byte) (err error) {
	return
}
func (this *SiemensS7Tcp) ReadBytes(address string, length uint) (base []byte, err error) {
	return
}

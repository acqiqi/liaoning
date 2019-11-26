package siemenss7tcp

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"vgateway/kernel/tcp"
)

func BuildWriteBoolCommand(address string, data bool) (cmd []byte, err error) {
	analysis, err := AnalysisAddress(address)
	if err != nil {
		return
	}
	buffer := [1]byte{}
	if data == true {
		buffer[0] = 0x01
	} else {
		buffer[0] = 0x00
	}

	_PLCCommand := make([]byte, 35+len(buffer))
	_PLCCommand[0] = 0x03
	_PLCCommand[1] = 0x00
	// 长度 -> length
	_PLCCommand[2] = byte((35 + len(buffer)) / 256)
	_PLCCommand[3] = byte((35 + len(buffer)) % 256)
	// 固定 -> fixed
	_PLCCommand[4] = 0x02
	_PLCCommand[5] = 0xF0
	_PLCCommand[6] = 0x80
	_PLCCommand[7] = 0x32
	// 命令 发 -> command to send
	_PLCCommand[8] = 0x01
	// 标识序列号 -> Identification serial Number
	_PLCCommand[9] = 0x00
	_PLCCommand[10] = 0x00
	_PLCCommand[11] = 0x00
	_PLCCommand[12] = 0x01
	// 固定 -> fixed
	_PLCCommand[13] = 0x00
	_PLCCommand[14] = 0x0E
	// 写入长度+4 -> Write Length +4
	_PLCCommand[15] = byte((4 + len(buffer)) / 256)
	_PLCCommand[16] = byte((4 + len(buffer)) % 256)
	// 命令起始符 -> Command start character
	_PLCCommand[17] = 0x05
	// 写入数据块个数 -> Number of data blocks written
	_PLCCommand[18] = 0x01
	_PLCCommand[19] = 0x12
	_PLCCommand[20] = 0x0A
	_PLCCommand[21] = 0x10
	// 写入方式，1是按位，2是按字 -> Write mode, 1 is bitwise, 2 is by word
	_PLCCommand[22] = 0x01
	// 写入数据的个数 -> Number of Write Data
	_PLCCommand[23] = byte(len(buffer) / 256)
	_PLCCommand[24] = byte(len(buffer) % 256)
	// DB块编号，如果访问的是DB块的话 -> DB block number, if you are accessing a DB block
	_PLCCommand[25] = byte(analysis.Content3.(uint16) / 256)
	_PLCCommand[26] = byte(analysis.Content3.(uint16) % 256)

	// 写入数据的类型 -> Types of writing data

	_PLCCommand[27] = byte(analysis.Content1.(int))
	// 偏移位置 -> Offset position
	_PLCCommand[28] = byte(analysis.Content2.(int32) / 256 / 256)
	_PLCCommand[29] = byte(analysis.Content2.(int32) / 256)
	_PLCCommand[30] = byte(analysis.Content2.(int32) % 256)
	// 按位写入 -> Bitwise Write
	_PLCCommand[31] = 0x00
	_PLCCommand[32] = 0x03
	// 按位计算的长度 -> The length of the bitwise calculation
	_PLCCommand[33] = byte(len(buffer) / 256)
	_PLCCommand[34] = byte(len(buffer) % 256)

	_PLCCommand[35] = buffer[0]
	cmd = _PLCCommand
	return
}

/// <summary>
/// 解析数据地址，解析出地址类型，起始地址，DB块的地址
/// </summary>
/// <param name="address">起始地址，例如M100，I0，Q0，DB2.100</param>
/// <returns>解析数据地址，解析出地址类型，起始地址，DB块的地址</returns>
func AnalysisAddress(address string) (result OperateResult, err error) {

	result.Content3 = uint16(0)
	if address[0] == 'I' {
		result.Content1 = 0x81
		result.Content2 = CalculateAddressStarted(address[1:])
	} else if address[0] == 'Q' {
		result.Content1 = 0x82
		result.Content2 = CalculateAddressStarted(address[1:])
	} else if address[0] == 'M' {
		result.Content1 = 0x83
		result.Content2 = CalculateAddressStarted(address[1:])
	} else if address[0] == 'D' || address[0:2] == "DB" {
		result.Content1 = 0x84
		adds := strings.Split(address, ".")
		if address[1] == 'B' {
			b, _ := strconv.ParseInt(adds[0][2:], 10, 16)
			result.Content3 = b
		} else {
			b, _ := strconv.ParseInt(adds[0][1:], 10, 16)
			result.Content3 = b
		}
		result.Content2 = CalculateAddressStarted(address[strings.Index(address, ".")+1 : 0])
	} else if address[0] == 'T' {
		result.Content1 = 0x1d
		result.Content2 = CalculateAddressStarted(address[1:])
	} else if address[0] == 'C' {
		result.Content1 = 0x1c
		result.Content2 = CalculateAddressStarted(address[1:])
	} else if address[0] == 'V' {
		result.Content1 = 0x84
		result.Content2 = CalculateAddressStarted(address[1:])
		result.Content3 = 1
	} else {
		return result, errors.New("NotSupportedDataType")
	}
	return
}

/// <summary>
/// 计算特殊的地址信息
/// </summary>
/// <param name="address">字符串地址</param>
/// <returns>实际值</returns>
func CalculateAddressStarted(address string) (base int32) {
	if strings.Index(address, ".") < 0 {
		b, _ := strconv.ParseInt(address, 10, 32)
		return int32(b * 8)
	} else {
		temp := strings.Split(address, ".")
		b0, _ := strconv.ParseInt(temp[0], 10, 32)
		b1, _ := strconv.ParseInt(temp[1], 10, 32)
		b := (b0 * 8) + b1
		return int32(b)
	}
}

var tcps *tcp.ClientShortNetworkLink = new(tcp.ClientShortNetworkLink)

// tcp短连接通讯发送数据并接收
func WriteShortTcpBytes(data []byte, ip string, port string) (base []byte, err error) {
	//tcps := tcp.ClientShortNetworkLink{}
	//tcps.Close()
	tcps.ServerIp = ip
	tcps.ServerPort = port
	if err := tcps.Init(); err != nil {
		log.Println("init?")
		return nil, err
	}
	//写第一个头文件

	if err := tcps.Write(plcHead1); err != nil {
		tcps.Close()
		return nil, err
	}
	_, err = tcps.Read()
	if err != nil {
		tcps.Close()
		return
	}
	//写第二个头文件
	if err := tcps.Write(plcHead2); err != nil {
		tcps.Close()
		return nil, err
	}
	_, err = tcps.Read()
	if err != nil {
		tcps.Close()
		return
	}
	//写报文
	if err := tcps.Write(data); err != nil {
		tcps.Close()
		return nil, err
	}
	bs, err := tcps.Read()
	if err != nil {
		tcps.Close()
		return
	}

	tcps.Close()
	return bs, nil
}

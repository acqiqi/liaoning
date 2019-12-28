package melsecfxserial

import (
	"bytes"
	"encoding/hex"
	"errors"
	"log"
	"strconv"
)

/// <summary>
/// 生成位写入的数据报文信息，该报文可直接用于发送串口给PLC 主要是bool类型
/// </summary>
/// <param name="address">地址信息，每个地址存在一定的范围，需要谨慎传入数据。举例：M10,S10,X5,Y10,C10,T10</param>
/// <param name="value"><c>True</c>或是<c>False</c></param>
/// <returns>带报文信息的结果对象</returns>
func BuildWriteBoolPacket(address string, value bool) (base []byte, err error) {
	// 初步解析，失败就返回
	analysis, err := FxAnalysisAddress(address)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// 二次运算起始地址偏移量，根据类型的不同，地址的计算方式不同
	startAddress := uint(analysis.Content1.(int64))
	if analysis.MelsecMCDataType == *MelsecMCM {
		if startAddress >= 8000 {
			startAddress = startAddress - 8000 + 0x0F00
		} else {
			startAddress = startAddress + 0x0800
		}
	} else if analysis.MelsecMCDataType == *MelsecMCS {
		startAddress = startAddress + 0x0000
	} else if analysis.MelsecMCDataType == *MelsecMCX {
		startAddress = startAddress + 0x0400
	} else if analysis.MelsecMCDataType == *MelsecMCY {
		startAddress = startAddress + 0x0500
	} else if analysis.MelsecMCDataType == *MelsecMCCS {
		startAddress += startAddress + 0x01C0
	} else if analysis.MelsecMCDataType == *MelsecMCCC {
		startAddress += startAddress + 0x03C0
	} else if analysis.MelsecMCDataType == *MelsecMCCN {
		startAddress += startAddress + 0x0E00
	} else if analysis.MelsecMCDataType == *MelsecMCTS {
		startAddress += startAddress + 0x00C0
	} else if analysis.MelsecMCDataType == *MelsecMCTC {
		startAddress += startAddress + 0x02C0
	} else if analysis.MelsecMCDataType == *MelsecMCTN {
		startAddress += startAddress + 0x0600
	} else {
		//return new OperateResult<byte[]>( StringResources.Language.MelsecCurrentTypeNotSupportedBitOperate );
	}
	var _PLCCommand = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	_PLCCommand[0] = 0x02
	if value {
		_PLCCommand[1] = 0x37
	} else {
		_PLCCommand[1] = 0x38
	}

	_PLCCommand[2] = BuildAsciiBytesFromX4(startAddress)[2] // 偏移地址
	_PLCCommand[3] = BuildAsciiBytesFromX4(startAddress)[3]
	_PLCCommand[4] = BuildAsciiBytesFromX4(startAddress)[0]
	_PLCCommand[5] = BuildAsciiBytesFromX4(startAddress)[1]

	_PLCCommand[6] = 0x03

	crc := FxCalculateCRC(_PLCCommand)
	_PLCCommand[7] = crc[0]
	_PLCCommand[8] = crc[1]

	return _PLCCommand, nil
}

/// <summary>
/// 根据类型地址长度确认需要读取的指令头  主要是bool类型
/// </summary>
/// <param name="address">起始地址</param>
/// <param name="length">bool数组长度</param>
/// <returns>带有成功标志的指令数据</returns>
func BuildReadBoolCommand(address string, length uint) (cmd []byte, c3 int, err error) {
	staraddress, content1, content3, err := FxCalculateBoolStartAddress(address)
	if err != nil {
		//错误处理
		return
	}

	// 计算下实际需要读取的数据长度
	length2 := (uint(content1)+length-1)/8 - (uint(content1) / 8) + 1

	startAddress := staraddress
	var _PLCCommand = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	_PLCCommand[0] = 0x02                                   // STX
	_PLCCommand[1] = 0x30                                   // Read
	_PLCCommand[2] = BuildAsciiBytesFromX4(startAddress)[0] // 偏移地址
	_PLCCommand[3] = BuildAsciiBytesFromX4(startAddress)[1]
	_PLCCommand[4] = BuildAsciiBytesFromX4(startAddress)[2]
	_PLCCommand[5] = BuildAsciiBytesFromX4(startAddress)[3]
	_PLCCommand[6] = BuildAsciiBytesFromX2(byte(length2))[0] // 读取长度
	_PLCCommand[7] = BuildAsciiBytesFromX2(byte(length2))[1]
	_PLCCommand[8] = 0x03              // ETX
	crc := FxCalculateCRC(_PLCCommand) // CRC
	_PLCCommand[9] = crc[0]
	_PLCCommand[10] = crc[1]
	return _PLCCommand, int(content3), nil
}

/// <summary>
/// 根据类型地址长度确认需要读取的指令头 字节型
/// </summary>
/// <param name="address">起始地址</param>
/// <param name="length">长度</param>
/// <returns>带有成功标志的指令数据</returns>
func BuildReadWordCommand(address string, length uint) (base []byte, err error) {
	addressResult, err := FxCalculateWordStartAddress(address)
	if err != nil {
		//错误处理
		return
	}
	length = uint(length * 2)
	startAddress := addressResult.Content1
	var _PLCCommand = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	_PLCCommand[0] = 0x02                                          //报头
	_PLCCommand[1] = 0x45                                          //指令形式
	_PLCCommand[2] = 0x30                                          //Read
	_PLCCommand[3] = 0x30                                          //Read
	_PLCCommand[4] = BuildAsciiBytesFromX4(startAddress.(uint))[0] //起始地址
	_PLCCommand[5] = BuildAsciiBytesFromX4(startAddress.(uint))[1] //起始地址
	_PLCCommand[6] = BuildAsciiBytesFromX4(startAddress.(uint))[2] //起始地址
	_PLCCommand[7] = BuildAsciiBytesFromX4(startAddress.(uint))[3] //起始地址
	_PLCCommand[8] = BuildAsciiBytesFromX2(byte(length))[0]        // 读取长度
	_PLCCommand[9] = BuildAsciiBytesFromX2(byte(length))[1]
	_PLCCommand[10] = 0x03
	//_PLCCommand[0] = 0x02;                                                    // STX
	//_PLCCommand[1] = 0x30;                                                    // Read
	//_PLCCommand[2] = MelsecHelper.BuildBytesFromData( startAddress )[0];      // 偏移地址
	//_PLCCommand[3] = MelsecHelper.BuildBytesFromData( startAddress )[1];
	//_PLCCommand[4] = MelsecHelper.BuildBytesFromData( startAddress )[2];
	//_PLCCommand[5] = MelsecHelper.BuildBytesFromData( startAddress )[3];
	//_PLCCommand[6] = MelsecHelper.BuildBytesFromData( (byte)length )[0];      // 读取长度
	//_PLCCommand[7] = MelsecHelper.BuildBytesFromData( (byte)length )[1];
	//_PLCCommand[8] = 0x03;                                                    // ETX

	crc := FxCalculateCRC(_PLCCommand) // CRC
	_PLCCommand[11] = crc[0]
	_PLCCommand[12] = crc[1]
	base = _PLCCommand
	return // Return
}

/// <summary>
/// 根据类型地址以及需要写入的数据来生成指令头
/// </summary>
/// <param name="address">起始地址</param>
/// <param name="value">实际的数据信息</param>
/// <returns>带有成功标志的指令数据</returns>
func BuildWriteWordCommand(address string, value []byte) (base []byte, err error) {
	addressResult, err := FxCalculateWordStartAddress(address)
	if err != nil {
		//错误处理
		return
	}
	// 字节数据转换成ASCII格式
	if value != nil {
		value = BuildAsciiBytesFromAscci(value)
	}
	startAddress := addressResult.Content1.(uint)
	_PLCCommand := make([]byte, bytes.Count(value, nil)+13)
	_PLCCommand[0] = 0x02                                   //报头
	_PLCCommand[1] = 0x45                                   //指令形式
	_PLCCommand[2] = 0x31                                   //Read
	_PLCCommand[3] = 0x30                                   //Read
	_PLCCommand[4] = BuildAsciiBytesFromX4(startAddress)[0] // Offect Address
	_PLCCommand[5] = BuildAsciiBytesFromX4(startAddress)[1]
	_PLCCommand[6] = BuildAsciiBytesFromX4(startAddress)[2]
	_PLCCommand[7] = BuildAsciiBytesFromX4(startAddress)[3]
	_PLCCommand[8] = BuildAsciiBytesFromX2((byte)(bytes.Count(value, nil) / 2))[0] // Read Length
	_PLCCommand[9] = BuildAsciiBytesFromX2((byte)(bytes.Count(value, nil) / 2))[1]
	for i := 0; i < len(value); i++ {
		_PLCCommand[i+10] = value[i]
	}
	_PLCCommand[9+bytes.Count(value, nil)] = 0x03

	crc := FxCalculateCRC(_PLCCommand) // CRC
	_PLCCommand[9+bytes.Count(value, nil)+1] = crc[0]
	_PLCCommand[9+bytes.Count(value, nil)+2] = crc[1]
	base = _PLCCommand
	return
}

/// <summary>
/// 返回读取的地址及长度信息
/// </summary>
/// <param name="address">读取的地址信息</param>
/// <returns>带起始地址的结果对象</returns>
func FxCalculateWordStartAddress(address string) (or OperateResult, err error) {
	// 初步解析，失败就返回
	analysis, err := FxAnalysisAddress(address)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// 二次解析
	startAddress := uint(analysis.Content1.(int64))
	if analysis.MelsecMCDataType == *MelsecMCD {
		if startAddress >= 8000 {
			startAddress = (startAddress-8000)*2 + 0x0E00
		} else {
			startAddress = startAddress*2 + 0x4000
		}
	} else if analysis.MelsecMCDataType == *MelsecMCCN {
		if startAddress >= 200 {
			startAddress = (startAddress-200)*4 + 0x0C00
		} else {
			startAddress = startAddress*2 + 0x0A00
		}
	} else if analysis.MelsecMCDataType == *MelsecMCTN {
		startAddress = startAddress*2 + 0x0800
	} else {
		return or, errors.New("不匹配")
	}
	log.Println(startAddress)
	or.Content1 = startAddress
	return
}

/// <summary>
/// 返回读取的地址及长度信息，以及当前的偏置信息
/// </summary><param name="address">读取的地址信息</param>
/// <returns>带起始地址的结果对象</returns>
func FxCalculateBoolStartAddress(address string) (staraddr uint, content1 int64, b8 uint, err error) {
	// 初步解析，失败就返回
	analysis, err := FxAnalysisAddress(address)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// 二次解析
	startAddress := uint(analysis.Content1.(int64))

	if analysis.MelsecMCDataType == *MelsecMCM {
		if startAddress >= 8000 {
			startAddress = (startAddress-8000)/8 + 0x01E0
		} else {
			startAddress = startAddress/8 + 0x0100
		}
	} else if analysis.MelsecMCDataType == *MelsecMCX {
		startAddress = startAddress/8 + 0x0080
	} else if analysis.MelsecMCDataType == *MelsecMCY {
		startAddress = startAddress/8 + 0x00A0
	} else if analysis.MelsecMCDataType == *MelsecMCS {
		startAddress = startAddress/8 + 0x0000
	} else if analysis.MelsecMCDataType == *MelsecMCCS {
		startAddress += startAddress/8 + 0x01C0
	} else if analysis.MelsecMCDataType == *MelsecMCCC {
		startAddress += startAddress/8 + 0x03C0
	} else if analysis.MelsecMCDataType == *MelsecMCTS {
		startAddress += startAddress/8 + 0x00C0
	} else if analysis.MelsecMCDataType == *MelsecMCTC {
		startAddress += startAddress/8 + 0x02C0
	} else {
		return staraddr, content1, b8, errors.New("地址错误")
	}

	return startAddress, analysis.Content1.(int64), uint(analysis.Content1.(int64) % 8), nil
}

/// <summary>
/// 解析数据地址成不同的三菱地址类型
/// </summary>
/// <param name="address">数据地址</param>
/// <returns>地址结果对象</returns>
func FxAnalysisAddress(address string) (or *OperateResult, err error) {

	result := new(OperateResult)

	switch address[0] {
	case 'M', 'm':
		result.MelsecMCDataType = *MelsecMCM
		if result.Content1, err = strconv.ParseInt(address[1:], MelsecMCM.FromBase, 16); err != nil {
			return
		}
		break
	case 'X', 'x':
		result.MelsecMCDataType = *MelsecMCX
		if result.Content1, err = strconv.ParseInt(address[1:], 8, 16); err != nil {
			return
		}
		break
	case 'Y', 'y':
		result.MelsecMCDataType = *MelsecMCY
		if result.Content1, err = strconv.ParseInt(address[1:], 8, 16); err != nil {
			return
		}
		break
	case 'D', 'd':
		result.MelsecMCDataType = *MelsecMCD
		if result.Content1, err = strconv.ParseInt(address[1:], MelsecMCD.FromBase, 16); err != nil {
			return
		}
		break
	case 'S', 's':
		result.MelsecMCDataType = *MelsecMCS
		if result.Content1, err = strconv.ParseInt(address[1:], MelsecMCS.FromBase, 16); err != nil {
			return
		}
		break
	case 'T', 't':
		if address[1] == 'N' || address[1] == 'n' {
			result.MelsecMCDataType = *MelsecMCTN
			if result.Content1, err = strconv.ParseInt(address[2:], MelsecMCTN.FromBase, 16); err != nil {
				return
			}
		} else if address[1] == 'S' || address[1] == 's' {
			result.MelsecMCDataType = *MelsecMCTS
			if result.Content1, err = strconv.ParseInt(address[2:], MelsecMCTS.FromBase, 16); err != nil {
				return
			}
		} else if address[1] == 'C' || address[1] == 'c' {
			result.MelsecMCDataType = *MelsecMCTC
			if result.Content1, err = strconv.ParseInt(address[2:], MelsecMCTC.FromBase, 16); err != nil {
				return
			}
		} else {
			return result, errors.New("格式不正确")
		}
		break
	case 'C', 'c':
		if address[1] == 'N' || address[1] == 'n' {
			result.MelsecMCDataType = *MelsecMCCN
			if result.Content1, err = strconv.ParseInt(address[2:], MelsecMCCN.FromBase, 16); err != nil {
				return
			}
		} else if address[1] == 'S' || address[1] == 's' {
			result.MelsecMCDataType = *MelsecMCCS
			if result.Content1, err = strconv.ParseInt(address[2:], MelsecMCCS.FromBase, 16); err != nil {
				return
			}
		} else if address[1] == 'C' || address[1] == 'c' {
			result.MelsecMCDataType = *MelsecMCCC
			if result.Content1, err = strconv.ParseInt(address[2:], MelsecMCCC.FromBase, 16); err != nil {
				return
			}
		} else {
			return result, errors.New("格式不正确")
		}
		break
	default:
		return result, errors.New("格式不正确")
	}
	//处理成功
	return result, nil
}

/// <summary>
/// 检查PLC返回的写入的数据是否是正常的
/// </summary>
/// <param name="ack">Plc反馈的数据信息</param>
/// <returns>检查结果</returns>
func CheckPlcWriteResponse(ack []byte) (err error) {
	if len(ack) == 0 {
		err = errors.New("数据为空")
		return
	}
	if ack[0] == 0x15 {
		err = errors.New("命令执行失败")
	}
	if ack[0] != 0x06 {
		err = errors.New("返回失败")
	}
	return
}

/// <summary>
/// 检查PLC返回的读取数据是否是正常的
/// </summary>
/// <param name="ack">Plc反馈的数据信息</param>
/// <returns>检查结果</returns>
func CheckPlcReadResponse(ack []byte) (err error) {
	if len(ack) == 0 {
		err = errors.New("数据为空")
		return
	}
	if ack[0] == 0x15 {
		err = errors.New("命令执行失败")
	}
	if ack[0] != 0x02 {
		err = errors.New("返回失败")
	}

	if !CheckCRC(ack) {
		err = errors.New("校验失败")
	}
	return
}

/// <summary>
/// 从PLC反馈的数据进行提炼操作
/// </summary>
/// <param name="response">PLC反馈的真实数据</param>
/// <returns>数据提炼后的真实数据</returns>
func ExtractActualData(response []byte) (base []byte) {

	data := make([]byte, (len(response)-4)/2)
	for i := 0; i < len(data); i++ {
		buffer := []byte{0x00, 0x00}
		buffer[0] = response[i*2+1]
		buffer[1] = response[i*2+2]

		bf, err := hex.DecodeString(string(buffer))
		if err == nil {
			data[i] = bf[0]
		}
	}
	return data

}

/// <summary>
/// 从PLC反馈的数据进行提炼bool数组操作
/// </summary>
/// <param name="response">PLC反馈的真实数据</param>
/// <param name="start">起始提取的点信息</param>
/// <param name="length">bool数组的长度</param>
/// <returns>数据提炼后的真实数据</returns>
func ExtractActualBoolData(response []byte, start int, length uint) (base []bool) {
	extraResult := ExtractActualData(response)

	// 转化bool数组

	data := make([]bool, length)
	array := ByteToBoolArray(extraResult, len(response)*8)
	for i := 0; i < int(length); i++ {
		data[i] = array[i+start]
	}
	base = data
	return
}

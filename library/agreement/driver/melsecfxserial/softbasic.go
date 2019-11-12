package melsecfxserial

import (
	"fmt"
	"log"
)

/// <summary>
/// 从ushort构建一个ASCII格式的数据内容
/// </summary>
/// <param name="value">数据</param>
/// <returns>ASCII格式的字节数组</returns>
func BuildAsciiBytesFromX2(value byte) (ret []byte) {
	s2x := fmt.Sprintf("%02X", value)
	b2x := []byte(s2x)
	return b2x
}

/// <summary>
/// 从byte构建一个ASCII格式的数据内容
/// </summary>
/// <param name="value">数据</param>
/// <returns>ASCII格式的字节数组</returns>
func BuildAsciiBytesFromX4(value uint) (ret []byte) {
	ss := fmt.Sprintf("%04X", value)
	ret = []byte(ss)
	return
}

/// <summary>
/// 从字节数组构建一个ASCII格式的数据内容
/// </summary>
/// <param name="value">字节信息</param>
/// <returns>ASCII格式的地址</returns>
func BuildAsciiBytesFromAscci(value []byte) (base []byte) {
	buffer := make([]byte, len(value)*2)
	for i := 0; i < len(value); i++ {
		buffer[i*2] = BuildAsciiBytesFromX2(value[i])[0]
		buffer[i*2+1] = BuildAsciiBytesFromX2(value[i])[1]
	}
	base = buffer
	return
}

/// <summary>
/// 计算Fx协议指令的和校验信息
/// </summary>
/// <param name="data">字节数据</param>
/// <returns>校验之后的数据</returns>
func FxCalculateCRC(data []byte) (ret []byte) {
	sum := uint(0)
	for i := 1; i < len(data)-2; i++ {
		sum = sum + uint(data[i])
	}
	return BuildAsciiBytesFromX2(byte(sum))
}

/// <summary>
/// 检查指定的和校验是否是正确的
/// </summary>
/// <param name="data">字节数据</param>
/// <returns>是否成功</returns>
func CheckCRC(data []byte) (isCheck bool) {
	crc := FxCalculateCRC(data)
	log.Println(crc)
	if crc[0] != data[len(data)-2] {
		return false
	}
	if crc[1] != data[len(data)-1] {
		return false
	}
	return true
}

/// <summary>
/// 从Byte数组中提取位数组，length代表位数 ->
/// Extracts a bit array from a byte array, length represents the number of digits
/// </summary>
/// <param name="InBytes">原先的字节数组</param>
/// <param name="length">想要转换的长度，如果超出自动会缩小到数组最大长度</param>
/// <returns>转换后的bool数组</returns>
/// <example>
/// <code lang="cs" source="HslCommunication_Net45.Test\Documentation\Samples\BasicFramework\SoftBasicExample.cs" region="ByteToBoolArray" title="ByteToBoolArray示例" />
/// </example>
func ByteToBoolArray(InBytes []byte, length int) (base []bool) {
	if InBytes == nil {
		return nil
	}

	if length > (len(InBytes) * 8) {
		length = len(InBytes) * 8
	}
	buffer := make([]bool, length)

	for i := 0; i < length; i++ {
		index := i / 8
		offect := i % 8

		temp := byte(0)
		switch offect {
		case 0:
			temp = 0x01
			break
		case 1:
			temp = 0x02
			break
		case 2:
			temp = 0x04
			break
		case 3:
			temp = 0x08
			break
		case 4:
			temp = 0x10
			break
		case 5:
			temp = 0x20
			break
		case 6:
			temp = 0x40
			break
		case 7:
			temp = 0x80
			break
		default:
			break
		}

		if (InBytes[index] & temp) == temp {
			buffer[i] = true
		}
	}
	return buffer
}

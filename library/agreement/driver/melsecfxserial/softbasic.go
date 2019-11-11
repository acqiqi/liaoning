package melsecfxserial

import (
	"bytes"
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
	for i := 1; i < bytes.Count(data, nil)-2; i++ {
		sum = sum + uint(data[i])
	}
	log.Println(sum)
	return BuildAsciiBytesFromX2(byte(sum))
}

/// <summary>
/// 检查指定的和校验是否是正确的
/// </summary>
/// <param name="data">字节数据</param>
/// <returns>是否成功</returns>
func CheckCRC(data []byte) (isCheck bool) {
	crc := FxCalculateCRC(data)
	if crc[0] != data[bytes.Count(data, nil)-2] {
		return false
	}
	if crc[1] != data[bytes.Count(data, nil)-1] {
		return false
	}
	return true
}

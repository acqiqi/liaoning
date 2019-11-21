package siemenss7tcp

import (
	"errors"
	"strconv"
	"strings"
)

/// <summary>
/// 解析数据地址，解析出地址类型，起始地址，DB块的地址
/// </summary>
/// <param name="address">起始地址，例如M100，I0，Q0，DB2.100</param>
/// <returns>解析数据地址，解析出地址类型，起始地址，DB块的地址</returns>
func AnalysisAddress(address string) (result OperateResult, err error) {
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

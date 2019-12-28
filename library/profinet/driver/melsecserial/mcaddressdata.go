package melsecserial

import "strconv"

/// <summary>
/// 三菱的数据地址表示形式
/// </summary>
type McAddressData struct {
	McDataType            MelsecMCDataType
	DeviceAddressDataBase DeviceAddressDataBase
}

/// <summary>
/// 从指定的地址信息解析成真正的设备地址信息，默认是三菱的地址
/// </summary>
/// <param name="address">地址信息</param>
/// <param name="length">数据长度</param>
func (this *McAddressData) Parse(address string, length uint) {

}

/// 从实际三菱的地址里面解析出
/// </summary>
/// <param name="address">三菱的地址数据信息</param>
/// <param name="length">读取的数据长度</param>
func (this *McAddressData) ParseMelsecFrom(address string, length uint) {

	this.DeviceAddressDataBase.Length = length
	switch address[0] {
	case 'M', 'm':
		this.McDataType = *MelsecMCX
		strconv.ParseInt(address[0:1], MelsecMCX.FromBase, 32)

	}

}

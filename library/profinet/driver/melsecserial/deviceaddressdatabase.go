package melsecserial

// 设备地址数据的信息，通常包含起始地址，数据类型，长度
type DeviceAddressDataBase struct {
	AddressStart int
	Length       uint
}

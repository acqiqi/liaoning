package melsecfxserial

import (
	"errors"
	"vgateway/kernel/serial"
)

/// <summary>
/// 三菱的串口通信的对象，适用于读取FX系列的串口数据，支持的类型参考文档说明
/// </summary>
/// <remarks>
/// 字读写地址支持的列表如下：
/// <list type="table">
///   <listheader>
///     <term>地址名称</term>
///     <term>地址代号</term>
///     <term>示例</term>
///     <term>地址范围</term>
///     <term>地址进制</term>
///     <term>备注</term>
///   </listheader>
///   <item>
///     <term>数据寄存器</term>
///     <term>D</term>
///     <term>D100,D200</term>
///     <term>D0-D511,D8000-D8255</term>
///     <term>10</term>
///     <term></term>
///   </item>
///   <item>
///     <term>定时器的值</term>
///     <term>TN</term>
///     <term>TN10,TN20</term>
///     <term>TN0-TN255</term>
///     <term>10</term>
///     <term></term>
///   </item>
///   <item>
///     <term>计数器的值</term>
///     <term>CN</term>
///     <term>CN10,CN20</term>
///     <term>CN0-CN199,CN200-CN255</term>
///     <term>10</term>
///     <term></term>
///   </item>
/// </list>
/// 位地址支持的列表如下：
/// <list type="table">
///   <listheader>
///     <term>地址名称</term>
///     <term>地址代号</term>
///     <term>示例</term>
///     <term>地址范围</term>
///     <term>地址进制</term>
///     <term>备注</term>
///   </listheader>
///   <item>
///     <term>内部继电器</term>
///     <term>M</term>
///     <term>M100,M200</term>
///     <term>M0-M1023,M8000-M8255</term>
///     <term>10</term>
///     <term></term>
///   </item>
///   <item>
///     <term>输入继电器</term>
///     <term>X</term>
///     <term>X1,X20</term>
///     <term>X0-X177</term>
///     <term>8</term>
///     <term></term>
///   </item>
///   <item>
///     <term>输出继电器</term>
///     <term>Y</term>
///     <term>Y10,Y20</term>
///     <term>Y0-Y177</term>
///     <term>8</term>
///     <term></term>
///   </item>
///   <item>
///     <term>步进继电器</term>
///     <term>S</term>
///     <term>S100,S200</term>
///     <term>S0-S999</term>
///     <term>10</term>
///     <term></term>
///   </item>
///   <item>
///     <term>定时器触点</term>
///     <term>TS</term>
///     <term>TS10,TS20</term>
///     <term>TS0-TS255</term>
///     <term>10</term>
///     <term></term>
///   </item>
///   <item>
///     <term>定时器线圈</term>
///     <term>TC</term>
///     <term>TC10,TC20</term>
///     <term>TC0-TC255</term>
///     <term>10</term>
///     <term></term>
///   </item>
///   <item>
///     <term>计数器触点</term>
///     <term>CS</term>
///     <term>CS10,CS20</term>
///     <term>CS0-CS255</term>
///     <term>10</term>
///     <term></term>
///   </item>
///   <item>
///     <term>计数器线圈</term>
///     <term>CC</term>
///     <term>CC10,CC20</term>
///     <term>CC0-CC255</term>
///     <term>10</term>
///     <term></term>
///   </item>
/// </list>

type MelsecFxSerial struct {
	SerialNo int //使用串口号 只针对使用串口协议
}

func (this *MelsecFxSerial) Begin() (err error) {
	return
}

// 操作写bool值
func (this *MelsecFxSerial) WriteBool(address string, status bool) (err error) {
	if !serial.IsSerialOpen {
		err = errors.New("请打开串口")
		return
	}
	cdata, err := BuildWriteBoolPacket(address, status)
	if err != nil {
		return
	}
	serial.SerialFlush()      //清空接收区
	serial.SerialWrite(cdata) //写入数据
	// 获取数据
	callbackdata, err := serial.ReadSerialOneData()
	if err != nil {
		err = errors.New("读取超时")
	}
	// 检测
	if err := CheckPlcWriteResponse(callbackdata); err != nil {

		return err
	}
	return
}

// 操作读取bool值
func (this *MelsecFxSerial) ReadBool(address string, length uint) (status []bool, err error) {
	if !serial.IsSerialOpen {
		err = errors.New("请打开串口")
		return
	}
	cdata, c3, err := BuildReadBoolCommand(address, length)
	if err != nil {
		return
	}
	serial.SerialFlush()      //清空接收区
	serial.SerialWrite(cdata) //写入数据
	// 获取数据
	callbackdata, err := serial.ReadSerialOneData()
	if err != nil {
		err = errors.New("读取超时")
	}
	// 检测
	if err := CheckPlcReadResponse(callbackdata); err != nil {
		return nil, err
	}
	status = ExtractActualBoolData(callbackdata, c3, length)
	return
}

// 操作写bytes
func (this *MelsecFxSerial) WriteBytes(address string, value []byte) (err error) {
	if !serial.IsSerialOpen {
		err = errors.New("请打开串口")
		return
	}
	cdata, err := BuildWriteWordCommand(address, value)
	if err != nil {
		return
	}
	serial.SerialFlush()      //清空接收区
	serial.SerialWrite(cdata) //写入数据
	// 获取数据
	callbackdata, err := serial.ReadSerialOneData()
	if err != nil {
		err = errors.New("读取超时")
	}
	// 检测
	if err := CheckPlcWriteResponse(callbackdata); err != nil {
		return err
	}
	return
}

func (this *MelsecFxSerial) ReadBytes(address string, length uint) (base []byte, err error) {
	if !serial.IsSerialOpen {
		err = errors.New("请打开串口")
		return
	}
	cdata, err := BuildReadWordCommand(address, length)
	if err != nil {
		return
	}
	serial.SerialFlush()      //清空接收区
	serial.SerialWrite(cdata) //写入数据
	// 获取数据
	callbackdata, err := serial.ReadSerialOneData()
	if err != nil {
		err = errors.New("读取超时")
	}
	// 检测
	if err := CheckPlcReadResponse(callbackdata); err != nil {
		return nil, err
	}
	//提炼数据
	base = ExtractActualData(callbackdata)
	return
}

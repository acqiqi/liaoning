package melsecserial

//--------------------------三菱PLC的数据类型，此处包含了几个常用的类型--------------------------
type MelsecA1EDataType struct {
	DataCode  []byte // 类型的代号值（软元件代码，用于区分软元件类型，如：D，R）
	DataType  byte   // 数据的类型，0代表按字，1代表按位
	AsciiCode string // 当以ASCII格式通讯时的类型描述
	FromBase  int    // 指示地址是10进制，还是16进制的

}

// 如果您清楚类型代号，可以根据值进行扩展
// code 数据类型的代号
// dataType 0或1，默认为0
//  ASCII格式的类型信息
// 指示地址的多少进制的，10或是16
func (this *MelsecA1EDataType) Init(code []byte, dataType byte, asciiCode string, fromBase int) {
	this.DataCode = code
	this.AsciiCode = asciiCode
	this.FromBase = fromBase
	if dataType < 2 {
		this.DataType = dataType
	}
}

// X输入寄存器
var MelsecA1EX = new(MelsecA1EDataType)

// Y输出寄存器
var MelsecA1EY = new(MelsecA1EDataType)

// M中间寄存器
var MelsecA1EM = new(MelsecA1EDataType)

// S状态寄存器
var MelsecA1ES = new(MelsecA1EDataType)

// D数据寄存器
var MelsecA1ED = new(MelsecA1EDataType)

// R文件寄存器
var MelsecA1ER = new(MelsecA1EDataType)

func init() {
	//type 初始化
	MelsecA1EX.Init([]byte{0x58, 0x20}, 0x01, "X*", 8)
	MelsecA1EY.Init([]byte{0x59, 0x20}, 0x01, "Y*", 8)
	MelsecA1EX.Init([]byte{0x4D, 0x20}, 0x01, "M*", 10)
	MelsecA1EX.Init([]byte{0x53, 0x20}, 0x01, "S*", 10)
	MelsecA1EX.Init([]byte{0x44, 0x20}, 0x00, "D*", 10)
	MelsecA1EX.Init([]byte{0x52, 0x20}, 0x00, "R*", 10)
}

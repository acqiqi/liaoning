package melsecserial

//--------------------------三菱PLC的数据类型，此处包含了几个常用的类型--------------------------
type MelsecMCDataType struct {
	DataCode  byte   // 类型的代号值（软元件代码，用于区分软元件类型，如：D，R）
	DataType  byte   // 数据的类型，0代表按字，1代表按位
	AsciiCode string // 当以ASCII格式通讯时的类型描述
	FromBase  int    // 指示地址是10进制，还是16进制的
}

// 如果您清楚类型代号，可以根据值进行扩展
// code 数据类型的代号
// dataType 0或1，默认为0
//  ASCII格式的类型信息
// 指示地址的多少进制的，10或是16
func (this *MelsecMCDataType) Init(code byte, dataType byte, asciiCode string, fromBase int) {
	this.DataCode = code
	this.AsciiCode = asciiCode
	this.FromBase = fromBase
	if dataType < 2 {
		this.DataType = dataType
	}
}

// X输入寄存器
var MelsecMCX = new(MelsecMCDataType)

// Y输出寄存器
var MelsecMCY = new(MelsecMCDataType)

// M中间寄存器
var MelsecMCM = new(MelsecMCDataType)

// D数据寄存器
var MelsecMCD = new(MelsecMCDataType)

// W链接寄存器
var MelsecMCW = new(MelsecMCDataType)

// L锁存继电器
var MelsecMCL = new(MelsecMCDataType)

// F报警器
var MelsecMCF = new(MelsecMCDataType)

// V边沿继电器
var MelsecMCV = new(MelsecMCDataType)

// B链接继电器
var MelsecMCB = new(MelsecMCDataType)

// R文件寄存器
var MelsecMCR = new(MelsecMCDataType)

// S步进继电器
var MelsecMCS = new(MelsecMCDataType)

// 变址寄存器
var MelsecMCZ = new(MelsecMCDataType)

// 定时器的当前值
var MelsecMCTN = new(MelsecMCDataType)

// 定时器的触点
var MelsecMCTS = new(MelsecMCDataType)

// 定时器的线圈
var MelsecMCTC = new(MelsecMCDataType)

// 累计定时器的触点
var MelsecMCSS = new(MelsecMCDataType)

// 累计定时器的线圈
var MelsecMCSC = new(MelsecMCDataType)

// 累计定时器的当前值
var MelsecMCSN = new(MelsecMCDataType)

// 计数器的当前值
var MelsecMCCN = new(MelsecMCDataType)

// 计数器的触点
var MelsecMCCS = new(MelsecMCDataType)

// 计数器的线圈
var MelsecMCCC = new(MelsecMCDataType)

// 文件寄存器ZR区
var MelsecMCZR = new(MelsecMCDataType)

func init() {
	//type 初始化
	MelsecMCX.Init(0x9C, 0x01, "X*", 16)
	MelsecMCY.Init(0x9D, 0x01, "Y*", 16)
	MelsecMCM.Init(0x90, 0x01, "M*", 10)
	MelsecMCD.Init(0xA8, 0x00, "D*", 10)
	MelsecMCW.Init(0xB4, 0x00, "W*", 16)
	MelsecMCL.Init(0x92, 0x01, "L*", 10)
	MelsecMCF.Init(0x93, 0x01, "F*", 10)
	MelsecMCV.Init(0x94, 0x01, "V*", 10)
	MelsecMCB.Init(0xA0, 0x01, "B*", 16)
	MelsecMCR.Init(0xAF, 0x00, "R*", 10)
	MelsecMCS.Init(0x98, 0x01, "S*", 10)
	MelsecMCZ.Init(0xCC, 0x00, "Z*", 10)
	MelsecMCTN.Init(0xC2, 0x00, "TN", 10)
	MelsecMCTS.Init(0xC1, 0x01, "TS", 10)
	MelsecMCTC.Init(0xC0, 0x01, "TC", 10)
	MelsecMCSS.Init(0xC7, 0x01, "SS", 10)
	MelsecMCSC.Init(0xC6, 0x01, "SC", 10)
	MelsecMCSN.Init(0xC8, 0x00, "SN", 100)
	MelsecMCCN.Init(0xC5, 0x00, "CN", 10)
	MelsecMCCS.Init(0xC4, 0x01, "CS", 10)
	MelsecMCCC.Init(0xC3, 0x01, "CC", 10)
	MelsecMCZR.Init(0xB0, 0x00, "ZR", 16)
}

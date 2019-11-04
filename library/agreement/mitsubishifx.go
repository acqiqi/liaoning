package agreement

import (
	"encoding/hex"
)

type MitsubishiFX struct {
}

func (this MitsubishiFX) GetDOOnOffCode(driveAddr uint16, funcaddr uint32, onOff bool) (callback []byte) {
	uSend := []byte{0x02, 0x37, 0x30, 0x30, 0x30, 0x30, 0x03, 0x30, 0x30}

	if onOff {
		uSend[1] = 0x37
	} else {
		uSend[1] = 0x38
	}

	uAddress_y := funcaddr
	uAddress_y = uAddress_y - (uAddress_y/100)*20
	uAddress_y = uAddress_y - (uAddress_y/10)*2
	uAddress_y = uAddress_y % 8
	funcaddr = funcaddr - (funcaddr/100)*20
	funcaddr = funcaddr - (funcaddr/10)*2
	funcaddr = funcaddr/8 + 0x00A0
	funcaddr = funcaddr*8 + uAddress_y
	uTmp := funcaddr & 0x000f
	if uint8(uTmp) < 10 {
		uSend[3] = uint8(uTmp) + 0x30
	} else {
		uSend[3] = uint8(uTmp) + 0x41 - 0xa
	}

	uTmp = (funcaddr >> 4) & 0x000f
	if uint8(uTmp) < 10 {
		uSend[2] = uint8(uTmp) + 0x30
	} else {
		uSend[2] = uint8(uTmp) + 0x41 - 0xa
	}

	uTmp = (funcaddr >> 8) & 0x000f
	if uint8(uTmp) < 10 {
		uSend[5] = uint8(uTmp) + 0x30
	} else {
		uSend[5] = uint8(uTmp) + 0x41 - 0xa
	}

	uTmp = (funcaddr >> 12) & 0x000f
	if uint8(uTmp) < 10 {
		uSend[4] = uint8(uTmp) + 0x30
	} else {
		uSend[4] = uint8(uTmp) + 0x41 - 0xa
	}

	var uSum uint32 = 0
	for i := 1; i < 7; i++ {
		uSum = uSum + uint32(uSend[i])
	}
	uTmp = uSum & 0x000f
	if uint8(uTmp) < 10 {
		uSend[8] = uint8(uTmp) + 0x30
	} else {
		uSend[8] = uint8(uTmp) + 0x41 - 0xa
	}

	uTmp = (uSum >> 4) & 0x000f
	if uint8(uTmp) < 10 {
		uSend[7] = uint8(uTmp) + 0x30
	} else {
		uSend[7] = uint8(uTmp) + 0x41 - 0xa
	}

	//for i := 0; i < 9; i++ {
	//	if i==0 || i== 6 {
	//
	//	} else {
	//		uSend[i] = uSend[i] - 18
	//	}
	//}

	return uSend
}

func (this MitsubishiFX) GetDoReadCode(driveAddr uint16, funcaddr uint32) (callback []byte) {
	uSend := []byte{0x02, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x03, 0x30, 0x30}
	//数据位初始化
	var uCount uint32 = 1

	if (uCount / 16) >= 10 {
		uSend[6] = uint8((uCount/16 + 0x41 - 10))
	} else {
		uSend[6] = uint8(uCount/16 + 0x30)
	}

	if (uCount % 16) >= 10 {
		uSend[7] = uint8(((uCount % 16) + 0x41 - 10))
	} else {
		uSend[7] = uint8((uCount % 16) + 0x30)
	}

	funcaddr = funcaddr - (funcaddr/100)*20
	funcaddr = funcaddr - (funcaddr/10)*2
	funcaddr = funcaddr/8 + 0x00A0
	var uTmp uint32 = funcaddr & 0x000f
	if uint8(uTmp) < 10 {
		uSend[5] = uint8(uTmp) + 0x30
	} else {
		uSend[5] = uint8(uTmp) + 0x41 - 0xa
	}
	uTmp = (funcaddr >> 4) & 0x000f
	if uint8(uTmp) < 10 {
		uSend[4] = uint8(uTmp) + 0x30
	} else {
		uSend[4] = uint8(uTmp) + 0x41 - 0xa
	}
	uTmp = (funcaddr >> 8) & 0x000f
	if uint8(uTmp) < 10 {
		uSend[3] = uint8(uTmp) + 0x30
	} else {
		uSend[3] = uint8(uTmp) + 0x41 - 0xa
	}
	uTmp = (funcaddr >> 12) & 0x000f
	if uint8(uTmp) < 10 {
		uSend[2] = uint8(uTmp) + 0x30
	} else {
		uSend[2] = uint8(uTmp) + 0x41 - 0xa
	}

	var uSum uint32 = 0
	for i := 1; i < 9; i++ {
		uSum = uSum + uint32(uSend[i])
	}

	uTmp = uSum & 0x000f
	if uint8(uTmp) < 10 {
		uSend[10] = uint8(uTmp) + 0x30
	} else {
		uSend[10] = uint8(uTmp) + 0x41 - 0xa
	}
	uTmp = (uSum >> 4) & 0x000f
	if uint8(uTmp) < 10 {
		uSend[9] = uint8(uTmp) + 0x30
	} else {
		uSend[9] = uint8(uTmp) + 0x41 - 0xa
	}

	callback = uSend
	return
}

/**
三菱PLC读DO返回数据 Bool
*/
func (this MitsubishiFX) DecodeDoReadCode(uData []byte, uAddress uint32) (callback bool) {

	be := []byte{uData[1], uData[2]}
	str := hex.EncodeToString(be)
	for {
		if len(str) >= 8 {
			break
		}
		str = "0" + str
	}

	uAddress = uAddress - (uAddress/100)*20
	uAddress = uAddress - (uAddress/10)*2
	uAddress = uAddress % 8

	newBee, _ := hex.DecodeString(string(be))

	if (newBee[0] >> uint8(uAddress) & 1) == 1 {
		return true
	} else {
		return false
	}
}

/**
FX系列校验
*/
func MitsubishiFxVerification(data []byte) (basedata []byte) {
	return
}

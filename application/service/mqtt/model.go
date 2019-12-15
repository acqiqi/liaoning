package mqtt

import (
	"errors"
	"log"
)

type PublishData struct {
	MessageId string      `json:"message_id"` //消息id
	Topic     string      `json:"topic"`
	Type      string      `json:"type"`    //类型
	Msg       string      `json:"msg"`     //消息
	Data      interface{} `json:"data"`    //数据
	Value     string      `json:"value"`   //值
	IsWait    bool        `json:"is_wait"` //是否等待数据反馈
	Code      int         `json:"code"`    //0正常 非0异常
}

type SubscribeData struct {
	MessageId string            `json:"message_id"` //消息id
	Topic     string            `json:"topic"`
	Type      string            `json:"type"`      //类型
	DataType  string            `json:"data_type"` //数据类型
	Msg       string            `json:"msg"`       //消息
	Data      map[string]string `json:"data"`      //数据
	Value     string            `json:"value"`     //值
	ValueInt  int               `json:"value_int"`
	ValFloat  float64           `json:"val_float"`
	IsWait    bool              `json:"is_wait"` //是否等待数据反馈
	Code      int               `json:"code"`    //0正常 非0异常
}

// 配置模式
func (s *SubscribeData) Config() (err error) {

	return
}

// 全局配置模式
func (s *SubscribeData) All() (err error) {

	return
}

// 读数据模式
func (s *SubscribeData) Read() (err error) {

	return
}

// 写模式
func (s *SubscribeData) Write() (err error) {

	return
}

// OTA
func (s *SubscribeData) OTA() (err error) {

	return
}

// 复位重置模式
func (s *SubscribeData) Rst() (err error) {

	return
}

// 功能锁
func (s *SubscribeData) Lock() (err error) {

	return
}

// Lib PLC协议配置
func (s *SubscribeData) Lib() (err error) {
	if s.DataType == "write" {
		LibDriverOne.DriverType = s.Data["driver_type"]
		LibDriverOne.DriverAddress = s.Data["driver_address"]
		LibDriverOne.DriverPort = s.Data["driver_port"]
		LibDriverOne.SerialNo = s.Data["serial_no"]
		LibDriverOne.PlcFlag = s.ValueInt //特殊plc型号标识
		log.Println(LibDriverOne)
		if err := LibDriverOne.Init(); err != nil {
			return err
		}
	} else if s.DataType == "read" {
		PublishCallbackSuccess(PublishData{
			MessageId: s.MessageId,
			Topic:     s.Topic,
			Type:      s.Type,
			Data:      *LibDriverOne,
		}, "success")
	} else {
		return errors.New("datatype err")
	}

	return
}

// 透传模式  如果开透传就不允许PLC协议配置
func (s *SubscribeData) PT() (err error) {

	return
}

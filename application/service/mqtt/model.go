package mqtt

import (
	"errors"
	"log"
	"strconv"
	"vgateway/common"
	"vgateway/kernel/config"
	"vgateway/library/agreement"
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
	IsWait    bool              `json:"is_wait"`   //是否等待数据反馈
	Code      int               `json:"code"`      //0正常 非0异常
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
	switch s.DataType {
	case "bool":
		value, err := strconv.Atoi(s.Value)
		if err != nil {
			return errors.New("value decode err")
		}
		val := common.If(value == 1, true, false)
		if err := agreement.LibDriverOne.WriteBool(s.Data["addr"], val.(bool)); err != nil {
			return err
		}
	}

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
		agreement.LibDriverOne.DriverType = s.Data["driver_type"]
		agreement.LibDriverOne.DriverAddress = s.Data["driver_address"]
		agreement.LibDriverOne.DriverPort = s.Data["driver_port"]
		agreement.LibDriverOne.SerialNo = s.Data["serial_no"]

		plc_flag, err := strconv.Atoi(s.Data["plc_flag"])
		if err != nil {
			return err
		}
		agreement.LibDriverOne.PlcFlag = plc_flag //特殊plc型号标识

		log.Println("标识")
		log.Println(agreement.LibDriverOne)
		if err := agreement.LibDriverOne.Init(); err != nil {
			return err
		}

		//初始化设备
		if err := agreement.LibDriverOne.InitDriver(); err != nil {
			return err
		}

		//保存配置文件
		config.ConfigObject.Lib.DriverType = agreement.LibDriverOne.DriverType
		config.ConfigObject.Lib.DriverAddress = agreement.LibDriverOne.DriverAddress
		config.ConfigObject.Lib.DriverPort = agreement.LibDriverOne.DriverPort
		config.ConfigObject.Lib.SerialNo = agreement.LibDriverOne.SerialNo
		config.ConfigObject.Lib.PlcFlag = plc_flag //特殊plc型号标识
		config.SaveConfig("lib")

		PublishCallbackSuccess(PublishData{
			MessageId: s.MessageId,
			Topic:     s.Topic,
			Type:      s.Type,
			Data:      *agreement.LibDriverOne,
		}, "success")

	} else if s.DataType == "read" {
		PublishCallbackSuccess(PublishData{
			MessageId: s.MessageId,
			Topic:     s.Topic,
			Type:      s.Type,
			Data:      *agreement.LibDriverOne,
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

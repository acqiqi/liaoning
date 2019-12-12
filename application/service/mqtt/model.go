package mqtt

type MqttPublishData struct {
	MessageId string      `json:"message_id"` //消息id
	Type      string      `json:"type"`       //类型
	Msg       string      `json:"msg"`        //消息
	Data      interface{} `json:"data"`       //数据
	Value     string      `json:"value"`      //值
	IsWait    bool        `json:"is_wait"`    //是否等待数据反馈
}

type MqttSubscribeData struct {
	MessageId string      `json:"message_id"` //消息id
	Type      string      `json:"type"`       //类型
	Msg       string      `json:"msg"`        //消息
	Data      interface{} `json:"data"`       //数据
	Value     string      `json:"value"`      //值
	IsWait    bool        `json:"is_wait"`    //是否等待数据反馈
}

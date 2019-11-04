package utils_wechat

// 微信小程序登录的
type RespWXSmallModel struct {
	Openid     string `json:"openid"`      //用户唯一标识
	Sessionkey string `json:"session_key"` //会话密钥
	Unionid    string `json:"unionid"`     //用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
	Errcode    int    `json:"errcode"`     //错误码
	ErrMsg     string `json:"errmsg"`      //错误信息
}

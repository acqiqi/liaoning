package common

type PageList struct {
	Page       int64       `json:"page"`
	Limit      int64       `json:"limit"`
	TotalCount int64       `json:"total_count"`
	List       interface{} `json:"list"`
}

type PushPostMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ApiCodeList struct {
	ApiSuccess         ApiCodeItem
	ApiError           ApiCodeItem
	ApiNotAuth         ApiCodeItem
	ApiNotBindMobile   ApiCodeItem
	ApiNotBindUserInfo ApiCodeItem
}
type ApiCodeItem struct {
	Code  int
	MsgCN string //中文信息
}

var apiCodeList = ApiCodeList{
	ApiSuccess: ApiCodeItem{
		Code:  0,
		MsgCN: "操作成功",
	}, //成功
	ApiError: ApiCodeItem{
		Code:  1,
		MsgCN: "操作失败",
	}, //失败
	ApiNotAuth: ApiCodeItem{
		Code:  45000,
		MsgCN: "登录超时",
	}, //登录超时
	ApiNotBindMobile: ApiCodeItem{
		Code:  41001,
		MsgCN: "请绑定手机号",
	}, //没有绑定手机号
	ApiNotBindUserInfo: ApiCodeItem{
		Code:  41002,
		MsgCN: "请绑定用户信息",
	}, //没有绑定用户信息
}

// Api状态码
func ApiStatusCode() ApiCodeList {
	return apiCodeList
}

// 提供空 obj类型
func GetEmptyStruct() interface{} {
	return struct {
	}{}
}

func GetListStruct(list interface{}, page int64, limit int) interface{} {
	return struct {
		List  interface{} `json:"list"`
		Page  int64       `json:"page"`
		Limit int         `json:"limit"`
	}{
		List:  list,
		Page:  page,
		Limit: limit,
	}
}

func ApiJsonSuccess(msg string, data interface{}) PushPostMsg {
	if msg == "" {
		msg = "操作成功"
	}
	obj := PushPostMsg{
		Code: ApiStatusCode().ApiSuccess.Code,
		Msg:  msg,
		Data: data,
	}
	return obj
}

func ApiJsonError(msg string) PushPostMsg {
	if msg == "" {
		msg = "操作有误"
	}
	obj := PushPostMsg{
		Code: ApiStatusCode().ApiError.Code,
		Msg:  msg,
	}
	return obj
}

func ApiJsonOth(code int, msg string, data interface{}) PushPostMsg {
	obj := PushPostMsg{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	return obj
}

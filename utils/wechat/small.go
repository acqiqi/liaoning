package utils_wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func LoginWXSmall(code string) (wxInfo RespWXSmallModel, err error) {
	//https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	var model RespWXSmallModel

	appId := "wx7565814c8826de10"
	appSecret := "49e4eb50a387a74981356c536d10deca"
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	resp, err := http.Get(fmt.Sprintf(url, appId, appSecret, code))
	if err != nil {
		return wxInfo, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RespWXSmallModel{}, err
	}
	fmt.Println("Http Body ---->", string(body))
	if err := json.Unmarshal(body, &model); err != nil {
		return RespWXSmallModel{}, errors.New("请输入正确的json格式")
	}
	if model.ErrMsg == "" {
		return model, nil
	} else {
		return RespWXSmallModel{}, errors.New("123123")
		//return model,errors.New(model.ErrMsg)
	}
}

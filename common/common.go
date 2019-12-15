package common

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PostJson struct {
	Data interface{}
}

func GetPostJsonObj(body []byte, obj interface{}) (err error) {
	if err := json.Unmarshal(body, &obj); err != nil {
		fmt.Println(err.Error())
		return errors.New("请输入正确的json格式")
	}
	return nil
}

func JsonDecode(body string, obj interface{}) (err error) {
	strByte := []byte(body)
	if err := json.Unmarshal(strByte, &obj); err != nil {
		fmt.Println(err.Error())
		return errors.New("string 格式不正确")
	}
	return nil
}

func JsonEncode(obj interface{}) string {
	b, _ := json.Marshal(obj)
	return string(b)
}

//加密密码
func PasswordEncrypt(password string) string {
	key := "vhake"
	password = password + key
	data := []byte(password)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

//解密密码 数据库密码  输入密码
func PasswordDecrypt(dbpassword string, password string) bool {
	key := "vhake"
	password = password + key
	data := []byte(password)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	if md5str == dbpassword {
		return true
	} else {
		return false
	}
}

// 发送短信验证码
func SendSMS(mobile string, msg string) (err error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST",
		"http://sms-api.luosimao.com/v1/send.json",
		strings.NewReader("mobile="+mobile+"&message="+msg+"【Vcloudshop】"))
	if err != nil {
		return errors.New(err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("api", "key-8038fe9d1e3f6f5710c6df1205c6e5aa")
	resp, err := client.Do(req)
	if err != nil {
		return errors.New(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(err.Error())
	}
	fmt.Println(string(body))
	return nil
}

func SetUserToken(uid uint64) (token string) {
	crutime := time.Now().UnixNano()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	tk := fmt.Sprintf("%x", h.Sum(nil))

	SetCache("VToken"+tk, uid, 2592000*time.Second)
	return tk
}

func GetUserToken(token string) (uid int64) {
	c, err := GetCache("VToken" + token)

	if err != nil {
		return 0
	} else {
		uId, _ := strconv.ParseInt(c, 10, 64)
		return uId
	}
}

func If(b bool, t, f interface{}) interface{} {
	if b {
		return t
	}
	return f
}

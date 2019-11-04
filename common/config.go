package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type CfgStruct struct {
	DbHost string `json:"db_host"`
}

var cfgObj = make(map[string]interface{})

func GetConfig(key string) string {
	val := cfgObj[key]
	if val == nil {
		return ""
	}
	return val.(string)
}

//配置文件初始化 如果失败必须停机

func init() {
	cfgByte, err := ioutil.ReadFile("./cfg.json")
	if err != nil {
		fmt.Println("配置文件打开错误")
		log.Println("config red err !")
		os.Exit(0)
	}
	//m := make(map[string]interface{})
	err = json.Unmarshal(cfgByte, &cfgObj)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

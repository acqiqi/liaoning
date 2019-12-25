package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func init() {
	initConfig()
}

//初始化配置文件
func initConfig() {
	readFile("system", &ConfigObject.System)
	readFile("lib", &ConfigObject.Lib)

}

// 保存配置文件
func SaveConfig(name string) {
	switch name {
	case "system": //系统配置
		writeFile("system", ConfigObject.System)
		break
	case "lib":
		writeFile("lib", ConfigObject.Lib)
	default: //全部配置
		writeFile("system", ConfigObject.System)
		writeFile("lib", ConfigObject.Lib)
	}
}

// 写配置文件
func writeFile(fileName string, data interface{}) (err error) {
	b, _ := json.Marshal(data)
	err = ioutil.WriteFile("./config/"+fileName+".json", b, 0666)
	return
}

// 读配置文件
func readFile(fileName string, data interface{}) (err error) {
	cfgByte, err := ioutil.ReadFile("./config/" + fileName + ".json")
	if err != nil {
		fmt.Println("配置文件打开错误")
		log.Println("config red err !")
		return
	}
	//m := make(map[string]interface{})
	err = json.Unmarshal(cfgByte, &data)
	return
}

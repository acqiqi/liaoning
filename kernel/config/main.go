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

func initConfig() {
	readFile("system", &ConfigObject.System)
}

func SaveConfig(name string) {
	switch name {
	case "system": //系统配置
		writeFile("system", ConfigObject.System)
		break
	default: //全部配置
		writeFile("system", ConfigObject.System)
	}
}

func writeFile(fileName string, data interface{}) (err error) {
	b, _ := json.Marshal(data)
	err = ioutil.WriteFile("./config/"+fileName+".json", b, 0666)
	return
}

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

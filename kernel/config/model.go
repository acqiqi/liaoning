package config

var ConfigObject = new(ConfigModel)

type ConfigModel struct {
	System SystemModel
	Lib    LibModel
}

type SystemModel struct {
	Cache struct {
		Driver         string `json:"driver"`
		Prefix         string `json:"prefix"`
		ServerHost     string `json:"server_host"`
		ServerPort     int    `json:"server_port"`
		ServerUsername string `json:"server_username"`
		ServerPassword string `json:"server_password"`
		FilesPath      string `json:"files_path"`
	} `json:"cache"`
	Database struct {
		Driver         string `json:"driver"`
		Prefix         string `json:"prefix"`
		ServerHost     string `json:"server_host"`
		ServerPort     int    `json:"server_port"`
		ServerUsername string `json:"server_username"`
		ServerPassword string `json:"server_password"`
		ServerName     string `json:"server_name"`
	} `json:"database"`
	Uploader struct {
		Driver string `json:"driver"`
		Domain string `json:"domain"`
		Ak     string `json:"ak"`
		Sk     string `json:"sk"`
		Bucket string `json:"bucket"`
	} `json:"uploader"`
	Sms struct {
		Driver string `json:"driver"`
		Key    string `json:"key"`
		Prefix string `json:"prefix"`
	} `json:"sms"`
	Email struct {
		Host     string `json:"host"`
		Password string `json:"password"`
	} `json:"email"`
}

type LibModel struct {
	DriverType    string `json:"driver_type"`
	DriverAddress string `json:"driver_address"`
	DriverPort    string `json:"driver_port"`
	SerialNo      string `json:"serial_no"`
	PlcFlag       int    `json:"plc_flag"`
}

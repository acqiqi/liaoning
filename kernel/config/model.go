package config

var ConfigObject = new(ConfigModel)

type ConfigModel struct {
	System SystemModel
}

type SystemModel struct {
	AdminUser string `json:"admin_user"`
	AdminPwd  string `json:"admin_pwd"`
}

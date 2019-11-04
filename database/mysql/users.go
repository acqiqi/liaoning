package mysql

import (
	"fmt"
	"vgateway/database/model"
)

type Users struct {
	model.Users
}

//简单的操作模型
func (this *Users) GetUserInfo(id int64) (user Users, err error) {
	if err := DB.Where("id = ?", id).Find(&this).Error; err != nil {
		fmt.Println(err.Error())
		return *this, err
	}
	return *this, nil
}

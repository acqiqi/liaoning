package model

import (
	"time"
)

type Users struct {
	Id          uint64    `gorm:"primary_key" json:"id"`
	Username    string    `gorm:"size(32)" json:"username"`
	Nickname    string    `gorm:"size(50)" json:"nickname"`
	Password    string    `gorm:"size(100)" json:"password"`
	Email       string    `gorm:"size(100)" json:"email"`
	Mobile      string    `gorm:"size(11)" json:"mobile"`
	Avatar      string    `gorm:"size(200)" json:"avatar"`
	Money       float64   `json:"money"`
	Score       int64     `json:"score"`
	SmallOpenid string    `json:"small_openid"` //small openid
	CreatedAt   time.Time `gorm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt   time.Time `gorm:"auto_now;type(datetime)" json:"update_at"`
	BindMobile  int       `gorm:"default:0" json:"bind_mobile"`
} // 默认表名是`users`

package model

import (
	"vgateway/utils"
)

type BaseModel struct {
	Id        uint64          `gorm:"primary_key" json:"id"`
	Flag      int8            `json:"flag"`
	CreatedAt utils.JSONTime  `json:"created_at"`
	UpdatedAt utils.JSONTime  `json:"updated_at"`
	DeletedAt *utils.JSONTime `json:"-"`
}

package service_swa

import (
	"github.com/ServiceWeaver/weaver"
	"time"
)

type ModDict struct {
	weaver.AutoMarshal
	ModBase
	Name   string `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`
	Type   string `json:"type" form:"type" gorm:"column:type;comment:字典名（英）"`
	Status *bool  `json:"status" form:"status" gorm:"column:status;comment:状态"`
	Detail string `json:"detail" form:"detail" gorm:"column:detail;comment:描述"`
	ModDictails []ModDictail `json:"swaDictail" gorm:"-"`
}

func (ModDict) TableName() string {
	return "swa_dicts"
}

type SearchInfoDict struct {
	weaver.AutoMarshal
	ModDict
	PageInfo
	StartCreatedAt time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	OrderKey       string    `json:"orderKey"`
	Desc           bool      `json:"desc"`
}

type SwaDictSearch struct {
	weaver.AutoMarshal
	ModDict
	PageInfo
}
